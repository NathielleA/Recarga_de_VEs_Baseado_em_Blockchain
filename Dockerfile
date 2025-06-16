# Dockerfile for Hardhat
FROM node:18

WORKDIR /app

# Instala o Hardhat globalmente (opcional)
RUN npm install -g hardhat

# Copia os arquivos do projeto (exceto node_modules se estiver no .dockerignore)
COPY package*.json ./

# Instala as dependências do projeto
RUN npm install

# Instala as dependências extras exigidas pelo hardhat-toolbox
RUN npm install --save-dev ts-node typechain typescript @types/mocha solidity-coverage @typechain/hardhat @typechain/ethers-v6 hardhat-gas-reporter @nomicfoundation/hardhat-verify @nomicfoundation/hardhat-ignition-ethers @nomicfoundation/hardhat-network-helpers

# Copia o restante dos arquivos do projeto
COPY . .

# Expõe a porta padrão usada pelo Hardhat
EXPOSE 8545

# Comando padrão: roda o nó local
CMD ["npx", "hardhat", "node"]
