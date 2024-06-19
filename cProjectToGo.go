package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error clearing terminal:", err)
	}
}

type CLASS_LINKED_LISTS struct {
	id             int
	name           string
	studentsNumber int
	date           string
	next           *CLASS_LINKED_LISTS
}

type STUDENT_LINKED_LISTS struct {
	id        int
	firstName string
	lastName  string
	age       int
	email     string
	className string
	date      string
	next      *STUDENT_LINKED_LISTS
}

var first *CLASS_LINKED_LISTS
var current *CLASS_LINKED_LISTS
var previous *CLASS_LINKED_LISTS

var first_s *STUDENT_LINKED_LISTS
var current_s *STUDENT_LINKED_LISTS
var previous_s *STUDENT_LINKED_LISTS

func backToMenu() int {
	var tmpVar int
	fmt.Printf("\n Si vous pouvez aller vers la page principale, tapez 0 :: ")
	fmt.Scanf("%d", &tmpVar)
	return tmpVar
}

func createNewClass(pLastClassId *int) {
	clearTerminal()

	current = new(CLASS_LINKED_LISTS)

	if first == nil {
		first = current
	}

	if previous != nil {
		previous.next = current
	}

	currentTime := time.Now()
	// Format the time as desired
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	var tmpName string
	var isValid bool = false
	var alreadyExists bool = false

	for !isValid {
		fmt.Printf("S'il te plaît, entre le nom de la classe :\n\t==> ")

		reader := bufio.NewReader(os.Stdin)
		tmpName, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		if len(tmpName) > 0 && tmpName[len(tmpName)-1] == '\n' {
			tmpName = tmpName[:len(tmpName)-1]
		}

		alreadyExists = false

		// check class name if alreay used by another class

		var tmpCurrent *CLASS_LINKED_LISTS = first
		for tmpCurrent != nil {
			if tmpCurrent.name == tmpName {
				clearTerminal()
				fmt.Printf("Ce nom de classe est déjà utilisé dans une autre classe.\n\t")
				alreadyExists = true
				break
			}
			tmpCurrent = tmpCurrent.next
		}

		if !isClassNameValid(tmpName) {
			clearTerminal()
			fmt.Printf("Ce nom de classe contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
		}
		if !alreadyExists && isClassNameValid(tmpName) {
			current.id = *pLastClassId
			*pLastClassId = (*pLastClassId) + 1
			current.name = tmpName
			current.studentsNumber = 0
			current.date = formattedTime
			isValid = true
			break
		}
	}

	current.next = nil
	previous = current
	fmt.Printf("\nL'opération a réussi avec succès.\n")
	tmpName = ""
	fmt.Printf("%s", tmpName)
}

func updateClass() {
	clearTerminal()
	current = first
	var tmpClassId int
	classExists := false

	fmt.Printf("Pour modifier les informations d'une classe, il faut d'abord son ID : ")
	fmt.Scanf("%d", &tmpClassId)

	for current != nil {
		if current.id == tmpClassId {
			classExists = true
			var tmpClassNameForUpdateStudents string
			tmpClassNameForUpdateStudents = current.name
			var tmpName string
			isClassNameValidBool := false
			for !isClassNameValidBool {
				fmt.Printf("S'il te plaît, entre le nouveau nom de la classe :\n\t==> ")

				// fgets(tmpName, sizeof(tmpName), stdin)

				reader := bufio.NewReader(os.Stdin)
				tmpName, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				if len(tmpName) > 0 && tmpName[len(tmpName)-1] == '\n' {
					tmpName = tmpName[:len(tmpName)-1]
				}
				isExists := false
				// check class name if alreay used by another class

				var tmpCurrent *CLASS_LINKED_LISTS = first
				for tmpCurrent != nil {
					if tmpCurrent.name == tmpName {
						clearTerminal()
						fmt.Printf("Ce nom de classe est déjà utilisé dans une autre classe.\n\t")
						isExists = true
						break
					}
					tmpCurrent = tmpCurrent.next
				}

				if !isExists {
					if !isClassNameValid(tmpName) {
						clearTerminal()
						fmt.Printf("Ce nom de classe contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
					} else {
						isClassNameValidBool = true
					}
				}
			}

			current.name = tmpName

			current_s = first_s
			for current_s != nil {
				if current_s.className == tmpClassNameForUpdateStudents {
					current_s.className = current.name
				} else {
					continue
				}
				current_s = current_s.next
			}
		}
		current = current.next
	}

	if !classExists {
		fmt.Printf("\nIl n'existe pas de classe avec cet identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func deleteClass() {

	clearTerminal()
	fmt.Printf("Pour supprimer une classe, il faut l'ID de cette classe avant la suppression : ")
	classExists := false
	var tmpClassId int
	fmt.Scanf("%d", &tmpClassId)
	var tmpClassNameForDeleteStudents string
	current = first
	previous = nil
	for current != nil {
		if current.id == tmpClassId {
			classExists = true
			tmpClassNameForDeleteStudents = current.name
			if first.next == nil {
				first = nil
			} else {
				if previous != nil {
					previous.next = current.next
					current = nil
				} else {
					first = first.next
					current = nil
				}
			}

			current_s = first_s
			previous_s = nil
			for current_s != nil {
				if current_s.className == tmpClassNameForDeleteStudents {
					if first_s.next == nil {

						first_s = nil
					} else {
						if previous_s != nil {
							previous_s.next = current_s.next

							current_s = nil
						} else {
							first_s = first_s.next

							current_s = nil
						}
					}
					break
				}
				previous_s = current_s
				current_s = current_s.next
			}
			break
		}
		previous = current
		current = current.next
	}
	if !classExists {
		fmt.Printf("\nIl n'existe pas de classe avec cet identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func classInformation() {
	var tmpClassId int
	clearTerminal()
	fmt.Printf("Pour voir les informations d'une classe, il faut l'ID de cette classe : ")

	fmt.Scanf("%d", &tmpClassId)

	clearTerminal()
	classExists := false

	current = first

	for current != nil {
		if current.id == tmpClassId {
			classExists = true
			fmt.Printf("_______________________________________________________\n\n")
			fmt.Printf("id                 :: %d\n", current.id)
			fmt.Printf("nom de class       :: %s\n", current.name)
			fmt.Printf("nombre d'élèves    :: %d\n", current.studentsNumber)
			fmt.Printf("date               :: %s\n", current.date)
			fmt.Printf("_______________________________________________________\n\n")
		}
		current = current.next
	}
	if !classExists {
		fmt.Printf("\nIl n'existe pas de classe avec ce identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func allClassesInformation() {
	clearTerminal()
	classExists := false
	current = first

	for current != nil {
		classExists = true
		fmt.Printf("_______________________________________________________\n\n")
		fmt.Printf("id                 :: %d\n", current.id)
		fmt.Printf("nom de class       :: %s\n", current.name)
		fmt.Printf("nombre d'élèves    :: %d\n", current.studentsNumber)
		fmt.Printf("date               :: %s\n", current.date)
		fmt.Printf("_______________________________________________________\n\n")
		current = current.next
	}

	if !classExists {
		fmt.Printf("\nIl n'y a aucune classe dans ces écoles.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func choseOperation() int {
	var crudOperationNumber int
	clearTerminal()
	fmt.Printf("##########################################################################\n")
	fmt.Printf("################ Bonjour, dans mon système scolaire ######################\n")
	fmt.Printf("##########################################################################\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#    ##########                                                          #\n")
	fmt.Printf("#    # classe #                                                          #\n")
	fmt.Printf("#    ##########                                                          #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#\t:: cC ==> 0 :: pour ajouter une classe                           #\n")
	fmt.Printf("#\t:: uC ==> 1 :: pour modifier les informations d'une classe       #\n")
	fmt.Printf("#\t:: dC ==> 2 :: pour supprimer une classe                         #\n")
	fmt.Printf("#\t:: rC ==> 3 :: pour voir les informations d'une classe           #\n")
	fmt.Printf("#\t:: aC ==> 4 :: pour voir les informations de toutes les classes  #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#    ############                                                        #\n")
	fmt.Printf("#    # étudiant #                                                        #\n")
	fmt.Printf("#    ############                                                        #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#\t:: cS ==> 5 :: Pour ajouter un étudiant                          #\n")
	fmt.Printf("#\t:: uS ==> 6 :: pour modifier les informations d'un étudiant      #\n")
	fmt.Printf("#\t:: dS ==> 7 :: pour supprimer un étudiant                        #\n")
	fmt.Printf("#\t:: rS ==> 8 :: pour voir les informations d'un étudiant          #\n")
	fmt.Printf("#\t:: aS ==> 9 :: pour voir les informations de tous les étudiant   #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#\t:: exit ==> 10 :: pour sortir                                    #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("#                                                                        #\n")
	fmt.Printf("##########################################################################\n\n\t")
	fmt.Printf("Quelle opération souhaitez-vous effectuer ? : ")

	for true {
		fmt.Scanf("%d", &crudOperationNumber)
		if crudOperationNumber < 11 && crudOperationNumber >= 0 {
			return crudOperationNumber
		} else {
			clearTerminal()
			fmt.Printf("Veuillez entrer un nombre entre 0 et 10.\n\n\t")
			fmt.Printf("Nombre d'opérations à effectuer :: ")
		}
	}
	return crudOperationNumber
}

func createNewStudent(pLastStudentId *int) {
	clearTerminal()

	currentTime := time.Now()
	// Format the time as desired
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	var tmpFirstName string
	var tmpLastName string
	var tmpEmail string
	var tmpClassName string
	var tmpAge int

	current_s = new(STUDENT_LINKED_LISTS)

	if first_s == nil {
		first_s = current_s
	}
	if previous_s != nil {
		previous_s.next = current_s
	}

	current_s.id = *pLastStudentId
	*pLastStudentId = (*pLastStudentId) + 1

	isFirstNameValid := false
	isLastNameValid := false
	isAgeValid := false
	isEmailValidBool := false

	for !isFirstNameValid {
		fmt.Printf("S'il te plaît, entre ton prénom :\n\t==> ")

		// reader := bufio.NewReader(os.Stdin)
		// tmpFirstName = reader.ReadString('\n')
		// tmpFirstName = tmpFirstName[:len(tmpFirstName)-1]

		reader := bufio.NewReader(os.Stdin)
		tmpFirstName, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if len(tmpFirstName) > 0 && tmpFirstName[len(tmpFirstName)-1] == '\n' {
			tmpFirstName = tmpFirstName[:len(tmpFirstName)-1]
		}

		if isHummanNameValid(tmpFirstName) {
			current_s.firstName = tmpFirstName
			isFirstNameValid = true
		} else {
			clearTerminal()
			fmt.Printf("Ce prénom contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
		}
	}

	for !isLastNameValid {
		fmt.Printf("S'il te plaît, entre ton nom:\n\t==> ")

		// reader := bufio.NewReader(os.Stdin)
		// tmpLastName = reader.ReadString('\n')
		// tmpLastName = tmpLastName[:len(tmpLastName)-1]

		reader := bufio.NewReader(os.Stdin)
		tmpLastName, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if len(tmpLastName) > 0 && tmpLastName[len(tmpLastName)-1] == '\n' {
			tmpLastName = tmpLastName[:len(tmpLastName)-1]
		}

		if isHummanNameValid(tmpLastName) {
			current_s.lastName = tmpLastName
			isLastNameValid = true
		} else {
			clearTerminal()
			fmt.Printf("Ce nom contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
		}
	}

	for !isAgeValid {
		fmt.Printf("S'il te plaît, entre ton âge:\n\t==> ")
		fmt.Scanf("%d", &tmpAge)
		if tmpAge > 15 && tmpAge < 25 {
			current_s.age = tmpAge
			isAgeValid = true
		} else {
			clearTerminal()
			fmt.Printf("S'il te plaît, entre un âge valide entre 15 et 25.\n\t")
		}
	}

	for !isEmailValidBool {
		fmt.Printf("S'il te plaît, entre ton email:\n\t==> ")

		// reader := bufio.NewReader(os.Stdin)
		// tmpEmail = reader.ReadString('\n')
		// tmpEmail = tmpEmail[:len(tmpEmail)-1]

		reader := bufio.NewReader(os.Stdin)
		tmpEmail, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if len(tmpEmail) > 0 && tmpEmail[len(tmpEmail)-1] == '\n' {
			tmpEmail = tmpEmail[:len(tmpEmail)-1]
		}

		if isEmailValid(tmpEmail) {
			current_s.email = tmpEmail
			current_s.date = formattedTime
			isEmailValidBool = true
		} else {
			clearTerminal()
			fmt.Printf("S'il te plaît, entre un email valide.\n\t")
		}
	}

	fmt.Printf("S'il te plaît, entre ta classe:\n\t==> ")
	isTheSame := false
	for !isTheSame {

		// reader := bufio.NewReader(os.Stdin)
		// tmpClassName = reader.ReadString('\n')
		// tmpClassName = tmpClassName[:len(tmpClassName)-1]

		reader := bufio.NewReader(os.Stdin)
		tmpClassName, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		if len(tmpClassName) > 0 && tmpClassName[len(tmpClassName)-1] == '\n' {
			tmpClassName = tmpClassName[:len(tmpClassName)-1]
		}

		var tmpCurrent *CLASS_LINKED_LISTS = first

		for tmpCurrent != nil {
			if tmpCurrent.name == tmpClassName {
				tmpCurrent.studentsNumber = tmpCurrent.studentsNumber + 1
				isTheSame = true
				break
			}
			tmpCurrent = tmpCurrent.next
		}

		if !isTheSame {
			clearTerminal()
			fmt.Printf("Cette classe n'existe pas dans l'école??\n")
			fmt.Printf("S'il te plaît, entre ta classe::\n\t")
		}
	}

	current_s.className = tmpClassName

	current_s.next = nil
	previous_s = current_s
	fmt.Printf("\nL'opération a réussi avec succès.\n")

	tmpFirstName = ""
	tmpLastName = ""
	tmpEmail = ""
	fmt.Printf("%s%s%s", tmpFirstName, tmpLastName, tmpEmail)
}

func updateStudent() {

	clearTerminal()
	studentExists := false
	fmt.Printf("Pour modifier les informations d'un étudiant, il faut d'abord l'ID de cet étudiant : ")

	var tmpStudentId int
	fmt.Scanf("%d", &tmpStudentId)

	current_s = first_s

	for current_s != nil {
		if current_s.id == tmpStudentId {
			var oldClassName string

			oldClassName = current_s.className

			studentExists = true

			var tmpFirstName string
			var tmpLastName string
			var tmpEmail string
			var tmpClassName string

			isFirstNameValid := false
			isLastNameValid := false
			isAgeValid := false

			for !isFirstNameValid {
				fmt.Printf("S'il te plaît, entre le nouveau prénom:\n\t==> ")

				reader := bufio.NewReader(os.Stdin)
				tmpFirstName, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				if len(tmpFirstName) > 0 && tmpFirstName[len(tmpFirstName)-1] == '\n' {
					tmpFirstName = tmpFirstName[:len(tmpFirstName)-1]
				}

				if isHummanNameValid(tmpFirstName) {
					current_s.firstName = tmpFirstName
					isFirstNameValid = true
				} else {
					clearTerminal()

					fmt.Printf("Ce prénom contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
				}
			}
			tmpFirstName = ""
			fmt.Printf("%s", tmpFirstName)
			for !isLastNameValid {
				fmt.Printf("S'il te plaît, entre le nouveau nom :\n\t==> ")
				// fgets(tmpLastName, sizeof(tmpLastName), stdin)

				reader := bufio.NewReader(os.Stdin)
				tmpLastName, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				if len(tmpLastName) > 0 && tmpLastName[len(tmpLastName)-1] == '\n' {
					tmpLastName = tmpLastName[:len(tmpLastName)-1]
				}

				current_s.lastName = tmpLastName

				if isHummanNameValid(tmpLastName) {
					current_s.lastName = tmpLastName
					isLastNameValid = true
				} else {
					clearTerminal()

					fmt.Printf("Ce nom contient des caractères interdits ou est trop court ou trop long.(5,100)\n\t")
				}
			}
			tmpLastName = ""
			fmt.Printf("%s", tmpLastName)
			for !isAgeValid {
				var tmpAge int
				fmt.Printf("S'il te plaît, entre le nouvel âge:\n\t==> ")
				fmt.Scanf("%d", &tmpAge)
				if tmpAge > 15 && tmpAge < 25 {
					current_s.age = tmpAge
					isAgeValid = true
				} else {
					clearTerminal()
					fmt.Printf("S'il te plaît, entre un âge valide entre 15 et 25.\n\t")
				}

			}

			isEmailValidBool := false
			for !isEmailValidBool {
				fmt.Printf("S'il te plaît, entre le nouvel email:\n\t==> ")
				// fgets(tmpEmail, sizeof(tmpEmail), stdin)

				reader := bufio.NewReader(os.Stdin)
				tmpEmail, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				if len(tmpEmail) > 0 && tmpEmail[len(tmpEmail)-1] == '\n' {
					tmpEmail = tmpEmail[:len(tmpEmail)-1]
				}

				if isEmailValid(tmpEmail) {
					current_s.email = tmpEmail
					isEmailValidBool = true
				} else {
					clearTerminal()

					fmt.Printf("S'il te plaît, entre un email valide.\n\t")
				}
			}
			tmpEmail = ""
			fmt.Printf("%s", tmpEmail)
			fmt.Printf("S'il te plaît, entre la nouvelle classe:\n\t==> ")
			isTheSame := false

			for !isTheSame {

				// fgets(tmpClassName, stdin)

				reader := bufio.NewReader(os.Stdin)
				tmpClassName, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				if len(tmpClassName) > 0 && tmpClassName[len(tmpClassName)-1] == '\n' {
					tmpClassName = tmpClassName[:len(tmpClassName)-1]
				}

				current = first
				for current != nil {
					if current.name == tmpClassName {
						isTheSame = true
						current_s.className = tmpClassName
						if oldClassName == tmpClassName {
							current.studentsNumber = current.studentsNumber + 1
							var tmpCurrent *CLASS_LINKED_LISTS = first
							for tmpCurrent != nil {
								if oldClassName == tmpCurrent.name {
									tmpCurrent.studentsNumber = tmpCurrent.studentsNumber - 1
									break
								}
								tmpCurrent = tmpCurrent.next
							}
						}
						break
					}
					current = current.next
				}

				if !isTheSame {
					clearTerminal()

					fmt.Printf("Cette classe n'existe pas dans l'école??\n")
					fmt.Printf("S'il te plaît, entre ta classe::\n\t")
				}
			}
			tmpClassName = ""
			fmt.Printf("%s", tmpClassName)

			break
		}
		current_s = current_s.next

	}
	if !studentExists {
		fmt.Printf("\nIl n'existe pas d'élève avec cet identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}

	// var tmpFirstName string
	// var tmpLastName string
	// var tmpEmail string
	// var tmpClassName string

}

func deleteStudent() {

	clearTerminal()

	fmt.Printf("Pour supprimer un étudiant, il faut l'ID de cet étudiant avant la suppression : ")

	studentExists := false
	var tmpStudentId int
	var tmpClasseName string
	fmt.Scanf("%d", &tmpStudentId)

	current = first

	current_s = first_s
	previous_s = nil
	for current_s != nil {
		if current_s.id == tmpStudentId {
			tmpClasseName = current_s.className
			if first_s.next == nil {
				first_s = nil
			} else {
				if previous_s != nil {
					previous_s.next = current_s.next
					current_s = nil
				} else {
					first_s = first_s.next
					current_s = nil
				}
			}

			for current != nil {
				if tmpClasseName == current.name {
					current.studentsNumber = current.studentsNumber - 1
					break
				}
				current = current.next
			}
			break
		}
		previous_s = current_s
		current_s = current_s.next
	}

	if !studentExists {
		fmt.Printf("\nIl n'existe pas d'élève avec cet identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func studentInformation() {

	studentExists := false
	clearTerminal()
	fmt.Printf("Pour voir les informations d'un étudiant, il faut l'ID de cet étudiant : ")

	var tmpStudentId int
	fmt.Scanf("%d", &tmpStudentId)

	clearTerminal()

	current_s = first_s
	for current_s != nil {
		if current_s.id == tmpStudentId {
			studentExists = true
			fmt.Printf("_______________________________________________________\n\n")
			fmt.Printf("id                   :: %d\n", current_s.id)
			fmt.Printf("prénom               :: %s\n", current_s.firstName)
			fmt.Printf("nom de famille       :: %s\n", current_s.lastName)
			fmt.Printf("âge                  :: %d\n", current_s.age)
			fmt.Printf("mail                 :: %s\n", current_s.email)
			fmt.Printf("classe               :: %s\n", current_s.className)
			fmt.Printf("date                 :: %s\n", current_s.date)
			fmt.Printf("_______________________________________________________\n\n")
		}
		current_s = current_s.next
	}

	if !studentExists {
		fmt.Printf("\nIl n'existe pas d'élève avec cet identifiant.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

func allStudentsInformation() {
	studentExists := false
	clearTerminal()

	current_s = first_s
	for current_s != nil {
		studentExists = true
		fmt.Printf("_______________________________________________________\n\n")
		fmt.Printf("id                   :: %d\n", current_s.id)
		fmt.Printf("prénom               :: %s\n", current_s.firstName)
		fmt.Printf("nom de famille       :: %s\n", current_s.lastName)
		fmt.Printf("âge                  :: %d\n", current_s.age)
		fmt.Printf("mail                 :: %s\n", current_s.email)
		fmt.Printf("classe               :: %s\n", current_s.className)
		fmt.Printf("date                 :: %s\n", current_s.date)
		fmt.Printf("_______________________________________________________\n\n")
		current_s = current_s.next
	}

	if !studentExists {
		fmt.Printf("\nIl n'existe pas d'élève dans cette école.\n")
	} else {
		fmt.Printf("\nL'opération a réussi avec succès.\n")
	}
}

// func getDataFromFileClasses(pLastClassId *int) {

//   FILE *pFile = nil
//   pFile = fopen("classes.txt", "r")
//   char myString[1000]
//   fgets(myString, 1000, pFile)
//   fclose(pFile)
//   pFile = nil

//   char subString[100]

//   const char delimiter[] = "|"
//   const char subDelimiter[] = ":,"

//   char *outer_saveptr = nil
//   char *inner_saveptr = nil

//   char *token = (char *)strtok_s(myString, delimiter, &outer_saveptr)

//   enum dataElements {
//     id,
//     idVal,
//     name,
//     namVal,
//     studentsNumber,
//     studentsNumberVal,
//     date,
//     dateVal
//   }
//   enum dataElements state = id

//   int lastClassId = 0
//   int number

//   for (token != nil) {
//     copy(subString, sizeof(subString), token)
//     char *subToken = (char *)strtok_s(subString, subDelimiter, &inner_saveptr)

//     state = id
//     current =
//         (struct CLASS_LINKED_LISTS *)malloc(sizeof(struct CLASS_LINKED_LISTS))

//     if (first == nil) {
//       first = current
//     }
//     if (previous != nil) {
//       previous.next = current
//     }
//     for (subToken != nil) {
//       switch (state) {
//       case 0:
//         // do nothing
//         break
//       case 1:
//         number = atoi(subToken)
//         current.id = number
//         break
//       case 2:
//         // do nothing
//         break
//       case 3:
//         copy(current.name, sizeof(current.name), subToken)
//         break
//       case 4:
//         // do nothing
//         break
//       case 5:
//         number = atoi(subToken)
//         current.studentsNumber = number
//         break
//       case 6:
//         // do nothing
//         break
//       case 7:
//         copy(current.date, sizeof(current.date), subToken)
//         *pLastClassId = *pLastClassId + 1
//         break
//       }
//       state = state + 1
//       subToken = (char *)strtok_s(nil, subDelimiter, &inner_saveptr)
//     }
//     current.next = nil
//     previous = current
//     token = (char *)strtok_s(nil, delimiter, &outer_saveptr)
//   }
// }

// void getDataFromFileStudents(int *pLastStudentId) {
//   FILE *pFile = nil
//   pFile = fopen("students.txt", "r")
//   char myString[1000]
//   fgets(myString, 1000, pFile)
//   fclose(pFile)
//   pFile = nil

//   char subString[200]

//   const char delimiter[] = "|"
//   const char subDelimiter[] = ":,"

//   char *outer_saveptr = nil
//   char *inner_saveptr = nil

//   char *token = (char *)strtok_s(myString, delimiter, &outer_saveptr)

//   enum dataElements {
//     id,
//     idVal,
//     firstName,
//     firstNameVal,
//     lastName,
//     lastNameVal,
//     age,
//     ageVal,
//     email,
//     emailVal,
//     className,
//     classNameVal,
//     datee,
//     dateeVal
//   }
//   enum dataElements state = id

//   int lastStudentId = 0
//   int number

//   for (token != nil) {
//     copy(subString, sizeof(subString), token)
//     char *subToken = (char *)strtok_s(subString, subDelimiter, &inner_saveptr)

//     state = id
//     current_s = (struct STUDENT_LINKED_LISTS *)malloc(
//         sizeof(struct STUDENT_LINKED_LISTS))

//     if (first_s == nil) {
//       first_s = current_s
//     }
//     if (previous_s != nil) {
//       previous_s.next = current_s
//     }
//     for (subToken != nil) {
//       switch (state) {
//       case 0:
//         // do nothing
//         break
//       case 1:
//         number = atoi(subToken)
//         current_s.id = number
//         break
//       case 2:
//         // do nothing
//         break
//       case 3:
//         copy(current_s.firstName, sizeof(current_s.firstName), subToken)
//         break
//       case 4:
//         // do nothing
//         break
//       case 5:
//         copy(current_s.lastName, sizeof(current_s.lastName), subToken)
//         break
//       case 6:
//         // do nothing
//         break
//       case 7:
//         number = atoi(subToken)
//         current_s.age = number
//         break
//       case 8:
//         // do nothing
//         break
//       case 9:
//         copy(current_s.email, sizeof(current_s.email), subToken)
//         break
//       case 10:
//         // do nothing
//         break
//       case 11:
//         copy(current_s.className, sizeof(current_s.className), subToken)
//         break
//       case 12:
//         // do nothing
//         break
//       case 13:
//         copy(current_s.date, sizeof(current_s.date), subToken)
//         lastStudentId = lastStudentId + 1
//         *pLastStudentId = *pLastStudentId + 1
//         break
//       }
//       state = state + 1
//       subToken = (char *)strtok_s(nil, subDelimiter, &inner_saveptr)
//     }
//     current_s.next = nil
//     previous_s = current_s
//     token = (char *)strtok_s(nil, delimiter, &outer_saveptr)
//   }
// }

// void putDataIntoFileClasses() {
//   FILE *pFile = nil
//   pFile = fopen("classes.txt", "w")
//   if (pFile != nil) {
//     ffmt.Printf(pFile, "")
//     fclose(pFile)
//     pFile = nil
//   }

//   fclose(pFile)
//   pFile = nil

//   pFile = fopen("classes.txt", "a")

//   current = first
//   previous = nil
//   if (pFile != nil) {

//     for (current != nil) {

//       ffmt.Printf(pFile, "id:%d,", current.id)
//       ffmt.Printf(pFile, "name:%s,", current.name)
//       ffmt.Printf(pFile, "studentsNumber:%d,", current.studentsNumber)
//       ffmt.Printf(pFile, "date:%s|", current.date)

//       previous = current
//       current = current.next
//       free(previous)
//       previous = nil
//     }
//   }

//   fclose(pFile)
//   pFile = nil
// }

// void putDataIntoFileStudents() {
//   FILE *pFile = nil
//   pFile = fopen("students.txt", "w")
//   if (pFile != nil) {
//     ffmt.Printf(pFile, "")
//     fclose(pFile)
//     pFile = nil
//   }

//   fclose(pFile)
//   pFile = nil
//   pFile = fopen("students.txt", "a")
//   if (pFile != nil) {

//     current_s = first_s
//     previous_s = nil

//     for (current_s != nil) {

//       ffmt.Printf(pFile, "id:%d,", current_s.id)
//       ffmt.Printf(pFile, "firstName:%s,", current_s.firstName)
//       ffmt.Printf(pFile, "lastName:%s,", current_s.lastName)
//       ffmt.Printf(pFile, "age:%d,", current_s.age)
//       ffmt.Printf(pFile, "email:%s,", current_s.email)
//       ffmt.Printf(pFile, "className:%s,", current_s.className)
//       ffmt.Printf(pFile, "date:%s|", current_s.date)

//       previous_s = current_s
//       current_s = current_s.next
//       free(previous_s)
//       previous_s = nil
//     }
//   }

//   fclose(pFile)
//   pFile = nil
// }

func isHummanNameValid(name string) bool {
	notAllowedChars := "123456789#$^&*_#{[]}\\@=+*"
	state := true
	if len(name) < 5 || len(name) > 100 {
		state = false
	}

	if state {
		for i := 0; i < len(name); i++ {
			for j := 0; j < len(notAllowedChars); j++ {
				if name[i] == notAllowedChars[j] {
					state = false
					break
				}
			}
		}
	}
	return state
}

func isClassNameValid(name string) bool {
	notAllowedChars := "#$^&*_#{[]}\\@=+*"
	state := true
	if len(name) < 5 || len(name) > 100 {
		state = false
	}
	if state {
		for i := 0; i < len(name); i++ {
			for j := 0; j < len(notAllowedChars); j++ {
				if name[i] == notAllowedChars[j] {
					state = false
					break
				}
			}
		}
	}
	return state
}

func isEmailValid(email string) bool {
	state := false
	for i := 0; i < len(email); i++ {
		if email[i] == '@' && i != 0 && i != len(email)-1 {
			state = true
		}
	}
	return state
}
func maindddd() {

	lastClassId := 0
	var pLastClassId *int = &lastClassId
	lastStudentId := 0
	var pLastStudentId *int = &lastStudentId

	//   getDataFromFileClasses(pLastClassId)

	//   getDataFromFileStudents(pLastStudentId)

	for true {
		var userNumber int
		userNumber = choseOperation()
		switch userNumber {
		case 0:
			createNewClass(pLastClassId)
			if backToMenu() == 0 {
				break
			}
			break
		case 1:
			updateClass()
			if backToMenu() == 0 {
				break
			}
			break
		case 2:
			deleteClass()
			if backToMenu() == 0 {
				break
			}
			break
		case 3:
			classInformation()
			if backToMenu() == 0 {
				break
			}
			break
		case 4:
			allClassesInformation()
			if backToMenu() == 0 {
				break
			}
			break
		case 5:
			createNewStudent(pLastStudentId)
			if backToMenu() == 0 {
				break
			}
			break
		case 6:
			updateStudent()
			if backToMenu() == 0 {
				break
			}
			break
		case 7:
			deleteStudent()
			if backToMenu() == 0 {
				break
			}
			break
		case 8:
			studentInformation()
			if backToMenu() == 0 {
				break
			}
			break
		case 9:
			allStudentsInformation()
			if backToMenu() == 0 {
				break
			}
			break
		case 10:
			clearTerminal()
			//   putDataIntoFileClasses()
			//   putDataIntoFileStudents()
			os.Exit(1)
			break
		default:
			userNumber = choseOperation()
			break
		}
	}
}
