// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract ChargingStation {
    address public owner;
    uint256 public pricePerSecond; // Preço em wei por segundo de recarga

    enum Status { Available, Reserved, Charging, OutOfOrder }
    Status public currentStatus;

    address public currentUser;
    uint256 public reservationTime;
    uint256 public chargeStartTime;
    uint256 public totalChargedTime;

    event StationReserved(address indexed user, uint256 timestamp);
    event ChargingStarted(address indexed user, uint256 timestamp);
    event ChargingCompleted(address indexed user, uint256 chargeTime, uint256 cost);
    event PaymentReceived(address indexed user, uint256 amount);

    constructor(uint256 _pricePerSecond) {
        owner = msg.sender;
        pricePerSecond = _pricePerSecond;
        currentStatus = Status.Available;
    }

    modifier onlyWhenAvailable() {
        require(currentStatus == Status.Available, "Station is not available");
        _;
    }

    modifier onlyReservedBy(address _user) {
        require(currentStatus == Status.Reserved && currentUser == _user, "Not reserved by this user");
        _;
    }
    
    modifier onlyDuringCharging(address _user) {
        require(currentStatus == Status.Charging && currentUser == _user, "Not charging for this user");
        _;
    }

    // Usuário (Carro) chama para reservar a estação
    function reserve() external onlyWhenAvailable {
        currentStatus = Status.Reserved;
        currentUser = msg.sender;
        reservationTime = block.timestamp;
        emit StationReserved(msg.sender, block.timestamp);
    }

    // Usuário (Carro) chama para iniciar a recarga após reservar
    function startCharge() external onlyReservedBy(msg.sender) {
        // Poderia ter uma lógica para expirar a reserva
        currentStatus = Status.Charging;
        chargeStartTime = block.timestamp;
        emit ChargingStarted(msg.sender, block.timestamp);
    }

    // Usuário (Carro) chama para finalizar a recarga
    function endCharge() external onlyDuringCharging(msg.sender) {
        totalChargedTime = block.timestamp - chargeStartTime;
        uint256 cost = totalChargedTime * pricePerSecond;
        
        currentStatus = Status.Available; // Fica disponível para pagamento e nova reserva
        emit ChargingCompleted(msg.sender, totalChargedTime, cost);
        
        // Em um caso real, o pagamento seria o próximo passo obrigatório
        // Aqui, simplificamos e liberamos a estação.
    }

    // Usuário (Carro) chama para pagar
    function payForCharge(uint256 _chargeTime) external payable {
        uint256 requiredAmount = _chargeTime * pricePerSecond;
        require(msg.value >= requiredAmount, "Insufficient payment");

        // Transfere o pagamento para o dono do posto
        payable(owner).transfer(msg.value);

        emit PaymentReceived(msg.sender, msg.value);

        // Zera o estado para o próximo cliente
        currentUser = address(0);
        chargeStartTime = 0;
        totalChargedTime = 0;
    }
}