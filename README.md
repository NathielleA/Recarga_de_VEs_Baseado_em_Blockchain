# ⚡ Sistema de Recarga de Veículos Elétricos Baseado em Blockchain 🔗


## Resumo e Objetivo

Este projeto propõe o desenvolvimento de um sistema de recarga de veículos elétricos (VEs) utilizando tecnologia blockchain. O objetivo é registrar de forma segura e transparente todas as transações de reserva, recarga e pagamento, promovendo confiança entre usuários e empresas, além de garantir a auditabilidade das operações. A solução visa atender às demandas de um cenário de mobilidade elétrica cada vez mais conectado e colaborativo.

## Arquitetura do Sistema

O sistema é composto por:

- **Rede Ethereum Privada:** Utiliza o consenso Proof of Authority (Clique) para garantir desempenho e controle sobre os participantes.
- **Contratos Inteligentes (Solidity):** Responsáveis pelo registro transparente das operações de reserva, recarga e pagamento.
- **Simulador em Golang:** Scripts em Go automatizam e integram as interações com a blockchain, simulando o comportamento de carros e postos de recarga.
- **Orquestração com Docker:** Permite a simulação de múltiplos participantes e ambientes distribuídos, facilitando testes e desenvolvimento.

### Fluxo Básico

1. **Reserva:** O usuário reserva um ponto de recarga via contrato inteligente.
2. **Recarga:** O processo de recarga é iniciado e monitorado na blockchain.
3. **Pagamento:** O pagamento é registrado, garantindo transparência e auditabilidade.

## Tecnologias Utilizadas

- **Solidity:** Para desenvolvimento dos contratos inteligentes, devido à sua integração nativa com Ethereum.
- **Golang:** Pela eficiência, concorrência e facilidade de integração com APIs Ethereum (via go-ethereum).
- **Docker:** Para orquestração de ambientes, simulação de nós e participantes distribuídos.
- **Hardhat:** Ferramenta para desenvolvimento, teste e deploy de contratos inteligentes.
- **Ethereum (Clique/PoA):** Rede privada para controle e desempenho em ambiente de testes.

Essas tecnologias foram escolhidas por sua robustez, comunidade ativa e alinhamento com os requisitos de transparência, descentralização e auditabilidade.

## Como Executar

### 1. Subir a Blockchain Privada

O projeto utiliza o Hardhat para simular a blockchain localmente. Para iniciar:

```sh
npm install node
docker-compose up --build
```

Isso irá construir a imagem Docker e iniciar um nó Hardhat acessível em `localhost:8545`.

### 2. Executar o Simulador em Golang

Com a blockchain rodando, abra um novo terminal e execute o simulador:

```sh
cd simulador_go
go run .
```

O simulador irá conectar-se à blockchain, criar entidades (carro e posto), e executar o fluxo de reserva, recarga e pagamento, registrando todas as operações na blockchain.

## Estrutura do Projeto

- `blockchain/contracts/RecargaVE.sol`: Contrato inteligente principal.
- `blockchain/contracts/recargave.go`: Bindings Go gerados para interação com o contrato.
- `simulador_go`: Scripts e simulador em Go.
- `Dockerfile` e `docker-compose.yml`: Orquestração do ambiente.
- `deploy`: Scripts para deploy do contrato na rede.

## Considerações Finais

O sistema demonstra a viabilidade de utilizar blockchain para garantir reservas sequenciais, reduzir riscos de interrupções e promover transparência no setor de mobilidade elétrica. A arquitetura modular e o uso de containers facilitam testes, desenvolvimento e futuras expansões do sistema.
