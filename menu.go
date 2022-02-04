package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

type Hangman struct {
	DeathCount     int
	Count          int
	WordHidden     string
	Word           string
	Guessedletter  []string
	Guessedletter1 []string
	Answer         string
	GameState      int
}

var hangman Hangman

func HangmanInit() {
	hangman = Hangman{
		DeathCount:     10,
		Count:          0,
		Guessedletter:  []string{},
		Guessedletter1: []string{},
		GameState : 0,
	}
	hangman.WordHidden = wordToUnderscore()
}

func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}

func SlowPrint(str ...string) {
	for _, strpart := range str {
		for _, char := range strpart {
			fmt.Print(string(char))

			time.Sleep(40_000_000 * time.Nanosecond)
		}
	}
}

func start() {
	fmt.Print("Bonjour et bienvenue dans le jeu du pendu\n")
	fmt.Println("1 = Démarrer l'éxécution")
	fmt.Println("2 = Non, je ne souhaite tuer personne")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		Clear()
		debut()
	case "2":
		os.Exit(2)
	}
}

func debut() {
	fmt.Println("Quelle bibliothèque de mot souhaitez vous choisir ? ")
	fmt.Println("1 = Choisir la premiere version")
	fmt.Println("2 = Choisir la deuxieme version")
	fmt.Println("3 = Choisir la troisieme version")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		startGame("../words.txt")
	case "2":
		startGame("../words2.txt")
	case "3":
		startGame("../words3.txt")
	}
}

func startGame(filename string) {
	HangmanInit()
	tw := Readword(filename)
	hangman.Word = tw[rand.Intn(len(tw))]

	hangman.WordHidden = wordToUnderscore()
	/*
		for {
			fmt.Println(hangman.WordHidden)
			if testmot() || !Contains(hangman.WordHidden, '_') {
				displayWinMessage()
				Retry()
			}
		}*/
	// trouve le mot et transforme le mot choisi en underscore

}

func Readword(filename string) []string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	word1 := ""
	var todowordos []string
	for _, char := range string(file) {
		if char == '\n' {
			todowordos = append(todowordos, word1)
			word1 = ""
		} else {
			word1 += string(char)
		}
	}
	return todowordos
}

func wordToUnderscore() string {
	sampleRegexp := regexp.MustCompile("[a-z,A-Z]")

	input := hangman.Word

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func findAndReplace(letterToReplace string) string {
	
	if len(letterToReplace) > 1 {
		if letterToReplace == hangman.Word {
			hangman.WordHidden = hangman.Word
		} else {
			hangman.DeathCount -= 2
		}
		if hangman.DeathCount < 0 {
			hangman.DeathCount = 0
		}
		return ""
	}
		isFound := strings.Index(hangman.Word, letterToReplace)
	if isFound == -1 {
		if hangman.DeathCount >= 1 {
			hangman.DeathCount--
			//deathCountStage(hangman.DeathCount)
			fmt.Println("raté")
			fmt.Println("Il vous reste", hangman.DeathCount, "essais")
			return hangman.WordHidden
			// mettre à jour le score
		}
		
	} else {
		str3 := []rune(hangman.WordHidden)
		for i, lettre := range hangman.Word {
			if string(lettre) == letterToReplace {
				str3[i] = lettre
				hangman.WordHidden = string(str3)
			}
		}
	}
	return hangman.WordHidden
}

func testmot() bool {
	hangman.Count++
	countPrint()
	fmt.Println("Veuillez saisir une lettre ou un mot")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console
	// lis ce que l'utilisateur a écrit
	println(hangman.WordHidden)
	lettreoumot := scanner.Text()
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if Contains1(hangman.Guessedletter, lettreoumot) {
		fmt.Println("vous avez utilisé les lettres :", hangman.Guessedletter)
		fmt.Println("vous avez deja rentré cette lettre")
	} else {
		if err != nil {
			fmt.Printf("Malheureusement cela ne marche pas ")
			fmt.Printf("Partir %v", lettreoumot)
			return testmot()
		}
		if !isALetter {
			fmt.Printf("Ce n'est pas une lettre !\n")
			return testmot()
		}
		if len(lettreoumot) == 1 {
			hangman.Guessedletter = append(hangman.Guessedletter, lettreoumot)
			fmt.Println("vous avez utilisé les lettres :", hangman.Guessedletter)
			findAndReplace(lettreoumot)
		} else if lettreoumot == hangman.Word {
			return true
		} else if (len(lettreoumot) == len(hangman.Word)) && hangman.WordHidden == hangman.Word {
			return true
		} else {
			hangman.DeathCount -= 2
			//deathCountStage(hangman.DeathCount)
			fmt.Println("Vous n'avez pas trouvé le bon mot")
			fmt.Println("Il vous reste", hangman.DeathCount, "essais")
		}
	}
	return false
}

func deathCountStage() int {

	file, err := os.Open("../hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	index := 0
	var death int
	var start int
	var end int
	if death == 9 {
		start = 0
		end = 7
		return 9
	}
	
	if death == 8 {
		start = 8
		end = 15
		return 8
	}
	if death == 7 {
		start = 16
		end = 23
		return 7
	}
	if death == 6 {
		start = 24
		end = 31
		return 6
	}
	if death == 5 {
		start = 32
		end = 39
		return 5
	}
	if death == 4 {
		start = 40
		end = 47
		return 4
	}
	if death == 3 {
		start = 48
		end = 55
		return 3
	}
	if death == 2 {
		start = 56
		end = 63
		return 2
	}
	if death == 1 {
		start = 64
		end = 71
		return 1
	}
	if death == 0 {
		start = 72
		end = 79
		return 0
	}
	for fileScanner.Scan() {
		if index >= start && index <= end {
			println(fileScanner.Text())
		}
		index++
	}
	return index
}

// Compte le nombre de tour
func countPrint() {

	if hangman.Count == 1 {
		fmt.Println("------------", hangman.Count, "er tour", "-------------")
	}
	if hangman.Count > 1 {
		fmt.Println("------------", hangman.Count, "ème tour", "-------------")
	}
}

func GameState() {
	if testmot() || !Contains(hangman.WordHidden, '_') {
		hangman.GameState = 1
		}
	if deathCountStage() == 0 {
		hangman.GameState = 2
		}
}

func Retry() {
	hangman.Count = 0
	hangman.DeathCount = 10
	hangman.Guessedletter = hangman.Guessedletter1
	SlowPrint("Voulez vous recommencer? \n")
	fmt.Println("1 = Oui")
	fmt.Println("2 = Non")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console

	// lis ce que l'utilisateur a écrit
	o := scanner.Text()
	switch o {
	case "1":
		Clear()
		hangman.Guessedletter = hangman.Guessedletter1
		debut()
	case "2":
		os.Exit(2)
	}
	if !Contains(o, '1') || !Contains(o, '2') {
		fmt.Println("Veuillez saisir une des réponses proposées")
		Retry()
	}
}

func displayWinMessage() {
	fmt.Println()
	fmt.Println("Tu as découvert le bon mot en ", hangman.Count, " essai")
	fmt.Println("Votre mot était: ", hangman.Word)
	fmt.Println("Bravo, vous avez sauvé le pendu")
}

func displayLoseMessage() {
	fmt.Println()
	fmt.Println("Raté ! Tu n'as pas réussi à découvrir le mot")
	fmt.Println("Votre mot choisi était : ", hangman.Word)
	fmt.Println("Vous essaierez de sauver le pendu une autre fois")
}

func Contains(s string, char rune) bool { // Si une string est contenue dans un tableau
	for _, a := range s {
		if a == char {
			return true
		}
	}
	return false
}

func Contains1(s []string, char string) bool { // Si une string est contenue dans un tableau
	for _, a := range s {
		if a == char {
			return true
		}
	}
	return false
}
