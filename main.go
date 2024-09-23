package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Print("trop d'arguments")
		return
	}
	fichier, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fichier)
	fileScanner.Split(bufio.ScanLines)
	mots := []string{}
	for fileScanner.Scan() {
		mots = append(mots, fileScanner.Text())
	}
	fichier.Close()
	mot := mots[rand.IntN(len(mots)-1)]
	motcacher := []string{}
	for i := 0; i < len(mot); i++ {
		motcacher = append(motcacher, "_")
	}
	fmt.Println(mot, motcacher)
}
