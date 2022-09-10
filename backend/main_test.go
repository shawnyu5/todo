package backend_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	// TODO: idk how to store the original file contents
	// It("should save tasks to config file", func() {
	// tasks := []backend.Task{
	// {Name: "test", DueDate: "test", Priority: 1, ReminderTime: "test"},
	// }
	// configDir := configdir.LocalConfig("todo")
	// configFile := path.Join(configDir, "tasks.json")
	// f, err := os.Open(configFile)
	// defer f.Close()
	// if err != nil {
	// panic(err)
	// }
	// // read the og contents of the file
	// og, err := ioutil.ReadAll(f)

	// backend.SaveTasks(tasks)

	// f, err = os.Open(configFile)
	// if err != nil {
	// panic(err)
	// }
	// newFile, err := ioutil.ReadAll(f)
	// if err != nil {
	// panic(err)
	// }
	// j, _ := json.Marshal(tasks)
	// Expect(string(newFile)).To(Equal(string(j)))

	// fmt.Println(fmt.Sprintf("string(og): %v", string(og))) // __AUTO_GENERATED_PRINT_VAR__
	// f.WriteString(string(og))
	// fmt.Println("end") // __AUTO_GENERATED_PRINTF__
	// })
})
