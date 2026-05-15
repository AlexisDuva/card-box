package domain

import (
	"errors"
	"fmt"
)

type Box struct {
	Id    string
	Title string
	Cells [][]Card
	Day   int
}

func AddCard(b Box, c Card) Box {
	b.Cells[0] = append(b.Cells[0], c)
	return b
}

func CleanLastCellIfNeeded(b *Box) {
	if NeedCleanLastCell(len((*b).Cells), (*b).Day) {
		(*b).Cells[(len((*b).Cells))-1] = []Card{}
	}
}

func NewBox(id string, title string, cells [][]Card, day int) (Box, error) {
	if id == "" {
		return Box{}, errors.New("newBox() : nil Id")
	}
	if title == "" {
		return Box{}, errors.New("newBox() : nil Title")
	}
	if day < 0 {
		return Box{}, errors.New("newBox() : negative Day")
	}
	var b Box
	if day == 0 {
		b = Box{Id: id, Title: title, Cells: cells, Day: 1}
	} else {
		b = Box{Id: id, Title: title, Cells: cells, Day: day}
	}
	return b, nil
}

func PrintBox(box Box) {
	for i, cell := range box.Cells {
		fmt.Printf("  [%d] (%d cartes)\n", i+1, len(cell))
		if len(cell) == 0 {
			fmt.Println("      —")
		}
		for _, card := range cell {
			fmt.Printf("      • %s\n", card.Title)
		}
		fmt.Println()
	}
}
