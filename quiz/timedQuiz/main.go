package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	que string
	ans string
}

func setProblems(lines [][]string) []problem {
	prblms := make([]problem, len(lines))
	for i, line := range lines {
		prblms[i] = problem{
			que: strings.TrimSpace(line[0]),
			ans: strings.TrimSpace(line[1]),
		}
	}
	return prblms
}

func main() {
	csvName := flag.String("probCsvPath", "problems.csv", "set path of csv of quiz problems")
	timeout := flag.Int("timeout", 20, "set timeout for quiz")
	flag.Parse()

	file, err := os.Open(*csvName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	problems := setProblems(lines)

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)

	var correct int
	fmt.Printf("Yov've got %v seconds to complete the quiz\n", *timeout)
	fmt.Println()

	for i, p := range problems {
		fmt.Printf("Q #%d : %s=", i+1, p.que)
		answer := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answer <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nScore: %d out of %d correct.", correct, len(problems))
			return
		case ans := <-answer:
			if ans == p.ans {
				correct++
			}
		}
	}
	fmt.Printf("Score: %d out of %d correct.\n", correct, len(problems))
}
