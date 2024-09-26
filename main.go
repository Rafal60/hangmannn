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
	lutil := []string{}
	nbrand := (len(mot) / 2) - 1
	for nbrand != 0 {
		n := rand.IntN(len(mot) - 1)
		if motcacher[n] == "_" {
			motcacher[n] = ToUpper(string(mot[n]))
			nbrand--
		}
	}
	nbessais := 10
	fmt.Printf("Tu as %d essais pour trouver le bon mot\n", nbessais)
	fmt.Print("Bonne chance\n")
	affichemot(motcacher)
	for nbessais != 0 {
		if MotFini(motcacher) {
			break
		}
		danslemot := 0
		l := ' '
		fmt.Print("Choisir une lettre :")
		fmt.Scanf("%c\n", &l)
		lettre := string(l)
		if !simplelettre(lettre) {
			fmt.Print("Caractère invalide\n")
			continue
		}
		if InTab(lutil, lettre) {
			fmt.Print("Lettre déjà utiliser réessayer\n")
			continue
		}
		lutil = append(lutil, lettre)
		for i := 0; i < len(mot); i++ {
			if string(mot[i]) == lettre {
				motcacher[i] = ToUpper(lettre)
				danslemot++
			}
		}
		if !InTab(motcacher, "_") {
			continue
		} else if danslemot == 0 {
			nbessais--
			fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", nbessais)
			affichemot(motcacher)
		} else {
			fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", nbessais)
			affichemot(motcacher)
		}
	}
	if nbessais == 0 {
		fmt.Printf("Perdu !! Le mot était : %s", ToUpper(mot))
	} else {
		fmt.Print("bravo vous avez trouvez le mot cacher : " + ToUpper(mot))
	}
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

func MotFini(tab []string) bool {
	for _, i := range tab {
		if i == "_" {
			return false
		}
	}
	return true
}

func InTab(tab []string, lettre string) bool {
	for _, i := range tab {
		if i == lettre {
			return true
		}
	}
	return false
}

func affichemot(tab []string) {
	affiche := ""
	for _, i := range tab {
		affiche += i + " "
	}
	fmt.Print(affiche + "\n")
}

func simplelettre(l string) bool {
	if len(l) != 1 {
		return false
	}
	if (rune(l[0]) < 'a' || rune(l[0]) > 'z') && (rune(l[0]) < 'A' || rune(l[0]) > 'Z') {
		return false
	}
	return true
}
