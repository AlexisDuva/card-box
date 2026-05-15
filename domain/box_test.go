package domain

import (
	"testing"
)

func makeBox(cells []map[string]Card) Box {
	return Box{Id: "1", Title: "test", Cells: cells, Age: 1}
}

func makeCard(title string) Card {
	return Card{Title: title, Recto: "r", Verso: "v"}
}

func TestNewBox(t *testing.T) {
	cells := []map[string]Card{{}, {}}
	tests := []struct {
		name    string
		id      string
		title   string
		cells   []map[string]Card
		age     int
		wantAge int
		wantErr bool
	}{
		{
			name: "valid box",
			id:   "1", title: "b1", cells: cells, age: 3,
			wantAge: 3, wantErr: false,
		},
		{
			name: "age 0 defaults to 1",
			id:   "1", title: "b1", cells: cells, age: 0,
			wantAge: 1, wantErr: false,
		},
		{
			name: "empty id",
			id:   "", title: "b1", cells: cells, age: 1,
			wantErr: true,
		},
		{
			name: "empty title",
			id:   "1", title: "", cells: cells, age: 1,
			wantErr: true,
		},
		{
			name: "negative age",
			id:   "1", title: "b1", cells: cells, age: -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBox(tt.id, tt.title, tt.cells, tt.age)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewBox() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got.Age != tt.wantAge {
				t.Errorf("NewBox() Age = %d, want %d", got.Age, tt.wantAge)
			}
		})
	}
}

func TestAddCard(t *testing.T) {
	tests := []struct {
		name      string
		initial   []map[string]Card
		card      Card
		wantTitle string
	}{
		{
			name:      "add card to empty cell 0",
			initial:   []map[string]Card{{}},
			card:      makeCard("Go"),
			wantTitle: "Go",
		},
		{
			name:      "add card to cell 0 already containing one card",
			initial:   []map[string]Card{{"Existing": makeCard("Existing")}},
			card:      makeCard("New"),
			wantTitle: "New",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := makeBox(tt.initial)
			got := AddCard(box, tt.card)
			if _, ok := got.Cells[0][tt.wantTitle]; !ok {
				t.Errorf("AddCard() card %q not found in cell 0", tt.wantTitle)
			}
		})
	}
}

func TestRightAnswer(t *testing.T) {
	tests := []struct {
		name           string
		cells          []map[string]Card
		idCell         int
		title          string
		wantInNextCell bool
		wantInCurrent  bool
	}{
		{
			name: "card moves to next cell",
			cells: []map[string]Card{
				{"Go": makeCard("Go")},
				{},
			},
			idCell:         0,
			title:          "Go",
			wantInNextCell: true,
			wantInCurrent:  false,
		},
		{
			name: "card in last cell is removed (learned)",
			cells: []map[string]Card{
				{},
				{"Go": makeCard("Go")},
			},
			idCell:         1,
			title:          "Go",
			wantInNextCell: false,
			wantInCurrent:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := makeBox(tt.cells)
			got := RightAnswer(box, tt.idCell, tt.title)

			_, inCurrent := got.Cells[tt.idCell][tt.title]
			if inCurrent != tt.wantInCurrent {
				t.Errorf("RightAnswer() card in cell %d = %v, want %v", tt.idCell, inCurrent, tt.wantInCurrent)
			}

			if tt.idCell < len(got.Cells)-1 {
				_, inNext := got.Cells[tt.idCell+1][tt.title]
				if inNext != tt.wantInNextCell {
					t.Errorf("RightAnswer() card in cell %d = %v, want %v", tt.idCell+1, inNext, tt.wantInNextCell)
				}
			}
		})
	}
}

func TestWrongAnswer(t *testing.T) {
	tests := []struct {
		name          string
		cells         []map[string]Card
		idCell        int
		title         string
		wantInCell0   bool
		wantInCurrent bool
	}{
		{
			name: "card moves back to cell 0",
			cells: []map[string]Card{
				{},
				{},
				{"Go": makeCard("Go")},
			},
			idCell:        2,
			title:         "Go",
			wantInCell0:   true,
			wantInCurrent: false,
		},
		{
			name: "card already in cell 0 stays in cell 0",
			cells: []map[string]Card{
				{"Go": makeCard("Go")},
				{},
			},
			idCell:        0,
			title:         "Go",
			wantInCell0:   true,
			wantInCurrent: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			box := makeBox(tt.cells)
			got := WrongAnswer(box, tt.idCell, tt.title)

			_, inCell0 := got.Cells[0][tt.title]
			if inCell0 != tt.wantInCell0 {
				t.Errorf("WrongAnswer() card in cell 0 = %v, want %v", inCell0, tt.wantInCell0)
			}

			if tt.idCell != 0 {
				_, inCurrent := got.Cells[tt.idCell][tt.title]
				if inCurrent != tt.wantInCurrent {
					t.Errorf("WrongAnswer() card in cell %d = %v, want %v", tt.idCell, inCurrent, tt.wantInCurrent)
				}
			}
		})
	}
}
