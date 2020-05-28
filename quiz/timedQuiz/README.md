# Timed quiz

This is a timer-based quiz. 

As soon as the timer expires quiz will end and will print result you scored.
The default time limit is 20 seconds, but the user can change via a flag `-timeout=40`.

Questions will be asked from the CSV file. 
The default CSV file is `problems.csv`. 
The user can provide another CSV filename via a flag `-probCsvPath=problems.csv` provide the format of CSV is same as given `problem.csv`, where the first column is a question and the second column in the same row is the answer to that question.

At the end of the quiz, the program prints the result. 
Questions given invalid answers are considered incorrect.

Output:

As soon as the timer expires result is printed.

```
go run main.go -timeout=5

You've got 5 seconds to complete the quiz

Q #1 : 5+5=10
Q #2 : 7-3=4
Q #3 : 1+1=
Score: 2 out of 13 correct.
```