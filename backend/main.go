package backend

import (
	"errors"
	"log"
	"os"
	"path"

	"github.com/kirsle/configdir"
)

type Task struct {
	Name         string
	DueDate      string
	Priority     int
	ReminderTime string
}

func CreateTask() {
	configDir := configdir.LocalConfig("todo")

	// if the config directory doesn't exist, create it
	if _, err := os.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	configFile := path.Join(configDir, "tasks.json")
	f, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString("hello")
}
