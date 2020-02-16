package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

//Student :nodoc:
type Student struct {
	StudentID     string
	Name          string
	FinalScore    float32
	Grade         string
	MidScore      int
	SemesterScore int
	Attendance    int
}

func isAnumber(param string) (int, error) {
	result, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func theFinalScore(attendance, mid, semester int) float32 {
	return (0.2 * float32(attendance)) + (0.4 * float32(mid)) + (0.4 * float32(semester))
}

func gradeScore(finalScore float32) string {
	var grade string

	switch {
	case (finalScore >= 85 && finalScore <= 100):
		grade = "A"

	case (finalScore >= 76 && finalScore <= 84):
		grade = "B"

	case (finalScore >= 61 && finalScore <= 75):
		grade = "C"

	case (finalScore >= 46 && finalScore <= 60):
		grade = "D"

	case (finalScore >= 0 && finalScore <= 45):
		grade = "E"
	}

	return grade
}

func main() {
	var (
		studentCountScan string
		studentCount     int
		students             = make([]Student, 0)
		Pass             int = 0
		Fail             int = 0
	)

	fmt.Print("Input the number of students :")
	for {
		fmt.Scan(&studentCountScan)
		result, err := isAnumber(studentCountScan)
		if err != nil {
			fmt.Printf("%s not a number, please input number : ", err)
			continue
		}
		studentCount = result
		break

	}

	for i := 0; i < studentCount; i++ {
		var (
			studentID       string
			name            string
			finalScore      float32
			midScore        int
			semesterScore   int
			attendanceScore int
		)
		fmt.Println(">> Students number : ", i+1)
		fmt.Print("Student ID : ")
		fmt.Scanln(&studentID)
		fmt.Print("Name : ")
		fmt.Scanln(&name)
		fmt.Print("Mid Term Test Score : ")
		fmt.Scanln(&midScore)
		fmt.Print("Semester Test Score : ")
		fmt.Scanln(&semesterScore)
		fmt.Print("Attendance : ")
		fmt.Scanln(&attendanceScore)

		finalScore = theFinalScore(midScore, semesterScore, attendanceScore)
		grade := gradeScore(finalScore)

		if grade == "D" || grade == "E" {
			Fail++
		} else {
			Pass++
		}
		students = append(students, Student{StudentID: studentID, Name: name, MidScore: midScore, SemesterScore: semesterScore, Attendance: attendanceScore, Grade: grade, FinalScore: finalScore})

	}
	fmt.Println("Report :")
	print(students, studentCount, Pass, Fail)

}

func print(data []Student, totalStudents int, passingStudents int, failingStudents int) {
	t := table.NewWriter()

	for i := 0; i < len(data); i++ {
		t.AppendRow(table.Row{i + 1, data[i].StudentID, data[i].Name, data[i].FinalScore, data[i].Grade})
	}
	t.SetStyle(table.Style{
		Name: "Mystyle",
		Box: table.BoxStyle{
			BottomLeft:       "=",
			BottomRight:      "=",
			BottomSeparator:  "=",
			Left:             " ",
			LeftSeparator:    "=",
			MiddleHorizontal: "=",
			MiddleSeparator:  "=",
			MiddleVertical:   " ",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ",
			RightSeparator:   "=",
			TopLeft:          "=",
			TopRight:         "=",
			TopSeparator:     "=",
			UnfinishedRow:    " ~~~",
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"No", "Student ID", "Name", "Final Score", "Grade"})
	t.Render()
	fmt.Printf("Number of Students         : %d\n", totalStudents)
	fmt.Printf("Number of Passing Students : %d\n", passingStudents)
	fmt.Printf("Number of Failing Students : %d\n", failingStudents)

}
