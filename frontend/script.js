const API_URL = 'http://localhost:8080'; // Altere se seu backend usar outra porta

document.getElementById('reservaForm').addEventListener('submit', async (e) => {
  e.preventDefault();

  const dados = {
    veiculoId: document.getElementById('veiculoId').value,
    pontoId: document.getElementById('pontoId').value,
    inicio: document.getElementById('inicio').value,
    fim: document.getElementById('fim').value,
    preco: document.getElementById('preco').value
  };

  const res = await fetch(`${API_URL}/reservas`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(dados)
  });

  const resultado = await res.json();
  alert(`Reserva criada. Hash: ${resultado.txHash || resultado.mensagem}`);
  carregarReservas();
});

async function carregarReservas() {
  const res = await fetch(`${API_URL}/reservas`);
  const reservas = await res.json();

  const div = document.getElementById('listaReservas');
  div.innerHTML = '';

  reservas.forEach(r => {
    const item = document.createElement('div');
    item.innerText = `Veículo ${r.veiculoId} -> Ponto ${r.pontoId} | ${r.inicio} até ${r.fim} | R$ ${r.preco}`;
    div.appendChild(item);
  });
}

carregarReservas();
