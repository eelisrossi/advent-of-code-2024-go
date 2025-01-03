package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)


func getInput(day int) string {
    inputDir := "inputs"
    inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

    if _, err := os.Stat(inputFile); err == nil {
        content, _ := os.ReadFile(inputFile)
        return string(content)
    }

    sessionCookie := os.Getenv("SESSION")

    url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
    req, _ := http.NewRequest("GET", url, nil)

    req.AddCookie(&http.Cookie{
        Name: "session",
        Value: sessionCookie,
    })

    client := &http.Client{}
    resp, _ := client.Do(req)

    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    _ = os.MkdirAll(inputDir, os.ModePerm)
    _ = os.WriteFile(inputFile, body, os.ModePerm)

    return string(body)
}

func getTestInput(day int) string {
    inputFile := fmt.Sprintf("./inputs/day%d-testinput.txt", day)
    content, _ := os.ReadFile(inputFile)
    return string(content)
}

func main() {

    godotenv.Load(".env")

    day := 6
    input := getInput(day)
    // input := getTestInput(day)

    p1 := day06_1(input)
    p2 := day06_2(input)

    fmt.Printf("Result for day %d part 1 is: %d\n", day, p1)
    fmt.Printf("Result for day %d part 2 is: %d\n", day, p2)
}

