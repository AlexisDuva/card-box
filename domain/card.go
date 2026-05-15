package domain

import "errors"

type Card struct {
	Title string
	Recto string
	Verso string
}

func NewCard(title string, recto string, verso string) (*Card, error) {
	if title == "" {
		return nil, errors.New("NewCard() : nil Title")
	}
	if recto == "" {
		return nil, errors.New("NewCard() : nil Recto")
	}
	if verso == "" {
		return nil, errors.New("NewCard() : nil Verso")
	}
	var c Card
	c = Card{Title: title, Recto: recto, Verso: verso}
	return &c, nil
}
