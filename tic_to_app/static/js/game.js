// static/js/game.js
document.addEventListener('DOMContentLoaded', () => {
    const board = document.getElementById('game-board');
    const cells = document.querySelectorAll('.cell');
    const status = document.getElementById('status');
    const urlParams = new URLSearchParams(window.location.search);
    const gameMode = urlParams.get('mode');
    
    let currentPlayer = 'X';
    let gameActive = true;

    cells.forEach(cell => {
        cell.addEventListener('click', () => handleCellClick(cell));
    });

    async function handleCellClick(cell) {
        if (!gameActive || cell.textContent !== '') return;

        cell.textContent = currentPlayer;
        
        if (checkWinner()) {
            status.textContent = `Player ${currentPlayer} wins!`;
            gameActive = false;
            return;
        }

        if ([...cells].every(cell => cell.textContent !== '')) {
            status.textContent = "It's a draw!";
            gameActive = false;
            return;
        }

        if (gameMode === 'single' && currentPlayer === 'X') {
            currentPlayer = 'O';
            makeComputerMove();
        } else {
            currentPlayer = currentPlayer === 'X' ? 'O' : 'X';
            status.textContent = `Player ${currentPlayer}'s turn`;
        }
    }

    function makeComputerMove() {
        if (!gameActive) return;

        const emptyCells = [...cells].filter(cell => cell.textContent === '');
        if (emptyCells.length > 0) {
            setTimeout(() => {
                const randomCell = emptyCells[Math.floor(Math.random() * emptyCells.length)];
                randomCell.click();
            }, 500);
        }
    }

    function checkWinner() {
        const winPatterns = [
            [0, 1, 2], [3, 4, 5], [6, 7, 8], // Rows
            [0, 3, 6], [1, 4, 7], [2, 5, 8], // Columns
            [0, 4, 8], [2, 4, 6] // Diagonals
        ];

        return winPatterns.some(pattern => {
            const [a, b, c] = pattern;
            return cells[a].textContent &&
                   cells[a].textContent === cells[b].textContent &&
                   cells[a].textContent === cells[c].textContent;
        });
    }
});

