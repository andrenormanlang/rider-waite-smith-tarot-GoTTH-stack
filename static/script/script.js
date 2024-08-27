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

document.addEventListener("DOMContentLoaded", function() {
    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) {
        shuffleButton.addEventListener('click', startShuffle);
    }
});

function startShuffle() {
    const cardContainer = document.getElementById('shuffled-cards');
    if (!cardContainer) return;

    const cards = Array.from(cardContainer.children);
    if (cards.length === 0) return;

    // Disable the shuffle button during animation
    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) shuffleButton.disabled = true;

    fisherYatesShuffle(cards).then(() => {
        // Re-enable the shuffle button after animation completes
        if (shuffleButton) shuffleButton.disabled = false;
    });
}

async function fisherYatesShuffle(cards) {
    for (let i = cards.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        await swapCards(cards[i], cards[j]);
    }
}

function swapCards(card1, card2) {
    return new Promise(resolve => {
        const parent = card1.parentNode;
        const nextSibling = card2.nextSibling;
        const card1Rect = card1.getBoundingClientRect();
        const card2Rect = card2.getBoundingClientRect();

        // Calculate the distance to move
        const deltaX = card2Rect.left - card1Rect.left;

        // Animate the swap
        card1.style.transition = card2.style.transition = 'transform 0.5s ease-in-out';
        card1.style.transform = `translateX(${deltaX}px)`;
        card2.style.transform = `translateX(${-deltaX}px)`;

        // After animation, swap the actual DOM elements
        setTimeout(() => {
            card1.style.transition = card2.style.transition = '';
            card1.style.transform = card2.style.transform = '';
            parent.insertBefore(card2, card1);
            if (nextSibling) {
                parent.insertBefore(card1, nextSibling);
            } else {
                parent.appendChild(card1);
            }
            resolve();
        }, 500);
    });
}