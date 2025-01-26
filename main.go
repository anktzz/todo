package main

func main() {
	todo := Task{}
	storage := NewStorage[Task]("task.json")
	storage.Load(&todo)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todo)
	storage.Save(todo)
}
