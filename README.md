# card-box

A CLI flashcard app based on the [Leitner system](https://en.wikipedia.org/wiki/Leitner_system) — a spaced repetition method that schedules cards for review based on how well you know them.

## How it works

Cards are organized into 7 cells. Each cell has a review cycle that doubles in length:

| Cell | Reviewed every |
| ---- | -------------- |
| 1    | 1 day          |
| 2    | 2 days         |
| 3    | 4 days         |
| ...  | ...            |
| 7    | 64 days        |

- **Correct answer** → card moves to the next cell
- **Wrong answer** → card goes back to cell 1
- A card that clears the last cell is considered learned and removed

## Installation

```bash
git clone https://github.com/AlexisDuva/card-box.git
cd card-box
go build ./...
```

## Usage

```bash
go run ./...
```

Data is saved to `~/card-box/box.txt`. Override with the `CARDBOX_DATA` environment variable:

```bash
CARDBOX_DATA=/custom/path/box.txt go run ./...
```

### Menu

| Key | Action                       |
| --- | ---------------------------- |
| `t` | Run a test session for today |
| `a` | Add a new card               |
| `p` | Print the box                |
| `q` | Quit                         |

## Running tests

```bash
go test ./...
```
