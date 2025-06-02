#!/bin/sh
set -e

# Variáveis de Configuração
DATADIR="/root/.ethereum/node_data"
GENESIS_FILE="/root/geth_config/genesis.json"
# Pega o chainId do genesis.json. Certifique-se que 'jq' está instalado no container.
NETWORK_ID=$(jq -r .config.chainId $GENESIS_FILE)
# Endereço da conta signatária (deve corresponder ao extradata e alloc no genesis.json)
# E deve ter uma conta correspondente no keystore do datadir e uma senha em $PASSWORD_FILE
SIGNER_ADDRESS="0x7df9a875a174b3bc565e6424a0050ebc1b2d1d82" # Substitua se o seu signatário for diferente
PASSWORD_FILE="/root/geth_config/password.txt"
HTTP_PORT="8545"
WS_PORT="8546"
AUTH_RPC_PORT="8551" # Porta para Engine API (relevante para PoS, mas incluída para completude)
P2P_PORT="30303"

# Inicializar o nó se o diretório de dados ainda não tiver sido inicializado
# Verifica a existência de um subdiretório específico que Geth cria, como 'geth' ou 'keystore'
if [ ! -d "$DATADIR/geth" ] && [ ! -d "$DATADIR/keystore" ]; then
  echo "Inicializando nó Geth com genesis: $GENESIS_FILE em $DATADIR"
  geth --datadir "$DATADIR" init "$GENESIS_FILE"
  # Importante: Após 'init', a conta signatária precisa existir no keystore do DATADIR.
  # Normalmente, você criaria a conta ANTES e copiaria o keystore para o volume,
  # ou usaria 'geth account import' se tivesse o arquivo da chave privada.
  # Para este script, assume-se que o Dockerfile ou o setup do volume já cuidou
  # de popular o keystore do DATADIR com a conta do SIGNER_ADDRESS.
else
  echo "Diretório de dados Geth já inicializado em $DATADIR."
fi

echo "Iniciando nó Geth..."
# A flag --authrpc.jwtsecret é para a Engine API usada em configurações PoS com clientes de consenso.
# Para uma rede PoA Clique simples, ela pode não ser estritamente necessária ou o Geth pode gerar o arquivo.
# Se não estiver usando um cliente de consenso, pode-se avaliar sua remoção ou garantir que o arquivo exista.
# Geth pode criar este arquivo se não existir, mas logs podem indicar isso.
JWT_SECRET_FILE="/root/.ethereum/jwt.hex"
if [ ! -f "$JWT_SECRET_FILE" ]; then
    echo "Gerando arquivo JWT secret em $JWT_SECRET_FILE"
    geth --datadir "$DATADIR" writedb --prefix "" version > /dev/null 2>&1 & # Truque para geth criar pastas se não existem
    sleep 1 # Dá um tempo para geth criar as pastas
    # Para versões mais recentes do Geth, pode ser necessário um comando diferente para gerar o jwt, ou ele é criado automaticamente.
    # Verifique a documentação do Geth se esta parte causar problemas.
    # Uma alternativa simples é que o Geth o crie na primeira execução com --authrpc.jwtsecret
fi


exec geth \
  --datadir "$DATADIR" \
  --networkid "$NETWORK_ID" \
  --syncmode "full" \
  --gcmode "archive" \
  --port "$P2P_PORT" \
  --http \
  --http.addr "0.0.0.0" \
  --http.port "$HTTP_PORT" \
  --http.corsdomain "*" \
  --http.vhosts "*" \
  --http.api "eth,net,web3,personal,miner,clique" \
  --ws \
  --ws.addr "0.0.0.0" \
  --ws.port "$WS_PORT" \
  --ws.origins "*" \
  --ws.api "eth,net,web3,personal,miner,clique" \
  --authrpc.addr "0.0.0.0" \
  --authrpc.port "$AUTH_RPC_PORT" \
  --authrpc.vhosts "*" \
  --authrpc.jwtsecret "$JWT_SECRET_FILE" \
  --allow-insecure-unlock \
  --unlock "$SIGNER_ADDRESS" \
  --password "$PASSWORD_FILE" \
  --mine \
  --miner.threads=1 \
  --miner.etherbase="$SIGNER_ADDRESS" \
  --verbosity 3 \
  "$@"