// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EVCharging {
    struct Reserva {
        uint id;
        address usuario;
        string veiculoId;
        string pontoId;
        string inicio;
        string fim;
        uint preco;
    }

    uint public contador = 0;
    mapping(uint => Reserva) public reservas;

    event NovaReserva(uint id, address usuario);

    function criarReserva(string memory veiculoId, string memory pontoId, string memory inicio, string memory fim, uint preco) public {
        reservas[contador] = Reserva(contador, msg.sender, veiculoId, pontoId, inicio, fim, preco);
        emit NovaReserva(contador, msg.sender);
        contador++;
    }

    function getReserva(uint id) public view returns (Reserva memory) {
        return reservas[id];
    }
}
