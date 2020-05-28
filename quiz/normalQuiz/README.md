# Normal Quiz

The program reads in a quiz provided via a CSV file and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong, the next question will be asked immediately afterwards.

The default CSV file is `problems.csv`, but the user can provide another CSV filename via a flag `-probCsvPath=problems.csv` but the format of CSV file should be same as given CSV file where the first column is a question and the second column in the same row is the answer to that question.

At the end of the quiz, the program outputs the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

Output: 

```
go run main.go
Q #1 : 5+5=10
Q #2 : 7-3=4
Q #3 : 1+1=2
Q #4 : 8*3=24
Q #5 : 1+2=3
Q #6 : 8/4=2
Q #7 : 3-1=2
Q #8 : 1*4=4
Q #9 : 5+1=6
Q #10 : 2+3=5
Q #11 : 3/3=3
Q #12 : 2*4=8
Q #13 : 5-2=3
Score: 12 out of 13 correct.
```