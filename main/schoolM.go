package main

import (
	"bufio"
	"fmt"
	"golang/main/greet"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type CLASSE struct {
	id   int
	name string
	next *CLASSE
}

var first *CLASSE = nil
var current *CLASSE = nil
var previous *CLASSE = nil
var lastClassId int = 0

var reader = bufio.NewReader(os.Stdin)

func clearTerminal() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
func userOperation() int {
	clearTerminal()
	userNumber := 0
	fmt.Println("Bonjour monsieur, comment allez-vous ?")
	fmt.Println("Pour commencer dans ce programme de gestion scolaire, veuillez d'abord entrer le numéro de l'opération :")
	fmt.Println("\t1 ==> Pour créer une classe, tapez sur")
	fmt.Println("\t2 ==> Pour modifier une classe, tapez sur")
	fmt.Println("\t3 ==> Pour supprimer une classe, tapez sur")
	fmt.Print("\t\tVotre numéro de l'opération s'il vous plaît : ")

	for true {
		var err error
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userNumber, err = strconv.Atoi(input)
		if err == nil && userNumber > 0 && userNumber < 4 {
			fmt.Printf("Vous avez choisi l'opération numéro : %d\n", userNumber)
			break
		} else {
			clearTerminal()
			fmt.Printf("\tVous avez choisi l'opération numéro ::")
		}
	}
	return userNumber
}

func createNewClass() {

	current = new(CLASSE)
	if first == nil {
		first = current
	}

	if previous != nil {
		previous.next = current
	}

	current.id = lastClassId
	lastClassId = lastClassId + 1
	clearTerminal()
	fmt.Printf("\n\nsil te plait le nom de la classe::")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	current.name = input
	current.next = nil

	previous = current

}

func seeAllClassesInforamtion() {
	clearTerminal()
	current = first
	for current != nil {
		fmt.Printf("-------------------------------------------------\n")

		fmt.Printf("l'ID de la classe       ==> %d\n", current.id)
		fmt.Printf("le nom de la classe     ==> %s\n", current.name)
		current = current.next
	}
	fmt.Printf("-------------------------------------------------\n")
}
func backToMenu() {
	fmt.Printf("l'operation a ete efectue avec succes , tapes sur 0 pour le retour al la page principale :: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return
}
func main() {
	greet.Hello()

	for true {
		switch userOperation() {
		case 1:
			fmt.Printf("tu choisiez 1")
			createNewClass()
			backToMenu()
		case 2:
			fmt.Printf("tu choisiez 2")
			seeAllClassesInforamtion()
			backToMenu()
		case 3:
			fmt.Printf("tu choisiez 3")
			seeAllClassesInforamtion()
			backToMenu()
		}
	}

}
