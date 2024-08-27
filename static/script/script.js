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
