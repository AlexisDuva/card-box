package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlexisDuva/card-box/domain"
)

func dataPath() string {
	if p := os.Getenv("CARDBOX_DATA"); p != "" {
		return p
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(home, "card-box", "box.txt")
}

func saveBox(box domain.Box, path string) error {
	data, err := json.Marshal(box)
	if err != nil {
		return fmt.Errorf("json.Marshal() : %s", err)
	}
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Errorf("os.WriteFile() : %s", err)
	}
	return nil
}

func loadData(path string) (domain.Box, error) {
	read, err := os.ReadFile(path)
	if err != nil {
		return domain.Box{}, fmt.Errorf("os.ReadFile() : %s", err)
	}
	var boxRead domain.Box
	if len(read) > 0 {
		err = json.Unmarshal(read, &boxRead)
	} else {
		cells := []map[string]domain.Card{{}, {}, {}, {}, {}, {}, {}}
		boxRead, err = domain.NewBox("1", "b1", cells, 0)
	}

	if err != nil {
		return domain.Box{}, fmt.Errorf("json.Unmarshal() : %s", err)
	}
	return boxRead, nil
}

func main() {
	fmt.Print("Welcome to CardBox\n")
	path := dataPath()
	box, err := loadData(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	readLine := func() string {
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	run := true
	for run {
		fmt.Println("1. Test\n2. Add Card\n3. Print Box\n4. Save and Quit")
		fmt.Println("What do you want to do ?")
		input := readLine()
		switch input {
		case "1":
			cellsToAssess := domain.Assessment(len(box.Cells), box.Age)
			fmt.Printf("cellsToAssess : %d\n", cellsToAssess)
			for _, id := range cellsToAssess {
				for _, card := range box.Cells[id] {
					fmt.Printf("Card : %s \n Question: %s\n", card.Title, card.Recto)
					fmt.Println("Press Enter to see answer")
					readLine()
					fmt.Printf("Answer: %s\n", card.Verso)
				answerLoop:
					for {
						fmt.Println("Was your answer correct ? y/n")
						input = readLine()
						switch input {
						case "y":
							box = domain.RightAnswer(box, id, card.Title)
							break answerLoop
						case "n":
							box = domain.WrongAnswer(box, id, card.Title)
							break answerLoop
						default:
							fmt.Println("Invalid answer. Please try again.")
						}
					}
				}
			}
			box.Age++
		case "2":
			fmt.Println("Title:")
			title := readLine()
			fmt.Println("Recto:")
			recto := readLine()
			fmt.Println("Verso:")
			verso := readLine()
			card, err := domain.NewCard(title, recto, verso)
			if err != nil {
				log.Fatalf("Option 2 : %s", err)
			}
			box = domain.AddCard(box, card)
		case "3":
			fmt.Print(box)
		case "4":
			run = false
		}
	}

	err = saveBox(box, path)
	if err != nil {
		log.Fatal(err)
	}
}
