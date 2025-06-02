// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title RecargaVE
 * @dev Contrato para gerenciar reservas e recargas de Veículos Elétricos.
 * Os preços e pagamentos são feitos em Wei (a menor unidade de Ether).
 */
contract RecargaVE {

    // Estruturas de Dados
    struct Posto {
        uint idPosto;
        address payable proprietario;
        string nome;
        string localizacao;
        uint precoPorUnidadeEnergia; // Em Wei por unidade de energia
        bool disponivelParaNovasReservas; // Se o posto está aceitando novas reservas
        uint256 totalRecebido; // Total de Ether recebido por este posto
    }

    struct Reserva {
        uint idReserva;
        uint idPosto;
        address enderecoVE;
        uint horarioInicio; // Timestamp Unix
        uint horarioFim;   // Timestamp Unix
        bool ativa;        // Se a reserva ainda é válida (não cancelada ou utilizada)
        bool sessaoIniciada;
        uint idSessaoVinculada; // ID da sessão de recarga se iniciada
    }

    struct SessaoRecarga {
        uint idSessao;
        uint idReserva;
        address enderecoVE;
        uint idPosto;
        uint horarioInicioReal; // Timestamp Unix
        uint horarioFimReal;   // Timestamp Unix
        uint energiaConsumida;
        uint custoTotal;       // Em Wei
        bool paga;
    }

    // Contadores para IDs
    uint public proximoIdPosto = 1;
    uint public proximoIdReserva = 1;
    uint public proximoIdSessao = 1;

    // Mapeamentos
    mapping(uint => Posto) public postos;
    mapping(uint => Reserva) public reservas;
    mapping(uint => SessaoRecarga) public sessoesRecarga;

    // Mapeamento para ajudar a encontrar reservas de um VE ou de um Posto
    mapping(address => uint[]) public reservasPorVE;
    mapping(uint => uint[]) public reservasPorPosto;


    // Eventos
    event PostoRegistrado(
        uint indexed idPosto,
        address indexed proprietario,
        string nome,
        string localizacao,
        uint precoPorUnidadeEnergia
    );

    event DisponibilidadePostoAtualizada(
        uint indexed idPosto,
        bool disponivel
    );

    event PrecoPostoAtualizado(
        uint indexed idPosto,
        uint novoPreco
    );

    event ReservaFeita(
        uint indexed idReserva,
        uint indexed idPosto,
        address indexed enderecoVE,
        uint horarioInicio,
        uint horarioFim
    );

    event ReservaCancelada(
        uint indexed idReserva,
        address indexed canceladoPor
    );

    event RecargaIniciada(
        uint indexed idSessao,
        uint indexed idReserva,
        uint idPosto,
        address indexed enderecoVE,
        uint horarioInicioReal
    );

    event RecargaFinalizada(
        uint indexed idSessao,
        uint idPosto,
        address indexed enderecoVE,
        uint energiaConsumida,
        uint custoTotal,
        bool paga
    );

    event PagamentoRetiradoPosto(
        uint indexed idPosto,
        address indexed proprietario,
        uint256 valor
    );

    // Modificadores
    modifier apenasProprietarioPosto(uint _idPosto) {
        require(postos[_idPosto].proprietario == msg.sender, "Apenas o proprietario do posto pode chamar esta funcao.");
        _;
    }

    modifier apenasEnderecoVE(uint _idReserva) {
        require(reservas[_idReserva].enderecoVE == msg.sender, "Apenas o VE da reserva pode chamar esta funcao.");
        _;
    }


    constructor() {
        // Opcional: lógica de inicialização
    }

    // Funções para Gerenciamento de Postos
    function registrarPosto(
        string calldata _nome,
        string calldata _localizacao,
        uint _precoPorUnidadeEnergia
    ) external {
        require(_precoPorUnidadeEnergia > 0, "Preco deve ser maior que zero.");
        uint idPosto = proximoIdPosto++;
        postos[idPosto] = Posto({
            idPosto: idPosto,
            proprietario: payable(msg.sender),
            nome: _nome,
            localizacao: _localizacao,
            precoPorUnidadeEnergia: _precoPorUnidadeEnergia,
            disponivelParaNovasReservas: true,
            totalRecebido: 0
        });
        emit PostoRegistrado(idPosto, msg.sender, _nome, _localizacao, _precoPorUnidadeEnergia);
    }

    function atualizarDisponibilidadePosto(uint _idPosto, bool _disponivel) external apenasProprietarioPosto(_idPosto) {
        postos[_idPosto].disponivelParaNovasReservas = _disponivel;
        emit DisponibilidadePostoAtualizada(_idPosto, _disponivel);
    }

    function atualizarPrecoPosto(uint _idPosto, uint _novoPreco) external apenasProprietarioPosto(_idPosto) {
        require(_novoPreco > 0, "Preco deve ser maior que zero.");
        postos[_idPosto].precoPorUnidadeEnergia = _novoPreco;
        emit PrecoPostoAtualizado(_idPosto, _novoPreco);
    }

    function sacarSaldoPosto(uint _idPosto) external apenasProprietarioPosto(_idPosto) {
        uint256 saldo = postos[_idPosto].totalRecebido;
        require(saldo > 0, "Posto sem saldo para sacar.");
        postos[_idPosto].totalRecebido = 0;
        (bool success, ) = postos[_idPosto].proprietario.call{value: saldo}("");
        require(success, "Falha ao transferir o saldo.");
        emit PagamentoRetiradoPosto(_idPosto, msg.sender, saldo);
    }

    // Funções para Gerenciamento de Reservas
    function fazerReserva(uint _idPosto, uint _horarioInicio, uint _duracaoMinutos) external {
        require(postos[_idPosto].idPosto != 0, "Posto nao existe.");
        require(postos[_idPosto].disponivelParaNovasReservas, "Posto nao esta disponivel para novas reservas.");
        require(_duracaoMinutos > 0, "Duracao deve ser maior que zero.");
        require(_horarioInicio >= block.timestamp, "Horario de inicio deve ser no futuro.");

        uint _horarioFim = _horarioInicio + (_duracaoMinutos * 1 minutes);
        require(_horarioFim > _horarioInicio, "Horario de fim invalido.");

        // Verificar sobreposição de reservas
        uint[] storage idsReservasDoPosto = reservasPorPosto[_idPosto];
        for (uint i = 0; i < idsReservasDoPosto.length; i++) {
            Reserva storage reservaExistente = reservas[idsReservasDoPosto[i]];
            if (reservaExistente.ativa) {
                // Verifica se (inicioNova < fimExistente) E (fimNova > inicioExistente)
                bool sobrepoe = (_horarioInicio < reservaExistente.horarioFim) && (_horarioFim > reservaExistente.horarioInicio);
                require(!sobrepoe, "Horario solicitado esta em conflito com reserva existente.");
            }
        }

        uint idReserva = proximoIdReserva++;
        reservas[idReserva] = Reserva({
            idReserva: idReserva,
            idPosto: _idPosto,
            enderecoVE: msg.sender,
            horarioInicio: _horarioInicio,
            horarioFim: _horarioFim,
            ativa: true,
            sessaoIniciada: false,
            idSessaoVinculada: 0
        });

        reservasPorVE[msg.sender].push(idReserva);
        reservasPorPosto[_idPosto].push(idReserva);

        emit ReservaFeita(idReserva, _idPosto, msg.sender, _horarioInicio, _horarioFim);
    }

    function cancelarReserva(uint _idReserva) external {
        Reserva storage reserva = reservas[_idReserva];
        require(reserva.idReserva != 0, "Reserva nao existe.");
        require(reserva.ativa, "Reserva ja esta inativa.");
        require(!reserva.sessaoIniciada, "Nao pode cancelar reserva com sessao iniciada.");

        // Permite cancelar se for o VE dono da reserva ANTES do horário de início
        // Ou qualquer um pode cancelar se o horário da reserva já passou e não foi iniciada
        bool podeCancelar = (reserva.enderecoVE == msg.sender && block.timestamp < reserva.horarioInicio);
        bool expiradaNaoIniciada = (block.timestamp > reserva.horarioFim && !reserva.sessaoIniciada);
        
        require(podeCancelar || expiradaNaoIniciada, "Condicoes para cancelamento nao atendidas.");

        reserva.ativa = false;
        emit ReservaCancelada(_idReserva, msg.sender);
    }

    // Funções para Gerenciamento de Sessões de Recarga
    function iniciarRecarga(uint _idReserva) external {
        Reserva storage reserva = reservas[_idReserva];
        require(reserva.idReserva != 0, "Reserva nao existe.");
        require(reserva.ativa, "Reserva nao esta ativa.");
        require(reserva.enderecoVE == msg.sender, "Apenas o VE da reserva pode iniciar.");
        require(!reserva.sessaoIniciada, "Sessao de recarga ja foi iniciada para esta reserva.");
        require(block.timestamp >= reserva.horarioInicio, "Ainda nao eh horario de inicio da reserva.");
        require(block.timestamp <= reserva.horarioFim, "Horario da reserva expirou."); // Ou uma pequena tolerância

        uint idSessao = proximoIdSessao++;
        sessoesRecarga[idSessao] = SessaoRecarga({
            idSessao: idSessao,
            idReserva: _idReserva,
            enderecoVE: msg.sender,
            idPosto: reserva.idPosto,
            horarioInicioReal: block.timestamp,
            horarioFimReal: 0, // Será definido ao finalizar
            energiaConsumida: 0, // Será definido ao finalizar
            custoTotal: 0, // Será definido ao finalizar
            paga: false
        });

        reserva.sessaoIniciada = true;
        reserva.idSessaoVinculada = idSessao;

        emit RecargaIniciada(idSessao, _idReserva, reserva.idPosto, msg.sender, block.timestamp);
    }

    function finalizarRecarga(uint _idSessao, uint _energiaConsumida) external payable {
        SessaoRecarga storage sessao = sessoesRecarga[_idSessao];
        Reserva storage reserva = reservas[sessao.idReserva];
        Posto storage posto = postos[sessao.idPosto];

        require(sessao.idSessao != 0, "Sessao de recarga nao existe.");
        require(sessao.enderecoVE == msg.sender, "Apenas o VE da sessao pode finalizar.");
        require(sessao.horarioFimReal == 0, "Sessao ja foi finalizada."); // Ainda não finalizada
        require(_energiaConsumida > 0, "Energia consumida deve ser maior que zero.");

        sessao.horarioFimReal = block.timestamp;
        sessao.energiaConsumida = _energiaConsumida;
        sessao.custoTotal = _energiaConsumida * posto.precoPorUnidadeEnergia;

        require(msg.value >= sessao.custoTotal, "Valor enviado insuficiente para cobrir o custo.");

        posto.totalRecebido += sessao.custoTotal; // Acumula no saldo do posto
        sessao.paga = true;
        reserva.ativa = false; // Marca a reserva como utilizada/finalizada

        // Reembolsar valor excedente, se houver
        if (msg.value > sessao.custoTotal) {
            payable(msg.sender).transfer(msg.value - sessao.custoTotal);
        }

        emit RecargaFinalizada(
            _idSessao,
            sessao.idPosto,
            msg.sender,
            sessao.energiaConsumida,
            sessao.custoTotal,
            true
        );
    }

    // Funções de Consulta (Getters públicos já são criados para mappings, mas exemplos abaixo)

    function getDetalhesPosto(uint _idPosto) external view returns (
        uint idPosto,
        address proprietario,
        string memory nome,
        string memory localizacao,
        uint precoPorUnidadeEnergia,
        bool disponivelParaNovasReservas,
        uint256 totalRecebido
    ) {
        Posto storage p = postos[_idPosto];
        return (
            p.idPosto,
            p.proprietario,
            p.nome,
            p.localizacao,
            p.precoPorUnidadeEnergia,
            p.disponivelParaNovasReservas,
            p.totalRecebido
        );
    }

    function getDetalhesReserva(uint _idReserva) external view returns (
        uint idReserva,
        uint idPosto,
        address enderecoVE,
        uint horarioInicio,
        uint horarioFim,
        bool ativa,
        bool sessaoIniciada,
        uint idSessaoVinculada
    ) {
        Reserva storage r = reservas[_idReserva];
        return (
            r.idReserva,
            r.idPosto,
            r.enderecoVE,
            r.horarioInicio,
            r.horarioFim,
            r.ativa,
            r.sessaoIniciada,
            r.idSessaoVinculada
        );
    }

    function getDetalhesSessaoRecarga(uint _idSessao) external view returns (
        uint idSessao,
        uint idReserva,
        address enderecoVE,
        uint idPosto,
        uint horarioInicioReal,
        uint horarioFimReal,
        uint energiaConsumida,
        uint custoTotal,
        bool paga
    ) {
        SessaoRecarga storage s = sessoesRecarga[_idSessao];
        return (
            s.idSessao,
            s.idReserva,
            s.enderecoVE,
            s.idPosto,
            s.horarioInicioReal,
            s.horarioFimReal,
            s.energiaConsumida,
            s.custoTotal,
            s.paga
        );
    }

    function getReservasPorVE(address _enderecoVE) external view returns (uint[] memory) {
        return reservasPorVE[_enderecoVE];
    }

    function getReservasPorPosto(uint _idPosto) external view returns (uint[] memory) {
        return reservasPorPosto[_idPosto];
    }
}