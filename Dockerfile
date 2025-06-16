# Dockerfile for Hardhat
FROM node:18

WORKDIR /app

RUN npm install -g hardhat

COPY . .

RUN npm install

CMD ["npx", "hardhat", "node"]
