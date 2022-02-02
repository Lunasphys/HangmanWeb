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
	deathCount int 
	count int 
	wordhidden string
	word string
	guessedletter []string
	guessedletter1 []string
	Answer string
}

var hangman Hangman

func HangmanInit() {
	hangman = Hangman{
		deathCount : 10,
		count : 0,
	}
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
		startGame("./fichiertxt/words.txt")
	case "2":
		startGame("./fichiertxt/words2.txt")
	case "3":
		startGame("./fichiertxt/words3.txt")
	}
}

func startGame(filename string) *Hangman {
	tw := Readword(filename)
	hangman.word = tw[rand.Intn(len(tw))]

	hangman.wordhidden = wordToUnderscore()

	for {
		fmt.Println(hangman.wordhidden)
		if testmot() || !Contains(hangman.wordhidden, '_') {
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

	input := hangman.word

	result := sampleRegexp.ReplaceAllString(input, "_")
	return (string(result))
}

func findAndReplace(letterToReplace string) string {
	isFound := strings.Index(hangman.word, letterToReplace)
	if isFound == -1 {
		if hangman.deathCount > 1 {
			hangman.deathCount--
			deathCountStage(hangman.deathCount)
			fmt.Println("raté")
			fmt.Println("Il vous reste", hangman.deathCount, "essais")
			return hangman.wordhidden
			// mettre à jour le score
		}
		if hangman.deathCount == 1 {
			hangman.deathCount--
			deathCountStage(hangman.deathCount)
			displayLoseMessage()
			Retry()
		}
	} else {
		str3 := []rune(hangman.wordhidden)
		for i, lettre := range hangman.word {
			if string(lettre) == letterToReplace {
				str3[i] = lettre
				hangman.wordhidden = string(str3)
				fmt.Println(hangman.wordhidden)
			}
		}
	}
	return hangman.wordhidden
}

func testmot() bool {
	hangman.count++
	countPrint()
	fmt.Println("Veuillez saisir une lettre ou un mot")
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan() // l'utilisateur input dans la console
	// lis ce que l'utilisateur a écrit
	println(hangman.wordhidden)
	lettreoumot := scanner.Text()
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if Contains1(hangman.guessedletter, lettreoumot) {
		fmt.Println("vous avez utilisé les lettres :", hangman.guessedletter)
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
			hangman.guessedletter = append(hangman.guessedletter, lettreoumot)
			fmt.Println("vous avez utilisé les lettres :", hangman.guessedletter)
			findAndReplace(lettreoumot)
		} else if lettreoumot == hangman.word {
			return true
		} else if (len(lettreoumot) == len(hangman.word)) && hangman.wordhidden == hangman.word {
			return true
		} else {
			hangman.deathCount -= 2
			deathCountStage(hangman.deathCount)
			fmt.Println("Vous n'avez pas trouvé le bon mot")
			fmt.Println("Il vous reste", hangman.deathCount, "essais")
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

	if hangman.count == 1 {
		fmt.Println("------------", hangman.count, "er tour", "-------------")
	}
	if hangman.count > 1 {
		fmt.Println("------------", hangman.count, "ème tour", "-------------")
	}
}

func Retry() {
	hangman.count = 0
	hangman.deathCount = 10
	hangman.guessedletter = hangman.guessedletter1
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
		hangman.guessedletter = hangman.guessedletter1
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
	fmt.Println("Tu as découvert le bon mot en ", hangman.count, " essai")
	fmt.Println("Votre mot était: ", hangman.word)
	fmt.Println("Bravo, vous avez sauvé le pendu")
}

func displayLoseMessage() {
	fmt.Println()
	fmt.Println("Raté ! Tu n'as pas réussi à découvrir le mot")
	fmt.Println("Votre mot choisi était : ", hangman.word)
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
