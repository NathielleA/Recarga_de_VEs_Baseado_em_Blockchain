# Use uma imagem oficial do Geth
FROM ethereum/client-go:stable

# Variáveis de ambiente (opcional, pode ser configurado no script)
ENV DATADIR="/root/.ethereum/node_data"
ENV GENESIS_FILE="/root/geth_config/genesis.json"
ENV PASSWORD_FILE="/root/geth_config/password.txt"

# Instalar jq para ler o chainId do genesis.json no script start_geth.sh
# e bash se o script usar funcionalidades específicas do bash (sh deve ser suficiente para o script acima)
RUN apk add --no-cache jq bash

# Criar diretório de configuração dentro do container
WORKDIR /root/geth_config

# Copiar o arquivo genesis para dentro do container
COPY genesis.json .

# Copiar o arquivo de senha da conta signatária
# IMPORTANTE: Este método de incluir a senha diretamente na imagem
# é adequado apenas para ambientes de teste locais.
# Para produção, use Docker Secrets ou outras formas seguras de gerenciar senhas.
COPY password.txt .

# Copiar o script de inicialização
COPY start_geth.sh /usr/local/bin/start_geth.sh
RUN chmod +x /usr/local/bin/start_geth.sh

# Criar o diretório de dados principal
RUN mkdir -p ${DATADIR} && chown -R root:root /root/.ethereum

# Expor as portas padrão do Geth
EXPOSE 8545  # HTTP-RPC
EXPOSE 8546  # WS-RPC
EXPOSE 8551  # Auth-RPC (Engine API)
EXPOSE 30303 # P2P TCP
EXPOSE 30303/udp # P2P UDP

# Volume para persistir os dados da blockchain fora do container
VOLUME ${DATADIR}

# Ponto de entrada para iniciar o Geth usando o script
ENTRYPOINT ["/usr/local/bin/start_geth.sh"]

# Comando padrão (pode ser usado para passar argumentos adicionais ao Geth via start_geth.sh)
CMD []