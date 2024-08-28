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

    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) {
        shuffleButton.addEventListener('click', toggleShuffle);
    }

    const doAnotherReadingButton = document.getElementById('doAnotherReading');
    if (doAnotherReadingButton) {
        doAnotherReadingButton.addEventListener('click', doAnotherReading);
    }

    loadSelectedCardsFromLocalStorage();
});

let isShuffling = false;
let shuffleInterval;

function toggleShuffle() {
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
        shuffleButton.classList.toggle('hover:bg-red-600', isShuffling);
        shuffleButton.classList.toggle('hover:bg-purple-600', !isShuffling);
    }
}

function swapCards(card1, card2) {
    const parent = card1.parentNode;
    const nextSibling = card2.nextSibling;
    const card1Rect = card1.getBoundingClientRect();
    const card2Rect = card2.getBoundingClientRect();

    const deltaX = card2Rect.left - card1Rect.left;

    card1.style.transition = card2.style.transition = 'transform 0.5s ease-in-out';
    card1.style.transform = `translateX(${deltaX}px)`;
    card2.style.transform = `translateX(${-deltaX}px)`;

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

function saveSelectedCardToLocalStorage(cardName) {
    let selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    selectedCards.push(cardName);
    localStorage.setItem('selectedCards', JSON.stringify(selectedCards));

    if (selectedCards.length >= 3) {
        showDoAnotherReadingButton();
    }
}

function loadSelectedCardsFromLocalStorage() {
    const selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    // Here you would typically update the UI to reflect the selected cards
    // This depends on how your server-side rendering works with HTMX
    if (selectedCards.length >= 3) {
        showDoAnotherReadingButton();
    }
}

function showDoAnotherReadingButton() {
    const button = document.getElementById('doAnotherReading');
    if (button) {
        button.style.display = 'block';
    }
}

function doAnotherReading() {
    console.log("Do Another Reading clicked");
    localStorage.removeItem('selectedCards');
    console.log("Local storage cleared");

    // Make an HTMX request to reset the server-side state
    htmx.ajax('GET', '/reset-reading', {
        target: 'body',
        swap: 'innerHTML',
        headers: {
            'HX-Request': 'true'
        },
        success: function() {
            console.log("Reset reading successful");
            // Force a re-render of the page
            location.reload();
        },
        error: function(xhr, status, error) {
            console.error("Error resetting reading:", status, error);
        }
    });
}

// Assuming you have a function that's called when a card is selected
function selectCard(cardName) {
    saveSelectedCardToLocalStorage(cardName);
    // Check if we've selected 3 cards
    let selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    if (selectedCards.length >= 3) {
        showDoAnotherReadingButton();
    }
}

// Add this function to help with debugging
function logLocalStorage() {
    console.log("Current localStorage:", JSON.parse(localStorage.getItem('selectedCards') || '[]'));
}
