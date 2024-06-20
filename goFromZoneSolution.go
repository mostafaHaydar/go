package main

import "fmt"

// import "unicode"

// #1
// func LatinAlphabet() {
// 	for i := 'a'; i < 'z'; i++ {
// 		fmt.Printf("%c", i)
// 	}
// }
// #2
// func printReverseAlphabet(){
// 	for i := 'z'; i > 'a'; i-- {
// 		fmt.Printf("%c",i)
// 	}
// }
// #3
// func numbers() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d", i)
// 	}
// }
// #4
// func IsNegative(nb int) {
// 	if(nb>0){
// 		fmt.Printf("posistive")
// 	}else{
// 		fmt.Printf("negative")
// 	}
// }
// #5
// func PrintComb() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			for x := 0; x < 10; x++ {
// 				if i < j && j < x {
// 					fmt.Printf("%d%d%d", i, j, x)
// 				}
// 			}
// 		}
// 	}
// }
// #6
// func PrintComb2() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if i < j {
// 				fmt.Printf("%d%d, ", i, j)
// 			}
// 		}
// 	}
// }
// #7
// func PrintNbr(n int) {
// 	fmt.Printf("%d", n)
// }
// #8
// func PrintCombN(x int) {
// 	n := 0
// 	index := -1
// 	var generateComb func(int, int, int)
// 	generateComb = func(index int, n int, number int) {
// 		if number == 0 {
// 			return
// 		}
// 		for j := index + 1; j < 10; j++ {
// 			n = n*10 + j
// 			generateComb(j, n, number-1)
// 			if number == 1 {
// 				fmt.Printf("%d ,", n)
// 			}
// 			n = (n - j) / 10
// 		}
// 	}
// 	generateComb(index, n, x)
// }
// #9
// func PointOne(n *int) {
// 	*n = -1
// }
// #10
// func UltimatePointOne(n ***int) {
// 	***n = -1
// }
// #11
// func DivMod(a int, b int, div *int, mod *int) {
// 	*div = a / b
// 	*mod = a % b
// }
// #12
// func UltimateDivMod(a *int, b *int) {
// 	*a = *a / *b
// 	*b = *a * *b % *b
// }
// #13
// func PrintStr(s string) {
// 	slice := []rune(s)
// 	for i := 0; i < len(slice); i++ {
// 		fmt.Printf("%c\n", slice[i])
// 	}
// }
// #14
// func StrLen(s string) int {
// 	return len(s)
// }
// #15
// func Swap(a *int, b *int) {
// 	tmp := *a
// 	*a = *b
// 	*b = tmp
// }
// #16
// func StrRev(s string) string {
// 	slice := []rune(s)
// 	var newSlice []rune
// 	for i := len(slice)-1; i >= 0; i-- {
// 		newSlice = append(newSlice, slice[i])
// 	}
// 	return string(newSlice)
// }
// #17
// func BasicAtoi(s string) int {
// 	slice:=[]rune(s)
// 	number:=0
// 	for i := 0; i < len(slice); i++ {
// 		if (unicode.IsDigit(slice[i])){
// 			switch expression {
// 			case condition:
// 				we use switch just for escaping the atoi  built in function
// 			}
// 		}
// 	}
// }
// #18
// the same as above function just a very samll change on it
// #19
// the same as above function just a very samll change on it  exactly adding the sign on the number
// func getThesmallOneAndRemoveIt(table []int) int {
// 	// Find the index of the smallest element
// 	min := 0
// 	for i := 0; i < len(table); i++ {
// 		if table[i] < table[min] {
// 			min = i
// 		}
// 	}

// 	// Remove the smallest element from the slice
// 	smallest := table[min]
// 	table = append(table[:min], table[min+1:]...)

// 	return smallest
// }

// func SortIntegerTable(table []int) {
// 	var newSlice []int
// 	for i := 0; i < len(table); i++ {
// 		newSlice = append(newSlice, getThesmallOneAndRemoveIt(table))
// 	}
// 	table = newSlice
// 	fmt.Println(table)
// }
// #20
// func IterativeFactorial(nb int) int {
// 	mod:=nb
// 	if(nb==0){
// 		return 1
// 	}
// 	return mod*IterativeFactorial(nb-1)
// }
// #21 and #22 and #23 are very simple and exactly like #20
// #24
// func Fibonacci(index int) int {
// 	if index == 0 || index == 1 {
// 		return index
// 	}
// 	return Fibonacci(index-1) + Fibonacci(index-2)
// }
// #25
// func Sqrt(nb int) int {
// 	sqrt := 0
// 	for i := 1; i < nb; i++ {
// 		if i*i == nb {
// 			sqrt = i
// 		}
// 	}
// 	return sqrt
// }
// #26 is very easy prime numbers
// #27 very easy also
// #28 eagth quen puzzle a bro very aeseu and doent help on the go lang
// #29
// func FirstRune(s string) rune {
// 	slice := []rune(s)
// 	return slice[0]
// }
// #30 the same as above  # 31 the same also
// #31
// func Compare(a, b string) int {
// 	state := 1
// 	firstString := []rune(a)
// 	secondString := []rune(a)
// 	if len(a) != len(b) {
// 		state = 0
// 		return state
// 	}
// 	for i := 0; i < len(a); i++ {
// 		if firstString[i] != secondString[i] {
// 			state = 0
// 			return state
// 		}
// 	}
// 	return state
// }
// #32 ..... exist here many athor question
// func MakeRange(min, max int) []int {
// 	var slice []int
// 	if min < max {
// 		for i := min; i < max; i++ {
// 			slice = append(slice, i)
// 		}
// 	}
// 	return slice
// }
// #33
// func ConcatParams(args []string) string {
// 	var slice []rune
// 	for i := 0; i < len(args); i++ {
// 		tmp := []rune(args[i])
// 		slice = append(slice, tmp...)
// 		tmp = []rune("\n")
// 		slice = append(slice, tmp...)
// 	}
// 	return string(slice)
// }
// #34 ....
// func IsPrime(nb int) bool {

// 	if nb%2 == 1 {
// 		return true
// 	}
// 	return false
// }
// func Map(f func(int) bool, a []int) []bool {
// 	var slice []bool
// 	for i := 0; i < len(a); i++ {
// 		slice = append(slice, IsPrime(a[i]))
// 	}
// 	return slice
// }
// #35
// func Rot14(s string) string {
// 	slice := []rune(s)
// 	for i := 0; i < len(s); i++ {
// 		slice[i] = slice[i] + 14
// 	}
// 	return string(slice)
// }
// #36

func main() {
}
