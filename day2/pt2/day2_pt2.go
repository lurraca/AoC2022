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

var desiredShapeResult = map[byte]string{
	'X': "LOSS",
	'Y': "DRAW",
	'Z': "WIN",
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
		opponentShape, desiredShape := shapesToByte(entryFromInput)

		desiredResult := desiredShapeResult[desiredShape]

		totalScore = totalScore + calculateScore(opponentShape, desiredShape, desiredResult)
	}
	fmt.Printf("Total score is everything goes exactly according to the strategy guide: %v\n", totalScore)
}

func calculateScore(opponentShape, desiredShape byte, desiredResult string) (score int) {
	switch {
	case desiredResult == "DRAW":
		score = shapeToScorePlayer[playerShape(opponentShape, desiredResult)] + drawScore
		fmt.Printf("opp: %c, des: %v, player: %c, score: %v\n",
			opponentShape, desiredResult, playerShape(opponentShape, desiredResult), score)
		break
	case desiredResult == "WIN":
		score = shapeToScorePlayer[playerShape(opponentShape, desiredResult)] + winScore
		fmt.Printf("opp: %c, des: %v, player: %c, score: %v\n",
			opponentShape, desiredResult, playerShape(opponentShape, desiredResult), score)
		break
	case desiredResult == "LOSS":
		score = shapeToScorePlayer[playerShape(opponentShape, desiredResult)] + lossScore
		fmt.Printf("opp: %c, des: %v, player: %c, score: %v\n",
			opponentShape, desiredResult, playerShape(opponentShape, desiredResult), score)
		break
	}

	return score
}

func playerShape(opponentShape byte, desiredResult string) (playerShape byte) {
	switch desiredResult {
	case "DRAW":
		if opponentShape == 'A' {
			playerShape = 'X'
		} else if opponentShape == 'B' {
			playerShape = 'Y'
		} else {
			playerShape = 'Z'
		}
		break
	case "WIN":
		if opponentShape == 'A' {
			playerShape = 'Y'
		} else if opponentShape == 'B' {
			playerShape = 'Z'
		} else {
			playerShape = 'X'
		}
		break
	case "LOSS":
		if opponentShape == 'A' {
			playerShape = 'Z'
		} else if opponentShape == 'B' {
			playerShape = 'X'
		} else {
			playerShape = 'Y'
		}
		break

	}
	return playerShape
}

func shapesToByte(round string) (byte, byte) {
	arrStrRound := strings.Split(round, " ")

	opponentShape := []byte(arrStrRound[0])[0]
	desiredShape := []byte(arrStrRound[1])[0]
	return opponentShape, desiredShape
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
