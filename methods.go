package main

import "fmt"
import "strings"

// methods
type PRS struct {
	id       int
	fullName string
}

func (prs PRS) prsInformation() {
	fmt.Printf("user id ::%d\nuser full name::%s", prs.id, prs.fullName)
}

func main() {
	slice := []string{"1", "2", "3"}
	fmt.Printf("%s", strings.Join(slice, "d"))
	prs := PRS{1, "mostaf"}
	prs.prsInformation()
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%s", slice[i])
	}
}

// interfaces
