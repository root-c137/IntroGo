package main

import (
	"bufio"
	"fmt"
	"os"
)

var Tasks []string

func main() {
	//var task string
	var response string
	fmt.Println("Ajouter une tâche (y/n): ")
	fmt.Scanln(&response)

	if response == "y" {
		addTask()
	} else if response == "n" {
		os.Exit(0)
	} else {
		fmt.Println("WTF...")
		main()
	}
}

func addTask() {
	scanner := bufio.NewScanner(os.Stdin)

	var response string = ""
	for len(response) == 0 {
		fmt.Println("Description de la tâche : ")
		scanner.Scan()
		response = scanner.Text()

		if len(response) > 0 {
			Tasks = append(Tasks, response)
			fmt.Println("La tâche a été ajoutée.")
			fmt.Println(Tasks)
			main()
		}
	}
}
