const { expect } = require("chai");
const { ethers, network } = require("hardhat");
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");

describe("ChargingStation", function () {
    let chargingStation, owner, car1, car2;
    const PRICE_PER_SECOND = ethers.parseUnits("10", "gwei"); // 10 Gwei por segundo

    beforeEach(async function () {
        [owner, car1, car2] = await ethers.getSigners();
        const ChargingStation = await ethers.getContractFactory("ChargingStation");
        chargingStation = await ChargingStation.deploy(PRICE_PER_SECOND);
    });

    it("Should be deployed with status Available", async function () {
        expect(await chargingStation.currentStatus()).to.equal(0); // Enum Status.Available
    });

    describe("Reservations", function () {
        it("A car should be able to reserve an available station", async function () {
            await expect(chargingStation.connect(car1).reserve())
                .to.emit(chargingStation, "StationReserved")
                .withArgs(car1.address, anyValue); // Verifica o usuário, mas aceita qualquer valor para o timestamp
            
            expect(await chargingStation.currentStatus()).to.equal(1); // Reserved
            expect(await chargingStation.currentUser()).to.equal(car1.address);
        });

        it("Should fail to reserve if already reserved", async function () {
            await chargingStation.connect(car1).reserve();
            await expect(chargingStation.connect(car2).reserve())
                .to.be.revertedWith("Station is not available");
        });
    });

    describe("Charging and Payment", function () {
        beforeEach(async function () {
            await chargingStation.connect(car1).reserve();
        });

        it("The reserving car should be able to start charging", async function () {
            await expect(chargingStation.connect(car1).startCharge())
                .to.emit(chargingStation, "ChargingStarted");
            expect(await chargingStation.currentStatus()).to.equal(2); // Charging
        });

        it("Should complete charge and calculate cost", async function () {
            await chargingStation.connect(car1).startCharge();

            // Avança o tempo da blockchain em 100 segundos
            await network.provider.send("evm_increaseTime", [100]);
            await network.provider.send("evm_mine");

            await expect(chargingStation.connect(car1).endCharge())
                .to.emit(chargingStation, "ChargingCompleted");

            const chargedTime = await chargingStation.totalChargedTime();
            expect(chargedTime).to.be.gte(100); // gte = Greater Than or Equal To
        });

        it("Should allow payment and transfer funds to the owner", async function () {
            const chargeTime = 100;
            const cost = BigInt(chargeTime) * BigInt(PRICE_PER_SECOND);

            await expect(
                chargingStation.connect(car1).payForCharge(chargeTime, { value: cost })
            ).to.changeEtherBalance(owner, cost);

            await expect(
                chargingStation.connect(car1).payForCharge(chargeTime, { value: cost })
            ).to.emit(chargingStation, "PaymentReceived").withArgs(car1.address, cost);
        });

        it("Should fail payment if amount is insufficient", async function () {
            const chargeTime = 100;
            const cost = BigInt(chargeTime) * BigInt(PRICE_PER_SECOND);
            const insufficientPayment = cost - BigInt(1);

            await expect(
                chargingStation.connect(car1).payForCharge(chargeTime, { value: insufficientPayment })
            ).to.be.revertedWith("Insufficient payment");
        });
    });
});
