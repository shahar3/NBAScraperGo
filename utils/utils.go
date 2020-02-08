package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Position string

const (
	Guard   Position = "Guard"
	Forward Position = "Forward"
	Center  Position = "Center"
)

const (
	FootToCm  float64 = 30.48
	InchToCm  float64 = 2.54
	PoundToKg float64 = 0.453592
)

func GetPosition(positions string) []Position {
	positionsArr := strings.Split(positions, "-")
	var convertedPosition []Position
	for _, position := range positionsArr {
		cvrt := Position(position)
		convertedPosition = append(convertedPosition, cvrt)
	}
	return convertedPosition
}

func GetHeightAndWeight(heightAndWeight string) (int, int) {
	fields := strings.Split(heightAndWeight, "|")
	foot, inches := getFeetAndInches(fields[0])
	cm := convertFeetAndInchesToCm(foot, inches)
	kg := convertPoundToKg(fields[1])
	return cm, kg
}

func convertFeetAndInchesToCm(feet string, inches string) int {
	feetInt, err := strconv.Atoi(feet)
	if err != nil {
		fmt.Println("Cannot convert feet to int, wrong input")
		return 0
	}
	inchInt, err := strconv.Atoi(inches)
	if err != nil {
		fmt.Println("Cannot convert inch to int, wrong input")
		return 0
	}
	cm := float64(feetInt)*FootToCm + float64(inchInt)*InchToCm
	cm = math.Round(cm)
	return int(cm)
}

func convertPoundToKg(pounds string) int {
	re := regexp.MustCompile("[0-9]+")
	poundsStr := re.FindString(pounds)
	poundsInt, err := strconv.Atoi(poundsStr)
	if err != nil {
		fmt.Println("Cannot convert Pound to int, wrong input")
		return 0
	}
	return int(math.Round(float64(poundsInt) * PoundToKg))
}

func getFeetAndInches(footAndInches string) (string, string) {
	fields := strings.Split(footAndInches, " ")
	feet := fields[0]
	inches := fields[2]
	return feet, inches
}

func GetDocument(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return dom
}

func BuildBoxScoreUrl(gameDate string, gameId string) string {
	return fmt.Sprintf("http://data.nba.net/10s/prod/v1/%s/%s_boxscore.json", gameDate, gameId)
}
