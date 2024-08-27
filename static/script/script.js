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
        shuffleButton.addEventListener('click', toggleShuffle);
    }
});

let isShuffling = false;
let shuffleInterval;

function toggleShuffle() {
    const shuffleButton = document.getElementById('shuffleButton');
    if (!shuffleButton) return;

    if (isShuffling) {
        stopShuffle();
    } else {
        startShuffle();
    }
}

function startShuffle() {
    const cardContainer = document.getElementById('shuffled-cards');
    if (!cardContainer) return;

    const cards = Array.from(cardContainer.children);
    if (cards.length === 0) return;

    isShuffling = true;
    updateShuffleButton(true);

    shuffleInterval = setInterval(() => {
        const i = Math.floor(Math.random() * (cards.length - 1)) + 1;
        const j = Math.floor(Math.random() * (i + 1));
        swapCards(cards[i], cards[j]);
    }, 500);
}

function stopShuffle() {
    clearInterval(shuffleInterval);
    isShuffling = false;
    updateShuffleButton(false);
}

function updateShuffleButton(isShuffling) {
    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) {
        shuffleButton.textContent = isShuffling ? 'Stop Shuffle' : 'Shuffle Cards';
        shuffleButton.classList.toggle('bg-red-500', isShuffling);
        shuffleButton.classList.toggle('bg-purple-500', !isShuffling);
    }
}

function swapCards(card1, card2) {
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
    }, 500);
}
