#!/bin/sh
set -e

# Variáveis de Configuração
DATADIR="/root/.ethereum/node_data"
GENESIS_FILE="/root/geth_config/genesis.json"
NETWORK_ID=13370
SIGNER_ADDRESS="0x65820a991dd674e7f46f9c0882830b5c2582d3f5"
PASSWORD_FILE="/root/geth_config/password.txt"
HTTP_PORT="8545"
WS_PORT="8546"
AUTH_RPC_PORT="8551"
P2P_PORT="30303"

# Inicializar o nó e copiar a conta do minerador, se necessário
if [ ! -d "$DATADIR/geth" ]; then
  echo "Inicializando nó Geth com genesis: $GENESIS_FILE"
  geth --datadir "$DATADIR" init "$GENESIS_FILE"

  echo "Copiando a conta do minerador para o diretório de dados..."
  mkdir -p "$DATADIR/keystore"
  cp /root/geth_config/keystore/* "$DATADIR/keystore/"
else
  echo "Diretório de dados Geth já inicializado."
fi

echo "Iniciando nó Geth..."
JWT_SECRET_FILE="/root/.ethereum/jwt.hex"

exec geth \
  --datadir "$DATADIR" \
  --networkid "$NETWORK_ID" \
  --syncmode "full" \
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
  --miner.etherbase="$SIGNER_ADDRESS" \
  --verbosity 3 \
  "$@"