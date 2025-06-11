// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title RecargaVESimplificado
 * @dev Contrato simplificado para agendamento e registro de recargas de VEs.
 * O custo é registrado on-chain, mas o pagamento é gerenciado off-chain.
 */
contract RecargaVESimplificado {

    // Enum para um controle de status mais claro e eficiente
    enum StatusSessao { RESERVADA, EM_ANDAMENTO, FINALIZADA, CANCELADA }

    struct Posto {
        uint idPosto;
        address proprietario;
        string nome;
        string localizacao;
        uint precoPorUnidadeEnergia; // Em uma unidade base (ex: centavos)
    }

    // Estrutura unificada para agendamento e recarga
    struct Sessao {
        uint idSessao;
        uint idPosto;
        address enderecoVE;
        uint horarioAgendadoInicio;
        uint horarioAgendadoFim;
        uint horarioRealInicio;
        uint energiaConsumida;
        uint custoTotal;
        StatusSessao status;
    }

    // Contadores
    uint public proximoIdPosto = 1;
    uint public proximoIdSessao = 1;

    // Mapeamentos
    mapping(uint => Posto) public postos;
    mapping(uint => Sessao) public sessoes;
    mapping(uint => uint[]) public sessoesPorPosto; // Para verificar conflitos

    // Eventos
    event PostoRegistrado(uint indexed idPosto, address indexed proprietario, string nome);
    event SessaoAgendada(uint indexed idSessao, uint indexed idPosto, address indexed enderecoVE, uint horarioInicio, uint horarioFim);
    event RecargaIniciada(uint indexed idSessao);
    event RecargaFinalizada(uint indexed idSessao, uint energiaConsumida, uint custoTotal);
    event SessaoCancelada(uint indexed idSessao);

    // Modificadores
    modifier apenasProprietarioPosto(uint _idPosto) {
        require(postos[_idPosto].proprietario == msg.sender, "Apenas o proprietario do posto.");
        _;
    }

    function registrarPosto(string calldata _nome, string calldata _localizacao, uint _precoPorUnidadeEnergia) external {
        uint idPosto = proximoIdPosto++;
        postos[idPosto] = Posto({
            idPosto: idPosto,
            proprietario: msg.sender,
            nome: _nome,
            localizacao: _localizacao,
            precoPorUnidadeEnergia: _precoPorUnidadeEnergia
        });
        emit PostoRegistrado(idPosto, msg.sender, _nome);
    }

    function agendarSessao(uint _idPosto, uint _horarioInicio, uint _duracaoMinutos) external {
        require(postos[_idPosto].idPosto != 0, "Posto nao existe.");
        uint _horarioFim = _horarioInicio + (_duracaoMinutos * 1 minutes);
        require(_horarioFim > _horarioInicio, "Duracao invalida.");

        // Lógica para verificar conflito de horário (mantida por ser essencial)
        uint[] storage idsSessoesDoPosto = sessoesPorPosto[_idPosto];
        for (uint i = 0; i < idsSessoesDoPosto.length; i++) {
            Sessao storage sessaoExistente = sessoes[idsSessoesDoPosto[i]];
            if (sessaoExistente.status == StatusSessao.RESERVADA || sessaoExistente.status == StatusSessao.EM_ANDAMENTO) {
                bool sobrepoe = (_horarioInicio < sessaoExistente.horarioAgendadoFim) && (_horarioFim > sessaoExistente.horarioAgendadoInicio);
                require(!sobrepoe, "Horario em conflito.");
            }
        }

        uint idSessao = proximoIdSessao++;
        sessoes[idSessao] = Sessao({
            idSessao: idSessao,
            idPosto: _idPosto,
            enderecoVE: msg.sender,
            horarioAgendadoInicio: _horarioInicio,
            horarioAgendadoFim: _horarioFim,
            horarioRealInicio: 0,
            energiaConsumida: 0,
            custoTotal: 0,
            status: StatusSessao.RESERVADA
        });
        
        sessoesPorPosto[_idPosto].push(idSessao);
        emit SessaoAgendada(idSessao, _idPosto, msg.sender, _horarioInicio, _horarioFim);
    }

    function iniciarRecarga(uint _idSessao) external {
        Sessao storage sessao = sessoes[_idSessao];
        require(sessao.enderecoVE == msg.sender, "Nao autorizado.");
        require(sessao.status == StatusSessao.RESERVADA, "Sessao nao esta reservada.");
        require(block.timestamp >= sessao.horarioAgendadoInicio && block.timestamp <= sessao.horarioAgendadoFim, "Fora do horario agendado.");

        sessao.status = StatusSessao.EM_ANDAMENTO;
        sessao.horarioRealInicio = block.timestamp;
        emit RecargaIniciada(_idSessao);
    }

    // Note que esta função NÃO é 'payable'
    function finalizarRecarga(uint _idSessao, uint _energiaConsumida) external {
        Sessao storage sessao = sessoes[_idSessao];
        Posto storage posto = postos[sessao.idPosto];
        require(sessao.enderecoVE == msg.sender, "Nao autorizado.");
        require(sessao.status == StatusSessao.EM_ANDAMENTO, "Recarga nao iniciada.");

        sessao.energiaConsumida = _energiaConsumida;
        sessao.custoTotal = _energiaConsumida * posto.precoPorUnidadeEnergia;
        sessao.status = StatusSessao.FINALIZADA;
        
        emit RecargaFinalizada(_idSessao, sessao.energiaConsumida, sessao.custoTotal);
    }
    
    function cancelarSessao(uint _idSessao) external {
        Sessao storage sessao = sessoes[_idSessao];
        require(sessao.enderecoVE == msg.sender, "Nao autorizado.");
        require(sessao.status == StatusSessao.RESERVADA, "So pode cancelar sessoes reservadas.");

        sessao.status = StatusSessao.CANCELADA;
        emit SessaoCancelada(_idSessao);
    }

    // Função de consulta para obter detalhes de uma sessão
    function getSessao(uint _idSessao) external view returns (Sessao memory) {
        return sessoes[_idSessao];
    }
}