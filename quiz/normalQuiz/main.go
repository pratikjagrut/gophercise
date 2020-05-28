package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
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

func showProblem(i int, c *int, p problem) {
	fmt.Printf("Q #%d : %s=", i+1, p.que)
	var ans string
	fmt.Scanf("%s\n", &ans)
	if ans == p.ans {
		*c++
	}
}

func main() {
	csvName := flag.String("probCsvPath", "problems.csv", "path of csv of quiz problems")
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

	var correct int

	for i, p := range problems {
		showProblem(i, &correct, p)
	}

	fmt.Printf("Score: %d out of %d correct.\n", correct, len(problems))
}
