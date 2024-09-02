let isShuffling = false; // Declare once at the top

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

let shuffleInterval;

function toggleShuffle() {
    if (isShuffling) {
        stopShuffle();
    } else {
        startShuffle();
    }
}

function openModal(cardName, cardMeaning) {
    // Implement the modal opening logic here
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

document.addEventListener('DOMContentLoaded', function() {
    var modals = document.querySelectorAll('.modal');
    modals.forEach(function(modal) {
        new bootstrap.Modal(modal);
    });
});

document.addEventListener('htmx:afterSwap', (event) => {
    if (event.detail.target.id === "cardModal") {
        var modalElement = document.getElementById('cardModal');
        var modal = new bootstrap.Modal(modalElement);
        modal.show();
    }
});

function openModal(cardName, cardMeaning) {
    // Assume `cardModal` is already in the DOM
    const modalElement = document.getElementById('cardModal');
    const modalTitle = modalElement.querySelector('.modal-title');
    const modalBody = modalElement.querySelector('.modal-body');

    // Update the modal content before showing
    modalTitle.textContent = cardName;
    modalBody.querySelector('.description').textContent = cardMeaning;

    // Show the modal
    var modal = new bootstrap.Modal(modalElement);
    modal.show();
}


function saveSelectedCardToLocalStorage(cardName) {
    let selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    selectedCards.push(cardName);
    localStorage.setItem('selectedCards', JSON.stringify(selectedCards));

    if (selectedCards.length >= 1) {
        hideShuffleButton(); // Hide the shuffle button after selecting the first card
    }

    if (selectedCards.length >= 3) {
        showDoAnotherReadingButton();
    }
}

function hideShuffleButton() {
    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) {
        shuffleButton.style.display = 'none';
    }
}

function showShuffleButton() {
    const shuffleButton = document.getElementById('shuffleButton');
    if (shuffleButton) {
        shuffleButton.style.display = 'block';
    }
}

function loadSelectedCardsFromLocalStorage() {
    const selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    if (selectedCards.length >= 1) {
        hideShuffleButton(); // Hide shuffle button if at least one card is selected
    }
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
            window.location.reload();
            showShuffleButton(); // Re-show the shuffle button after reset
        },
        error: function(xhr, status, error) {
            console.error("Error resetting reading:", status, error);
        }
    });
}

function selectCard(cardName) {
    saveSelectedCardToLocalStorage(cardName);
    let selectedCards = JSON.parse(localStorage.getItem('selectedCards') || '[]');
    if (selectedCards.length >= 3) {
        showDoAnotherReadingButton();
    }
}

function logLocalStorage() {
    console.log("Current localStorage:", JSON.parse(localStorage.getItem('selectedCards') || '[]'));
}
