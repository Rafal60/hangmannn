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
	nbrand := (len(mot) / 2) - 1
	for nbrand != 0 {
		n := rand.IntN(len(mot) - 1)
		if motcacher[n] == "_" {
			motcacher[n] = ToUpper(string(mot[n]))
			nbrand--
		}
	}
	fmt.Println(mot, motcacher)

}

func ToUpper(s string) string {
	h := []rune(s)
	result := ""
	for i := 0; i < len(h); i++ {
		if (h[i] >= 'a') && (h[i] <= 'z') {
			h[i] = h[i] - 32
		}
		result += string(h[i])
	}
	return result
}
