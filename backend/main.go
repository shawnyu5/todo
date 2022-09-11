package backend

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"time"

	"github.com/kirsle/configdir"
	"github.com/spf13/afero"
)

type Task struct {
	Name         string    `json:"name"`
	DueDate      time.Time `json:"due_date"`
	Priority     int       `json:"priority"`
	ReminderTime time.Time `json:"reminder_time"`
}

var AppFs = afero.NewOsFs()

// SaveTasks saves the current tasks to the config directory as json objects. Note this over writes the contents of the config file, does not append to it
func SaveTasks(items []Task) {
	configDir := configdir.LocalConfig("todo")

	// if the config directory doesn't exist, create it
	if _, err := AppFs.Stat(configDir); errors.Is(err, os.ErrNotExist) {
		err := AppFs.MkdirAll(configDir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	configFile := path.Join(configDir, "tasks.json")
	f, err := AppFs.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	j, err := json.Marshal(items)
	f.WriteString(string(j))
}

// LoadTasks loads the tasks from the config directory.
func LoadTasks() []Task {
	configDir := configdir.LocalConfig("todo")
	configFile := path.Join(configDir, "tasks.json")
	f, err := AppFs.Open(configFile)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	var tasks []Task
	err = json.NewDecoder(f).Decode(&tasks)
	if err != nil {
		log.Println(err)
	}
	return tasks
}
