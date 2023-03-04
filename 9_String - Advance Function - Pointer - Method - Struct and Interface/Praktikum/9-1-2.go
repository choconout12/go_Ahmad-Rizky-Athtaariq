package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func (s *Student) AddScore(score int) {
	s.score = append(s.score, score)
}

func (s *Student) AverageScore() float64 {
	sum := 0
	for _, score := range s.score {
		sum += score
	}
	return float64(sum) / float64(len(s.score))
}

func (s *Student) MinScore() int {
	min := s.score[0]
	for _, score := range s.score {
		if score < min {
			min = score
		}
	}
	return min
}

func (s *Student) MaxScore() int {
	max := s.score[0]
	for _, score := range s.score {
		if score > max {
			max = score
		}
	}
	return max
}

func main() {
	students := make([]Student, 0, 5)
	for i := 0; i < 5; i++ {
		fmt.Printf("Input data for student %d:\n", i+1)
		var name string
		var score int
		fmt.Print("Name: ")
		fmt.Scanln(&name)
		fmt.Print("Score: ")
		fmt.Scanln(&score)
		students = append(students, Student{name, []int{score}})
	}

	// Print average score, min score, and max score for all students
	var sum float64
	var min, max int
	for i, student := range students {
		fmt.Printf("Student %d - Name: %s, Scores: %v\n", i+1, student.name, student.score)
		sum += student.AverageScore()
		if i == 0 {
			min = student.MinScore()
			max = student.MaxScore()
		} else {
			if student.MinScore() < min {
				min = student.MinScore()
			}
			if student.MaxScore() > max {
				max = student.MaxScore()
			}
		}
	}
	fmt.Printf("Average score: %.2f\n", sum/float64(len(students)))
	fmt.Printf("Minimum score: %d\n", min)
	fmt.Printf("Maximum score: %d\n", max)
}
