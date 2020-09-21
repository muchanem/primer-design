package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)
 
func main() {
	fmt.Println(reverse(complement(getFile("testing"))))
}
func reverse(strand string) string {
    r := []rune(strand)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
func complement(strand string) string {
	var returnValue string
	for _, char := range(strand) {
		switch string(char) {
		case "A":
			returnValue += "T"
		case "T":
			returnValue += "A"
		case "C":
			returnValue += "G"
		case "G":
			returnValue += "C"
		}
	}
	return returnValue
}

func meltingPoint(strand string) int {
	var A int
	var T int
	var C int
	var G int 
	for _, char := range(strand) {
		switch string(char) {
		case "A":
			A += 1
		case "T":
			T += 1
		case "C":
			C += 1
		case "G":
			G += 1 
		}
	}
		
	return ((2*(A + T) + 4*(C + G)) - 5)
}

func getFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	var returnValue string	
	check(err)
    for _, char := range strings.ToUpper(string(dat)) {
		if string(char) == "A" || string(char) == "T" || string(char) == "C" || string(char) == "G" {
			returnValue += string(char)
		}
	}
	return returnValue

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}