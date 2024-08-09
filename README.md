# Solitaire Game

## Overview

This project is a command-line implementation of the classic Solitaire card game. The game is written in Go and provides a simple yet engaging way to play Solitaire directly from your terminal.

## Features

- **Classic Solitaire Rules**: Implements the traditional rules of Solitaire.
- **Command-Line Interface**: Play the game directly from your terminal.
- **Randomized Deck**: Each game starts with a shuffled deck for a unique experience every time.
- **Move Validation**: Ensures that all moves are legal according to Solitaire rules.
- **Undo Functionality**: Allows players to undo their last move.

## Installation

To install and run the Solitaire game, you need to have Go installed on your machine. Follow the steps below to get started:

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/yourusername/solitaire.git
    cd solitaire
    ```

2. **Build the Game**:
    ```sh
    go build -o solitaire
    ```

3. **Run the Game**:
    ```sh
    ./solitaire
    ```

## How to Play

### Starting the Game

After running the game, you will see the initial layout of the cards. The game follows the standard Solitaire layout with seven columns of cards, a draw pile, and four foundation piles.

### Commands

- **Draw Card**: To draw a card from the draw pile, use the command `D`.
- **Move Card**: To move a card, use the command `MXX`, where `XX` represents the move details:
  - `M` - Move
  - `FromPile` - The pile number you are moving from
  - `FromPileIndex` - The index of the card in the pile you are moving from
  - `ToPile` - The pile number you are moving to
  - Example: `M12` moves a card from pile 1 to pile 2.
- **Complete Move**: To move a card to the foundation pile, use the command `CX`, where `X` represents the pile number you are moving from:
  - `C` - Complete
  - `FromPile` - The pile number you are moving from
  - Example: `C1` moves a card from pile 1 to the foundation pile.
- **Undo Move**: To undo the last move, use the command `undo`.
- **Quit Game**: To quit the game, use the command `quit`.

### Winning the Game

The objective is to move all cards to the four foundation piles, sorted by suit and in ascending order from Ace to King.

