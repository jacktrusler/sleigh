package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env variables
	err := godotenv.Load()
	if err != nil {
		// no env variables
	}
}

func FetchPuzzleInput(year, day int, session string) (str string, status int, err error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", 0, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", 0, err
	}
	defer resp.Body.Close()

	var bodyString string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body: ", err)
		return "", 0, err
	}
	bodyString = string(bodyBytes)

	fmt.Println("----- Fetching Input -----")
	fmt.Println("Status code:", resp.StatusCode)
	return bodyString, resp.StatusCode, nil
}

func main() {
	currentTime := time.Now()
	thisYear := currentTime.Year()
	month := currentTime.Month()
	today := currentTime.Day()
	if month != 12 {
		today = 1
	}

	envCookie := os.Getenv("SESSION_COOKIE")
	sessionCookie := flag.String("c", "", "Cookie for your unique Advent of Code session.")
	year := flag.Int("year", thisYear, fmt.Sprintf("Year of the problem (2015-%d).", thisYear))
	day := flag.Int("day", today, "Day of the problem (1-25).")
	name := flag.String("name", "", "Optionally set a filename. (default yyyy-dd-input.txt)")
	flag.Parse()

	if *sessionCookie == "" && envCookie == "" {
		fmt.Println("Requires a session cookie flag '-c', enter your session cookie to retrieve your input. \n\nOptionally, you can also create an .env file with: \nSESSION_COOKIE=<session_cookie>\nNote: AoC Session cookies last 30 days.")
		return
	}
	// 2015 was the first year
	if *year > thisYear || *year < 2015 {
		fmt.Printf("Year is out of range, please enter a year between %d and %d, if left blank, defaults to current year (%d).", 2015, thisYear, thisYear)
		return
	}
	if *day > 25 || *day < 1 {
		fmt.Printf("Day is out of range, please enter a day between 1 and 25, if left blank, defaults to today (%d) if run in the month of December.", today)
		return
	}

	fileName := *name
	if fileName == "" {
		dayFormatted := fmt.Sprintf("%02d", *day)
		fileName = fmt.Sprintf("%d-%s-input.txt", *year, dayFormatted)
	}

	var puzzleInput string
	var statusCode int
	var err error
	if envCookie != "" {
		puzzleInput, statusCode, err = FetchPuzzleInput(*year, *day, envCookie)
	} else {
		puzzleInput, statusCode, err = FetchPuzzleInput(*year, *day, *sessionCookie)
	}

	if err != nil {
		fmt.Println(err)
	}

	if statusCode == 400 {
		fmt.Printf("Response from adventofcode.com: %s\n", puzzleInput)
		fmt.Println("This indicates there was a problem with your SESSION_COOKIE")
		return
	}

	if statusCode == 200 {
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		file.WriteString(strings.Trim(puzzleInput, "\n"))

		fmt.Println("File created successfully:", fileName)
	}

}
