package main

import "fmt"

func main() {
	var score float64
	var nxtcal string
	for {
		fmt.Print("Please input the student's score: ")
		fmt.Scan(&score)
		if score >= 0 && score <= 100 {
			grade := gradecal(score)
			fmt.Println("The grade is ", grade)
			nxtcal = nextcal()
		} else {
			fmt.Println("Invalid score.")
			continue
		}

		if nxtcal == "n" || nxtcal == "N" {
			fmt.Println("Good bye.")
			break
		} else {
			continue
		}

	}

}

func gradecal(score float64) (grade string) {

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}
	return
}

func nextcal() (nxtcal string) {
	for {
		fmt.Print("Next calculation? [Y/N]")
		fmt.Scan(&nxtcal)
		if nxtcal == "Y" || nxtcal == "y" || nxtcal == "N" || nxtcal == "n" {
			return
		} else {
			fmt.Println("Invalid input.")
			continue
		}
	}
}
