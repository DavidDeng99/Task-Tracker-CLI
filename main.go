package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Task struct {
	id   int
	task string
}

const INVALID = "Invalid command - type help for more info"

func main() {
	tasks := make([]Task, 0)
	fmt.Println("Welcome to Task Tracker CLI - type \"exit\" to quit")

	file, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if reader.Scan() {
			line := reader.Text()
			arguments := strings.Fields(line)
			switch arguments[0] {
			case "add":
				add(arguments, tasks)
			case "remove": //do1
			case "show":
				show(arguments, tasks)
			case "help": //do2
			case "exit":
				os.Exit(0)
			default:
				fmt.Println(INVALID)
			}
		}
	}
}

func add(args []string, tasks []Task) {
	taskArgument := args[2:]
	task := Task{len(tasks), strings.Join(taskArgument, " ")}
	tasks = append(tasks, task)
}

func show(args []string, tasks []Task) {
	for _, task := range tasks {
		fmt.Println(task)
	}
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d, task: %s", t.id, t.task)
}
