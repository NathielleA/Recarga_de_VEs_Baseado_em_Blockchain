FROM ethereum/client-go:stable

# Instala git
RUN apk add --no-cache git

# Clona o repositório automaticamente
RUN git clone https://github.com/a1brz/eth-private-network.git /eth

# Define o diretório de dados como volume
VOLUME /root/.ethereum

# Usa os dados da rede local criada pelo eth-private-network
WORKDIR /eth
