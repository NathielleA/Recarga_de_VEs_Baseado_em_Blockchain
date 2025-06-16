const { ethers } = require("hardhat");
const fs = require('fs');
const path = require('path');

async function main() {
  // 1. Obter a conta do deployer
  const [deployer] = await ethers.getSigners();

  console.log("--------------------------------------------------");
  console.log("Iniciando o deploy de contratos...");
  console.log("Realizando deploy com a conta:", deployer.address);

  const provider = ethers.provider;
  const balance = await provider.getBalance(deployer.address);
  console.log("Saldo atual da conta:", balance.toString(), "Wei");
  console.log("--------------------------------------------------");

  // 2. Definir o preço por segundo para o construtor do contrato
  const pricePerSecond = ethers.parseUnits("0.0000001", "ether"); // 100 Gwei
  console.log(`Definindo pricePerSecond para o contrato: ${pricePerSecond.toString()} wei (${ethers.formatEther(pricePerSecond)} ETH)`);

  // 3. Obter o ContractFactory do seu contrato
  const ContractFactory = await ethers.getContractFactory("ChargingStation");

  // 4. Implanta o contrato, passando o argumento para o construtor
  const chargingStation = await ContractFactory.deploy(pricePerSecond);

  // 5. Esperar até que o contrato esteja completamente implantado
  await chargingStation.waitForDeployment();

  console.log("\n--------------------------------------------------");
  console.log("✅ Contrato 'ChargingStation' implantado com sucesso!");
  console.log("➡️ Endereço do Contrato:", chargingStation.target);
  console.log("--------------------------------------------------");

  // 6. Salvar o endereço do contrato em um arquivo
  const contractsDataDir = path.join(__dirname, '..', 'contracts_data');
  const contractAddressFilePath = path.join(contractsDataDir, 'ChargingStationAddress.txt');

  if (!fs.existsSync(contractsDataDir)) {
    fs.mkdirSync(contractsDataDir);
  }

  fs.writeFileSync(contractAddressFilePath, chargingStation.target);
  console.log(`\n➡️ Endereço do contrato salvo em: ${contractAddressFilePath}`);
  console.log("--------------------------------------------------");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
