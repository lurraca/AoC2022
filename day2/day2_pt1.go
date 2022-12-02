package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var shapeToScorePlayer = map[byte]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var roundToResult = map[string]string{
	"A X": "DRAW",
	"A Y": "WIN",
	"A Z": "LOSS",
	"B X": "LOSS",
	"B Y": "DRAW",
	"B Z": "WIN",
	"C X": "WIN",
	"C Y": "LOSS",
	"C Z": "DRAW",
}

const (
	winScore  = 6
	lossScore = 0
	drawScore = 3
)

func main() {
	var totalScore int

	input := readInput("input.txt")

	for _, entryFromInput := range input {
		roundResult := roundToResult[entryFromInput]
		totalScore = totalScore + calculateScore(playerShapeToByte(entryFromInput), roundResult)
	}
	fmt.Printf("Total score is everything goes exactly according to the strategy guide: %v\n", totalScore)
}

func calculateScore(playerShape byte, roundResult string) (score int) {
	switch {
	case roundResult == "DRAW":
		score = shapeToScorePlayer[playerShape] + drawScore
		break
	case roundResult == "WIN":
		score = shapeToScorePlayer[playerShape] + winScore
		break
	case roundResult == "LOSS":
		score = shapeToScorePlayer[playerShape] + lossScore
		break
	}

	return score
}

func playerShapeToByte(round string) byte {
	arrStrRound := strings.Split(round, " ")

	playerShape := []byte(arrStrRound[1])[0]
	return playerShape
}

func readInput(fileName string) (entries []string) {

	inputFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries
}
