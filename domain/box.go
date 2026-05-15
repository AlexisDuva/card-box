package domain

import (
	"errors"
	"fmt"
	"strings"
)

type Box struct {
	Id    string
	Title string
	Cells []map[string]Card
	Age   int
}

func AddCard(b Box, c Card) Box {
	b.Cells[0][c.Title] = c
	return b
}

func RightAnswer(b Box, idCell int, title string) Box {
	if idCell < len(b.Cells)-1 {
		b.Cells[idCell+1][title] = b.Cells[idCell][title]
	}
	delete(b.Cells[idCell], title)
	return b
}

func WrongAnswer(b Box, idCell int, title string) Box {
	if idCell == 0 {
		return b
	}
	b.Cells[0][title] = b.Cells[idCell][title]
	delete(b.Cells[idCell], title)
	return b
}

func NewBox(id string, title string, cells []map[string]Card, age int) (Box, error) {
	if id == "" {
		return Box{}, errors.New("newBox() : nil Id")
	}
	if title == "" {
		return Box{}, errors.New("newBox() : nil Title")
	}
	if age < 0 {
		return Box{}, errors.New("newBox() : negative Age")
	}
	var b Box
	if age == 0 {
		b = Box{Id: id, Title: title, Cells: cells, Age: 1}
	} else {
		b = Box{Id: id, Title: title, Cells: cells, Age: age}
	}
	return b, nil
}

func (b Box) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Box : %s, Age : %d\n", b.Title, b.Age)
	for i, cell := range b.Cells {
		fmt.Fprintf(&sb, "  [%d] (%d cartes)\n", i+1, len(cell))
		if len(cell) == 0 {
			sb.WriteString("      —\n")
		}
		for _, card := range cell {
			fmt.Fprintf(&sb, "      • %s\n", card.Title)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
