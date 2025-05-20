package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Id        int
	Task      string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

const INVALID = "Invalid command - type help for more info"

func main() {
	tasks := load()
	fmt.Println("Welcome to Task Tracker CLI - type \"exit\" to quit")

	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if reader.Scan() {
			line := reader.Text()
			args := strings.Fields(line)
			switch args[0] {
			case "add":
				add(strings.Join(args[1:], " "), &tasks)
			case "delete":
				id, _ := strconv.Atoi(args[1])
				delete(id, &tasks)
			case "update":
				id, _ := strconv.Atoi(args[1])
				update(id, strings.Join(args[2:], " "), &tasks)
			case "mark-in-progress":
				id, _ := strconv.Atoi(args[1])
				mark(id, &tasks, "in-progress")
			case "mark-done":
				id, _ := strconv.Atoi(args[1])
				mark(id, &tasks, "done")
			case "list":
				if len(args) == 1 {
					listAll(&tasks)
				} else {
					listFiltered(&tasks, args[1])
				}
			case "exit":
				os.Exit(0)
			default:
				fmt.Println(INVALID)
			}
		}
		save(&tasks)
	}
}

func add(task string, tasks *[]Task) {
	*tasks = append(*tasks, Task{len(*tasks), task, "todo", time.Now(), time.Now()})
	fmt.Printf("Task added successfully (ID: %d)\n", len(*tasks)-1)
}

func delete(id int, tasks *[]Task) {
	idx := findIndexById(id, tasks)
	*tasks = append((*tasks)[:idx], (*tasks)[idx+1:]...)
}

func update(id int, task string, tasks *[]Task) {
	idx := findIndexById(id, tasks)
	(*tasks)[idx].Task = task
}

func mark(id int, tasks *[]Task, status string) {
	idx := findIndexById(id, tasks)
	(*tasks)[idx].Status = status
}

func listAll(tasks *[]Task) {
	for _, task := range *tasks {
		fmt.Println(task)
	}
}

func listFiltered(tasks *[]Task, status string) {
	filtered := filterStatus(tasks, status)
	listAll(&filtered)
}

func (t Task) String() string {
	return fmt.Sprintf("ID: %d, task: %s, status: %s", t.Id, t.Task, t.Status)
}

func findIndexById(id int, tasks *[]Task) int {
	for idx, task := range *tasks {
		if id == task.Id {
			return idx
		}
	}

	return -1
}

func load() []Task {
	data, err := os.ReadFile("tasks.json")
	tasks := make([]Task, 0)
	if err != nil {
		return tasks
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks
	}

	return tasks
}

func save(tasks *[]Task) {
	data, err := json.MarshalIndent(*tasks, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("tasks.json", data, 0644)
}

func filterStatus(tasks *[]Task, status string) []Task {
	filtered := make([]Task, 0)
	for _, task := range *tasks {
		if strings.Compare(task.Status, status) == 0 {
			filtered = append(filtered, task)
		}
	}
	return filtered
}
