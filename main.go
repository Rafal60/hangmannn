package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Print("trop d'arguments")
		return
	}
	mot := ""
	motcacher := []string{}
	nbessais := 0
	lutil := []string{}
	position := pospendu()
	d := 0
	difficulté := ' '
	for d == 0 {
		fmt.Print("Choisir une difficulté :")
		fmt.Scanf("%c\n", &difficulté)
		if difficulté == 'f' {
			d++
			mot = ouvrirfich("facile.txt")
			motcacher = motcache(mot)
			nbrand := (len(mot) / 2) - 1
			for nbrand != 0 {
				n := rand.IntN(len(mot) - 1)
				if motcacher[n] == "_" {
					motcacher[n] = ToUpper(string(mot[n]))
					nbrand--
				}
			}
			nbessais = 10
		} else if difficulté == 'm' {
			d++
			mot = ouvrirfich("moyen.txt")
			motcacher = motcache(mot)
			motcacher[0] = ToUpper(string(mot[0]))
			nbessais = 8
		} else if difficulté == 'd' {
			d++
			mot = ouvrirfich("difficile.txt")
			motcacher = motcache(mot)
			nbessais = 5
		} else {
			fmt.Print("Caractère invalide\n")
			continue
		}
	}
	fmt.Printf("\n \n Tu as %d essais pour trouver le bon mot\n", nbessais)
	fmt.Print("\t \tBonne chance\n \n")
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
			if string(mot[i]) == lettre && string(motcacher[i]) == "_" {
				motcacher[i] = ToUpper(lettre)
				danslemot++
			}
		}
		if !InTab(motcacher, "_") {
			continue
		} else if danslemot == 0 {
			nbessais--
			if nbessais == 0 {
				continue
			}
			fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", nbessais)
			fmt.Print(string(position[10-nbessais-1]))
			affichemot(motcacher)
		} else {
			fmt.Printf("Il te reste %d essais pour trouver le bon mot\n", nbessais)
			affichemot(motcacher)
		}
	}
	if nbessais == 0 {
		fmt.Print(string(position[9]))
		fmt.Printf("Perdu !! Le mot était : %s", ToUpper(mot))
	} else {
		fmt.Print("bravo vous avez trouvez le mot cacher : " + ToUpper(mot))
	}
}

func ouvrirfich(fich string) string {
	fichier, err := os.Open(fich)
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
	return mot
}

func motcache(mot string) []string {
	motcacher := []string{}
	for i := 0; i < len(mot); i++ {
		motcacher = append(motcacher, "_")
	}
	return motcacher
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

func pospendu() []string {
	tab := []string{}
	fichier, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Print(err)
	}
	fileScanner := bufio.NewScanner(fichier)
	fileScanner.Split(bufio.ScanLines)
	pos := ""
	lscan := 0
	for fileScanner.Scan() {
		if lscan < 0 {
			lscan++
			continue
		}
		pos += fileScanner.Text() + "\n"
		lscan++
		if lscan%7 == 0 {
			tab = append(tab, pos)
			pos = ""
			lscan -= 8
		}
	}
	fichier.Close()
	return tab
}
