package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", n.Title, n.Content)
}

func (note Note) Save() error {
	filename := strings.ToLower((strings.ReplaceAll(note.Title, " ", "_")))
	filename = filename + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)

}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
