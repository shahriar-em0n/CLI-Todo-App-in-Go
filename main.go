package main

// import "fmt"

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	Storage.Save(todos)
}
