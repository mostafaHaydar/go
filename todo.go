package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func isTaskExists(ID int) bool {
	state := false
	current := first
	for current != nil {
		if current.ID == ID {
			state = true
			break
		}
		current = current.next
	}
	return state
}
func clearScreen() {
	clearCmd := exec.Command("cmd", "/c", "cls")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

func waitForKeyPress() {
	fmt.Println("Operation completed successfully.\n\t")
	fmt.Printf("Press Enter to return to the menu... ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

type Task struct {
	ID          int
	Name        string
	Description string
	Done        bool
	next        *Task
}

var lastTaskId = 0
var first *Task

func addTask() {
	clearScreen()
	var current *Task
	var previous *Task
	current = first
	for current != nil {
		previous = current
		current = current.next
	}
	current = new(Task)
	if first == nil {
		first = current
	}
	if previous != nil {
		previous.next = current
	}
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Printf("first write the name of your task please:: \n")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	current.Name = input
	fmt.Printf("second write the description of your task please:: \n")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	current.Description = input
	current.ID = lastTaskId
	lastTaskId++
	previous = current

}

func taskDone() {
	clearScreen()
	fmt.Printf("Please enter the task ID to update its information: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, _ := strconv.Atoi(input)
	current := first
	if isTaskExists(number) {
		for current != nil {
			if current.ID == number {
				if !current.Done {
					current.Done = true
				}
				return
			}
			current = current.next
		}
	} else {
		fmt.Println("There is no task with this ID.")
		return
	}
}
func updateTask() {
	clearScreen()
	fmt.Printf("Please enter the task ID to update its information: ")
	reader := bufio.NewReader(os.Stdin)
	var input string
	var number int
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, _ = strconv.Atoi(input)
	current := first
	if current == nil {
		fmt.Println("The to-do list is empty. Please add tasks.")
		return
	}
	taskExists := false
	for current != nil {
		if current.ID == number {
			taskExists = true
			clearScreen()
			fmt.Printf("please entre the new name::\n")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			current.Name = input
			fmt.Printf("please entre the new description ::\n")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			current.Description = input
			return
		}
		current = current.next
	}
	if !taskExists {
		fmt.Println("There is no task with this ID.")
	}
}
func removeTask() {
	clearScreen()
	fmt.Printf("Please enter the task ID to delete it : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, _ := strconv.Atoi(input)
	current := first
	if current == nil {
		fmt.Println("The to-do list is empty. Please add tasks.")
		return
	}
	taskExists := false

	var previous *Task

	for current != nil {
		if current.ID == number {

			if previous != nil {
				previous.next = current.next
			} else {
				first = current.next
			}
			return

		}
		previous = current
		current = current.next
	}

	if !taskExists {
		fmt.Println("There is no task with this ID.")
	}

}
func getTaskInformation() {
	clearScreen()

	fmt.Printf("Please enter the task ID to view its information: ")
	reader := bufio.NewReader(os.Stdin)
	var input string
	var number int
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, _ = strconv.Atoi(input)
	current := first
	if current == nil {
		fmt.Println("The to-do list is empty. Please add tasks.")
		return
	}
	taskExists := false
	for current != nil {
		if current.ID == number {
			taskExists = true
			clearScreen()
			fmt.Printf("######  start  #####\n\n\t")
			fmt.Printf("Task ID             : %d\n\t", current.ID)
			fmt.Printf("Task Name           : %s\n\t", current.Name)
			fmt.Printf("Task Description    : %s\n\t", current.Description)
			fmt.Printf("Task Status         : %t\n\n", current.Done)
			fmt.Printf("######  end  #####\n")
			return
		}
		current = current.next
	}
	if !taskExists {
		fmt.Println("There is no task with this ID.")
	}
}
func getAllTasksInformations() {
	clearScreen()
	current := first
	if current == nil {
		fmt.Println("The to-do list is empty. Please add tasks.")
		return
	}
	for current != nil {
		fmt.Printf("######  %d  #####\n\n\t", current.ID)
		fmt.Printf("Task ID             : %d\n\t", current.ID)
		fmt.Printf("Task Name           : %s\n\t", current.Name)
		fmt.Printf("Task Description    : %s\n\t", current.Description)
		fmt.Printf("Task Status         : %t\n\n", current.Done)
		current = current.next
	}
	fmt.Printf("######  end  #####\n")

}
func getUserOperationNumber() int {
	clearScreen()
	fmt.Println("Welcome to my first full program written in Go language.")
	fmt.Println("This system is a very basic to-do list app for daily tasks.")
	fmt.Println("To create a new task              :: 0")
	fmt.Println("To update a task                  :: 1")
	fmt.Println("To delete a task                  :: 2")
	fmt.Println("To get task info                  :: 3")
	fmt.Println("To get all tasks info             :: 4")
	fmt.Println("'done' remark for completed tasks :: 5")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please enter your choice (0-4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err == nil && number >= 0 && number <= 5 {
			return number
		}
		clearScreen()
		fmt.Println("Invalid input. Please enter a number between 0 and 4.")
	}
}

func maind() {

	for {
		number := getUserOperationNumber()
		switch number {
		case 0:
			addTask()
			waitForKeyPress()
		case 1:
			updateTask()
			waitForKeyPress()
		case 2:
			removeTask()
			waitForKeyPress()
		case 3:
			getTaskInformation()
			waitForKeyPress()
		case 4:
			getAllTasksInformations()
			waitForKeyPress()
		case 5:
			taskDone()
			waitForKeyPress()
		default:
			fmt.Printf("please entre a valid number between 0 and 4")
			waitForKeyPress()
		}

	}

}
