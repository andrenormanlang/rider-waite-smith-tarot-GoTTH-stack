document.addEventListener("DOMContentLoaded", function() {
    const cards = document.querySelectorAll('.card');

    function revealCard(card, delay) {
        setTimeout(() => {
            card.classList.add('flip');
        }, delay);
    }

    cards.forEach((card, index) => {
        revealCard(card, index * 500);  // Adjust delay between reveals
    });
});

function populateModal(cardString) {
    // Parse the stringified card object
    const card = JSON.parse(cardString);

    // Set the content of the modal
    document.getElementById('cardModalLabel').textContent = card.Name;
    document.getElementById('modal-desc').innerHTML = `
        <p class="text-sm text-white flex-grow overflow-y-auto">${card.MeaningUp}</p>
        <p class="text-sm text-white flex-grow overflow-y-auto">${card.Description}</p>
    `;
}

