package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
)

func main() {
	manageFile()
}

func manageFile() {
	var match bool = false

	for !match {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select your action: (create / read / write / delete) file ? [C,R,W,D]")
		
		text, _ := reader.ReadString('\n')

		if len(strings.TrimSpace(text)) > 0 {
			match = true
			switch strings.TrimSpace(text) {
				case "C":
					createFile()
				case "R":
					readFile()
				case "W":
					writeFile()
				case "D":
					deleteFile()
				default:
					match = false
					fmt.Println("No matched result ! Try again")
			}
		} else {
			fmt.Println("Please select one to continue")
		}
	}
}

func createFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name :")
	fmt.Print("=> ")

	fileName, _ := reader.ReadString('\n')
	
	os.Create(strings.TrimSpace(fileName))
	fmt.Printf("File %v created successfully !", fileName)
}

func readFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name :")
	fmt.Print("=> ")

	fileName, _ := reader.ReadString('\n')
	
	content, _ := ioutil.ReadFile(strings.TrimSpace(fileName))
	fmt.Println(string(content))
}

func writeFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name :")
	fmt.Print("=> ")

	fileName, _ := reader.ReadString('\n')

	fmt.Println(fileName)
}

func deleteFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter file name :")
	fmt.Print("=> ")

	fileName, _ := reader.ReadString('\n')
	
	os.Remove(strings.TrimSpace(fileName))
	fmt.Printf("The file %v removed successfully !", fileName)
}