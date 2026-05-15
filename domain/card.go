package domain

import "errors"

type Card struct {
	Title string
	Recto string
	Verso string
}

func NewCard(title string, recto string, verso string) (Card, error) {
	if title == "" {
		return Card{}, errors.New("NewCard() : nil Title")
	}
	if recto == "" {
		return Card{}, errors.New("NewCard() : nil Recto")
	}
	if verso == "" {
		return Card{}, errors.New("NewCard() : nil Verso")
	}
	return Card{Title: title, Recto: recto, Verso: verso}, nil
}
