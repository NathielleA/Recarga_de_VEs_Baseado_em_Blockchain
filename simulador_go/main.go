package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	// Importe o pacote gerado pelo abigen.
	// Substitua "NOME_DO_SEU_MODULO_GO/simulador_go/recargave_contract" pelo caminho correto
	// onde NOME_DO_SEU_MODULO_GO é o nome do seu módulo Go (definido no go.mod).
	// Se RecargaVEAbigen.go estiver no mesmo diretório e com `package main`,
	// você não precisaria de um import separado para ele, mas usar um pacote distinto é mais limpo.
	// Para este exemplo, vamos assumir que você criou um pacote `recargave_contract`:
	"charging-system/blockchain/contracts" // <<< SUBSTITUA ESTE CAMINHO

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// --- Configurações e Placeholders - SUBSTITUA PELOS SEUS VALORES ---
const (
	// URL do seu nó Geth (ajuste se estiver usando Docker Compose, ex: "http://geth_node:8545")
	gethRPCURL = "http://localhost:8545"

	// Endereço do contrato RecargaVE implantado na sua rede Geth
	contractAddressHex = "0xSEU_ENDERECO_DE_CONTRATO_AQUI" // <<< SUBSTITUA

	// Chave privada da conta que será o proprietário do posto (pré-financiada)
	// !!! NUNCA FAÇA HARDCODE DE CHAVES PRIVADAS EM PRODUÇÃO !!!
	postoOwnerPrivateKeyHex = "2307470b0ba60fd8410dcab163c45c513dc1bf44b7b1c1f5a0a94c2ed572c2ac" // <<< SUBSTITUA (ex: "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")

	// Chave privada da conta que será o VE (pré-financiada)
	// !!! NUNCA FAÇA HARDCODE DE CHAVES PRIVADAS EM PRODUÇÃO !!!
	veOwnerPrivateKeyHex = "3874744cc85007e018e4f22b0b66c310793196755fee92e87878aa3ea4c983e0" // <<< SUBSTITUA (ex: "df57089feb5f7fbf800908df002191758743875e5c39e8799d681825e68369f9")

	// Chain ID da sua rede Geth privada (o mesmo do genesis.json)
	chainIDValue = 13370 // <<< SUBSTITUA se o seu ChainID for diferente
)

// --- Fim das Configurações ---

func main() {
	// Conectar ao cliente Ethereum (Geth)
	client, err := ethclient.Dial(gethRPCURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao nó Ethereum: %v", err)
	}
	defer client.Close()
	fmt.Println("Conectado ao nó Ethereum!")

	// Carregar chaves privadas e criar objetos de autenticação para transações
	postoOwnerAuth, postoOwnerAddress := loadSigner(postoOwnerPrivateKeyHex, client)
	veOwnerAuth, veOwnerAddress := loadSigner(veOwnerPrivateKeyHex, client)

	fmt.Printf("Endereço do Dono do Posto: %s\n", postoOwnerAddress.Hex())
	fmt.Printf("Endereço do VE: %s\n", veOwnerAddress.Hex())

	// Carregar o contrato inteligente
	contractAddr := common.HexToAddress(contractAddressHex)
	recargaVEInstance, err := recargave_contract.NewRecargaveContract(contractAddr, client) // Use o nome do pacote/construtor gerado pelo abigen
	if err != nil {
		log.Fatalf("Falha ao instanciar o contrato RecargaVE: %v", err)
	}
	fmt.Printf("Contrato RecargaVE carregado no endereço: %s\n", contractAddressHex)

	// --- 1. Registrar um Posto de Recarga ---
	fmt.Println("\n--- Registrando Posto ---")
	nomePosto := "Posto Rápido Central"
	localizacaoPosto := "Rua Principal, 123"
	precoPorUnidadeEnergia := big.NewInt(1000000000000000) // 0.001 Ether (em Wei) por unidade de energia

	txPosto, err := recargaVEInstance.RegistrarPosto(postoOwnerAuth, nomePosto, localizacaoPosto, precoPorUnidadeEnergia)
	if err != nil {
		log.Fatalf("Falha ao enviar transação RegistrarPosto: %v", err)
	}
	fmt.Printf("Transação RegistrarPosto enviada! Hash: %s\n", txPosto.Hash().Hex())
	fmt.Println("Aguardando mineração da transação RegistrarPosto...")
	receiptPosto, err := bind.WaitMined(context.Background(), client, txPosto)
	if err != nil {
		log.Fatalf("Falha ao minerar transação RegistrarPosto: %v", err)
	}
	if receiptPosto.Status != 1 {
		log.Fatalf("Transação RegistrarPosto falhou! Status: %d", receiptPosto.Status)
	}
	fmt.Println("Transação RegistrarPosto minerada com sucesso!")

	// É uma boa prática obter o ID do posto a partir de um evento ou de uma chamada de contrato
	// Para este exemplo, vamos assumir que o primeiro posto registrado tem ID 1.
	// Em um cenário real, você ouviria o evento PostoRegistrado.
	idPostoRegistrado := big.NewInt(1) // Assumindo que proximoIdPosto começa em 1 no contrato
	fmt.Printf("Posto '%s' registrado (assumindo ID: %s).\n", nomePosto, idPostoRegistrado.String())

	// --- 2. VE Faz uma Reserva ---
	fmt.Println("\n--- VE Fazendo Reserva ---")
	// Horário de início: agora + 5 minutos (em timestamp Unix)
	// Duração: 30 minutos
	horarioInicioReserva := time.Now().Add(5 * time.Minute).Unix()
	duracaoMinutosReserva := big.NewInt(30)

	txReserva, err := recargaVEInstance.FazerReserva(veOwnerAuth, idPostoRegistrado, big.NewInt(horarioInicioReserva), duracaoMinutosReserva)
	if err != nil {
		log.Fatalf("Falha ao enviar transação FazerReserva: %v", err)
	}
	fmt.Printf("Transação FazerReserva enviada! Hash: %s\n", txReserva.Hash().Hex())
	fmt.Println("Aguardando mineração da transação FazerReserva...")
	receiptReserva, err := bind.WaitMined(context.Background(), client, txReserva)
	if err != nil {
		log.Fatalf("Falha ao minerar transação FazerReserva: %v", err)
	}
	if receiptReserva.Status != 1 {
		log.Fatalf("Transação FazerReserva falhou! Status: %d", receiptReserva.Status)
	}
	fmt.Println("Transação FazerReserva minerada com sucesso!")
	// Obter o ID da reserva a partir do evento ReservaFeita (ou assumir para o exemplo)
	// Para este exemplo, vamos assumir que a primeira reserva tem ID 1.
	idReservaFeita := big.NewInt(1) // Assumindo que proximoIdReserva começa em 1 no contrato
	fmt.Printf("Reserva feita para o posto ID %s pelo VE %s (assumindo ID Reserva: %s).\n", idPostoRegistrado.String(), veOwnerAddress.Hex(), idReservaFeita.String())

	// Simular espera até o horário da reserva (ou um pouco depois para testar)
	fmt.Printf("Aguardando horário da reserva (início em: %s)...\n", time.Unix(horarioInicioReserva, 0).Format(time.RFC1123))
	// time.Sleep(time.Duration(horarioInicioReserva-time.Now().Unix()+5) * time.Second) // Espera real
	time.Sleep(5 * time.Second) // Pequena pausa para simulação

	// --- 3. VE Inicia a Recarga ---
	fmt.Println("\n--- VE Iniciando Recarga ---")
	txIniciar, err := recargaVEInstance.IniciarRecarga(veOwnerAuth, idReservaFeita)
	if err != nil {
		log.Fatalf("Falha ao enviar transação IniciarRecarga: %v", err)
	}
	fmt.Printf("Transação IniciarRecarga enviada! Hash: %s\n", txIniciar.Hash().Hex())
	fmt.Println("Aguardando mineração da transação IniciarRecarga...")
	receiptIniciar, err := bind.WaitMined(context.Background(), client, txIniciar)
	if err != nil {
		log.Fatalf("Falha ao minerar transação IniciarRecarga: %v", err)
	}
	if receiptIniciar.Status != 1 {
		log.Fatalf("Transação IniciarRecarga falhou! Status: %d", receiptIniciar.Status)
	}
	fmt.Println("Transação IniciarRecarga minerada com sucesso!")
	// Obter o ID da sessão a partir do evento RecargaIniciada (ou assumir para o exemplo)
	idSessaoIniciada := big.NewInt(1) // Assumindo que proximoIdSessao começa em 1
	fmt.Printf("Recarga iniciada para reserva ID %s (assumindo ID Sessão: %s).\n", idReservaFeita.String(), idSessaoIniciada.String())

	// Simular tempo de recarga
	fmt.Println("Simulando tempo de recarga por 10 segundos...")
	time.Sleep(10 * time.Second)

	// --- 4. VE Finaliza a Recarga e Paga ---
	fmt.Println("\n--- VE Finalizando Recarga ---")
	energiaConsumida := big.NewInt(25) // Ex: 25 unidades de energia

	// Calcular custo: energiaConsumida * precoPorUnidadeEnergia
	// Precisamos obter o preço do posto do contrato ou armazená-lo
	// Para simplificar, estamos usando o mesmo precoPorUnidadeEnergia usado no registro.
	custoTotal := new(big.Int).Mul(energiaConsumida, precoPorUnidadeEnergia)
	fmt.Printf("Energia consumida: %s unidades, Preço por unidade: %s Wei, Custo total calculado: %s Wei\n",
		energiaConsumida.String(), precoPorUnidadeEnergia.String(), custoTotal.String())

	// Para transações pagáveis, precisamos definir o campo Value no TransactOpts
	veOwnerAuthPagamento := veOwnerAuth     // Copiar para não modificar o original
	veOwnerAuthPagamento.Value = custoTotal // Definir o valor a ser enviado com a transação

	txFinalizar, err := recargaVEInstance.FinalizarRecarga(veOwnerAuthPagamento, idSessaoIniciada, energiaConsumida)
	if err != nil {
		log.Fatalf("Falha ao enviar transação FinalizarRecarga: %v", err)
	}
	fmt.Printf("Transação FinalizarRecarga enviada! Hash: %s\n", txFinalizar.Hash().Hex())
	fmt.Println("Aguardando mineração da transação FinalizarRecarga...")
	receiptFinalizar, err := bind.WaitMined(context.Background(), client, txFinalizar)
	if err != nil {
		log.Fatalf("Falha ao minerar transação FinalizarRecarga: %v", err)
	}
	if receiptFinalizar.Status != 1 {
		log.Fatalf("Transação FinalizarRecarga falhou! Status: %d", receiptFinalizar.Status)
	}
	fmt.Println("Transação FinalizarRecarga minerada com sucesso! Pagamento efetuado.")

	fmt.Println("\nSimulação básica concluída!")

	// Aqui você poderia adicionar lógica para ouvir eventos do contrato
	// Exemplo de como obter saldo de uma conta:
	balance, err := client.BalanceAt(context.Background(), postoOwnerAddress, nil)
	if err != nil {
		log.Printf("Falha ao obter saldo do dono do posto: %v", err)
	} else {
		fmt.Printf("Saldo atual do Dono do Posto (%s): %s Wei\n", postoOwnerAddress.Hex(), balance.String())
	}

	balanceVE, err := client.BalanceAt(context.Background(), veOwnerAddress, nil)
	if err != nil {
		log.Printf("Falha ao obter saldo do VE: %v", err)
	} else {
		fmt.Printf("Saldo atual do VE (%s): %s Wei\n", veOwnerAddress.Hex(), balanceVE.String())
	}
}

// loadSigner carrega uma chave privada e retorna um objeto TransactOpts e o endereço da conta.
// !!! ESTA FUNÇÃO USA CHAVE PRIVADA HARDCODED E É INSEGURA PARA PRODUÇÃO !!!
func loadSigner(privateKeyHex string, client *ethclient.Client) (*bind.TransactOpts, common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatalf("Erro ao carregar chave privada: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Erro ao converter chave pública para ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Obter nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("Erro ao obter nonce pendente: %v", err)
	}

	// Obter preço do gás sugerido
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Erro ao sugerir preço do gás: %v", err)
	}

	chainID := big.NewInt(chainIDValue)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Erro ao criar transator com chave: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // Valor em Wei a ser enviado com a transação (padrão 0)
	auth.GasLimit = uint64(3000000) // Limite de gás (ajuste conforme necessário)
	auth.GasPrice = gasPrice

	return auth, fromAddress
}

// Para rodar este código:
// 1. Certifique-se que seu nó Geth está rodando e acessível na `gethRPCURL`.
// 2. Certifique-se que o contrato `RecargaVE.sol` está implantado e o `contractAddressHex` está correto.
// 3. Certifique-se que as chaves privadas `postoOwnerPrivateKeyHex` e `veOwnerPrivateKeyHex` são válidas,
//    correspondem a contas com fundos na sua rede Geth, e estão corretas.
// 4. Gere o binding Go para seu contrato com `abigen` e certifique-se que o import para
//    `recargave_contract` (ou o nome que você deu) está correto.
// 5. Execute `go mod init seu_projeto_go` (se ainda não o fez) e `go mod tidy` para baixar as dependências.
// 6. Execute `go run main.go`.
