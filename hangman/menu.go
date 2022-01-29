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

var deathCount int = 10
var count int = 0
var wordhidden string
var word string
var guessedletter []string
var guessedletter1 []string

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
		startGame("./fichiertxt/words.txt")
	case "2":
		startGame("./fichiertxt/words2.txt")
	case "3":
		startGame("./fichiertxt/words3.txt")
	}
}

func startGame(filename string) {
	tw := Readword(filename)
	word = tw[rand.Intn(len(tw))]

	wordhidden = wordToUnderscore()

	for {
		fmt.Println(wordhidden)
		if testmot() || !Contains(wordhidden, '_') {
			displayWinMessage()
			Retry()
		}
	}
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

	input := word

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func findAndReplace(letterToReplace string) string {
	isFound := strings.Index(word, letterToReplace)
	if isFound == -1 {
		if deathCount > 1 {
			deathCount--
			deathCountStage(deathCount)
			fmt.Println("raté")
			fmt.Println("Il vous reste", deathCount, "essais")
			return wordhidden
			// mettre à jour le score
		}
		if deathCount == 1 {
			deathCount--
			deathCountStage(deathCount)
			displayLoseMessage()
			Retry()
		}
	} else {
		str3 := []rune(wordhidden)
		for i, lettre := range word {
			if string(lettre) == letterToReplace {
				str3[i] = lettre
				wordhidden = string(str3)
				fmt.Println(wordhidden)
			}
		}
	}
	return wordhidden
}

func testmot() bool {
	count++
	countPrint()
	fmt.Println("Veuillez saisir une lettre ou un mot")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console
	// lis ce que l'utilisateur a écrit
	println(wordhidden)
	lettreoumot := scanner.Text()
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if Contains1(guessedletter, lettreoumot) {
		fmt.Println("vous avez utilisé les lettres :", guessedletter)
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
			guessedletter = append(guessedletter, lettreoumot)
			fmt.Println("vous avez utilisé les lettres :", guessedletter)
			findAndReplace(lettreoumot)
		} else if lettreoumot == word {
			return true
		} else if (len(lettreoumot) == len(word)) && wordhidden == word {
			return true
		} else {
			deathCount -= 2
			deathCountStage(deathCount)
			fmt.Println("Vous n'avez pas trouvé le bon mot")
			fmt.Println("Il vous reste", deathCount, "essais")
		}
	}
	return false
}

func deathCountStage(death int) {

	file, err := os.Open("./fichiertxt/hangman.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	index := 0

	var start int
	var end int

	if death == 9 {
		start = 0
		end = 7
	}
	if death == 8 {
		start = 8
		end = 15
	}
	if death == 7 {
		start = 16
		end = 23
	}
	if death == 6 {
		start = 24
		end = 31
	}
	if death == 5 {
		start = 32
		end = 39
	}
	if death == 4 {
		start = 40
		end = 47
	}
	if death == 3 {
		start = 48
		end = 55
	}
	if death == 2 {
		start = 56
		end = 63
	}
	if death == 1 {
		start = 64
		end = 71
	}
	if death == 0 {
		start = 72
		end = 79
	}
	for fileScanner.Scan() {
		if index >= start && index <= end {
			println(fileScanner.Text())
		}
		index++
	}

}

// Compte le nombre de tour
func countPrint() {
	if count == 1 {
		fmt.Println("------------", count, "er tour", "-------------")
	}
	if count > 1 {
		fmt.Println("------------", count, "ème tour", "-------------")
	}
}

func Retry() {
	count = 0
	deathCount = 10
	guessedletter = guessedletter1
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
		guessedletter = guessedletter1
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
	fmt.Println("Tu as découvert le bon mot en ", count, " essai")
	fmt.Println("Votre mot était: ", word)
	fmt.Println("Bravo, vous avez sauvé le pendu")
}

func displayLoseMessage() {
	fmt.Println()
	fmt.Println("Raté ! Tu n'as pas réussi à découvrir le mot")
	fmt.Println("Votre mot choisi était : ", word)
	fmt.Println("Vous essaierez de sauver le pendu une autre fois")
}
