package backend_test

import (
	"encoding/json"
	"log"
	"path"
	"time"
	"todo/backend"

	"github.com/google/go-cmp/cmp"
	"github.com/kirsle/configdir"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Main", func() {
	Context("SaveTasks()", func() {
		It("should save tasks to existing config file", func() {
			// use mock file system
			backend.AppFs = afero.NewMemMapFs()
			tasks := []backend.Task{
				{Name: "test", DueDate: time.Now(), Priority: 1, ReminderTime: time.Now()},
			}
			backend.SaveTasks(tasks)

			configDir := configdir.LocalConfig("todo")
			configFile := path.Join(configDir, "tasks.json")
			f, err := backend.AppFs.Open(configFile)
			if err != nil {
				log.Fatal(err)
			}
			var fileTasks []backend.Task
			json.NewDecoder(f).Decode(&fileTasks)
			same := cmp.Equal(tasks, fileTasks)
			Expect(same).To(BeTrue(), "config file contents are incorrect")
		})
	})

	Context("LoadTasks()", func() {
		It("should load tasks from existing config file", func() {
			// use mock file system
			backend.AppFs = afero.NewMemMapFs()
			tasks := []backend.Task{
				{Name: "test", DueDate: time.Now(), Priority: 1, ReminderTime: time.Now()},
			}
			backend.SaveTasks(tasks)

			loadedTasks := backend.LoadTasks()
			same := cmp.Equal(tasks, loadedTasks)
			Expect(same).To(BeTrue(), "loaded tasks are incorrect")
		})

		It("should return an empty array if config file does not exist", func() {
			// use mock file system
			backend.AppFs = afero.NewMemMapFs()
			tasks := backend.LoadTasks()
			Expect(tasks).To(BeEmpty(), "tasks should be empty")
		})

		Context("AddTask()", func() {
			It("should add a task", func() {
				backend.AppFs = afero.NewMemMapFs()
				newTask := backend.Task{Name: "test", Desription: "test test test", DueDate: time.Now(), Priority: 1, ReminderTime: time.Now()}
				backend.AddTask(newTask)
				allTasks := backend.LoadTasks()

				found := false
				for _, task := range allTasks {
					if cmp.Equal(task, newTask) {
						found = true
					}
				}
				Expect(found).To(BeTrue(), "task was not added")
			})
		})
	})
})
