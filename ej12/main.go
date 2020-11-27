package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile()
	readFile2()
	recordFile()
	recordFile2()
}

func readFile() {
	file, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Error")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			register := scanner.Text()
			fmt.Printf("Linea > " + register + "\n")
		}
	}
	file.Close()
}

func recordFile() {
	file, err := os.Create("./newFile.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}
	fmt.Fprintln(file, "THis a new line, perfect")
	file.Close()
}

func recordFile2() {
	fileName := "./newFile.txt"
	if Append(fileName, "\n This is a second line") == false {
		fmt.Println("Error in the second line")
	}
}

func Append(file string, text string) bool {
	fil, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error")
		return false
	}
	_, err = fil.WriteString(text)
	if err != nil {
		fmt.Println("Error")
		return false
	}
	return true
}
