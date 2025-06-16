const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
const { parseUnits } = require("ethers");

module.exports = buildModule("ChargingStationModule", (m) => {
  // Preço: 10 Gwei por segundo
  const pricePerSecond = m.getParameter("price", parseUnits("10", "gwei"));

  const station = m.contract("ChargingStation", [pricePerSecond]);

  return { station };
});