package main

import "fmt"

func main() {
	fmt.Println("CLI Todo app")

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos) // passed as a reference to populate todos

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	storage.Save(todos)
}
