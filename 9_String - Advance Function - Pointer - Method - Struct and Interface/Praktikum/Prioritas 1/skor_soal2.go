package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) averageScore() float64 {
	totalScore := 0.0
	for _, score := range s.score {
		totalScore += float64(score)
	}
	return totalScore / float64(len(s.name))
}

func (s Student) scoreMethod() {
	var totalSiswa int
	fmt.Print("Ingin memasukkan berapa siswa : ")
	fmt.Scanln(&totalSiswa)
	fmt.Println()

	for i := 1; i <= totalSiswa; i++ {
		var score int
		fmt.Print("Input ", i, " Student's Name : ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		name := scanner.Text()
		fmt.Print("Input ", i, " Student's Score : ")
		fmt.Scanln(&score)
		s.name = append(s.name, name)
		s.score = append(s.score, score)
		fmt.Println()
		if i == 5 || i%5 == 0 {
			maxScore := s.score[0]
			minScore := s.score[0]
			maxName := s.name[0]
			minName := s.name[0]
			totalScore := 0
			for j := 0; j < i; j++ {
				totalScore += s.score[j]
				if s.score[j] < minScore {
					minScore = s.score[j]
					minName = s.name[j]
				}
				if s.score[j] > maxScore {
					maxScore = s.score[j]
					maxName = s.name[j]
				}
			}
			avgScore := s.averageScore()
			fmt.Println("Average Score: ", avgScore)
			fmt.Println("Min Score of Students :", minName, "(", minScore, ")")
			fmt.Println("Max Score of Students : ", maxName, "(", maxScore, ")")

		}
	}

}

func main() {
	var student = Student{}
	student.scoreMethod()
}
