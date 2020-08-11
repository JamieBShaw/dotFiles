package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getPackageList() {
	file, err := os.Open("packages-pacman.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("sudo pacman -S ", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {
	getPackageList()
}
