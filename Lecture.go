package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Lecture() string {
	file, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("00 Lecture du fichier imposibe vérifier le nom et le type de fichier utiliser")
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close()

	return os.Args[1]
}
