package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"

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
		GameState:      0,
	}
}




func startGame(filename string) {
	HangmanInit()
	Readword(filename)
	ChoseWord(filename)
	
	
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
func ChoseWord(filename string) {
	tw := Readword(filename)
	hangman.Word = tw[rand.Intn(len(tw)-1)]
	hangman.Word = hangman.Word[:len(hangman.Word)-1]
	hangman.WordHidden = wordToUnderscore()
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

func findAndReplace(letterToReplace string) {
hangman.Count++
	if len(letterToReplace) != 0 {
		for _,guess := range hangman.Guessedletter {
			if letterToReplace == guess {
				return 
			}
		}
		hangman.Guessedletter = append(hangman.Guessedletter, letterToReplace)
	}	

	if len(letterToReplace) > 1 {
		if letterToReplace == hangman.Word {
			print(2)
			hangman.WordHidden = hangman.Word
		} else {
			hangman.DeathCount -= 2
		}
		if hangman.DeathCount < 0 {
			hangman.DeathCount = 0
		}
		return
	}
	
	isFound := strings.Index(hangman.Word, letterToReplace)
	if isFound == -1 {
		if hangman.DeathCount >= 1 {
			hangman.DeathCount--
			//deathCountStage(hangman.DeathCount)
			fmt.Println("raté")
			fmt.Println("Il vous reste", hangman.DeathCount, "essais")
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
}

func testEndGame() {
	if hangman.WordHidden == hangman.Word {
		hangman.GameState = 1
	}
}

func testmot() bool {
	hangman.Count++
	// créer une var scanner qui va lire ce que l'utilisateur va écrire
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // l'utilisateur input dans la console
	// lis ce que l'utilisateur a écrit
	lettreoumot := scanner.Text()
	lettreoumot = strings.ToLower(lettreoumot)
	// peret à l'utilisateur de savoir qu'il ne doit mettre que des lettres contenues dans l'alphabet latin
	isALetter, err := regexp.MatchString("^[a-zA-Z]", lettreoumot)
	if Contains1(hangman.Guessedletter, lettreoumot) {
	} else {
		if err != nil {
			return testmot()
		}
		if !isALetter {
			return testmot()
		}
		if len(lettreoumot) == 1 {
			
			findAndReplace(lettreoumot)
		} else if lettreoumot == hangman.Word {
			return true
		} else if (len(lettreoumot) == len(hangman.Word)) && hangman.WordHidden == hangman.Word {
			return true
		} else {
			hangman.DeathCount -= 2
			//deathCountStage(hangman.DeathCount)
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


func GameState() {
	if testmot() || !Contains(hangman.WordHidden, '_') {
		hangman.GameState = 1
	}
	if deathCountStage() == 0 {
		hangman.GameState = 2
		}
	if testmot() || Contains(hangman.WordHidden, '_') {
		hangman.GameState = 0
		}
}

func Retry() {
	hangman.Count = 0
	hangman.DeathCount = 10
	hangman.Guessedletter = hangman.Guessedletter1
	
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
