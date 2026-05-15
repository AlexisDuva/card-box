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

func loadData(path string) (domain.Box, error) {
	dirs, err := os.ReadDir(filepath.Dir(path))
	if os.IsNotExist(err) {
		fmt.Printf("User data folder %s not found.\nCreating folder...\n", filepath.Dir(path))
		if err := os.Mkdir(filepath.Dir(path), os.ModeDir); err != nil {
			log.Fatal(err)
		}
		fmt.Println("User data folder successfully created !")
	} else if err != nil {
		log.Fatal(err)
	}
	dataFileExists := false
	for _, d := range dirs {
		if d.Name() == "box.txt" && d.Type().IsRegular() {
			dataFileExists = true
			break
		}
	}
	if !dataFileExists {
		fmt.Printf("Box data file %s not found.\nCreating file...\n", path)
		if _, err := os.Create(path); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Box data file successfully created !")
	}
	read, err := os.ReadFile(path)
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

func autosave(box domain.Box, path string) {
	if err := saveBox(box, path); err != nil {
		fmt.Printf("Save error: %s\n", err)
	}
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
			autosave(box, path)
		case "2":
			fmt.Println("Title:")
			title := readLine()
			fmt.Println("Recto:")
			recto := readLine()
			fmt.Println("Verso:")
			verso := readLine()
			card, err := domain.NewCard(title, recto, verso)
			if err != nil {
				log.Printf("Option 2 : %s", err)
				continue
			}
			box = domain.AddCard(box, card)
			autosave(box, path)
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
