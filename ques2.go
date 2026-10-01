// package main

//  type Student struct {
// 		Name string
// 		Id int
// 		Marks int
// 		feesPaid bool
// 	 }

// func main() {

// 	 studentsName := []string {"Kaka", "kuku", "Kiki", "Keke"}
//    stu1 := Student {
// 	Name: "Kaka" ,
// 	Id: 202601,
// 	Marks: 348,
// 	feesPaid: true,

//    }

// }

package main

import "fmt"

type Student struct {
	Name     string
	Age      int
	Course   string
	FeesPaid bool
}

func printStudents(students []Student) {

	for _, student := range students {
		fmt.Println("Name:", student.Name)
		fmt.Println("Age:", student.Age)
		fmt.Println("Course:", student.Course)
		fmt.Println("Fees Paid:", student.FeesPaid)

		switch student.Course {
		case "CSE":
			fmt.Println("Course Category: Computer Science")
		case "ECE":
			fmt.Println("Course Category: Electronics")
		case "ME":
			fmt.Println("Course Category: Mechanical")
		default:
			fmt.Println("Course Category: Other Course")
		}

		fmt.Println()
	}
}

func main() {

	students := []Student{
		{Name: "Palak", Age: 18, Course: "CSE", FeesPaid: true},
		{Name: "Rahul", Age: 19, Course: "ECE", FeesPaid: false},
		{Name: "Aman", Age: 18, Course: "ME", FeesPaid: true},
		{Name: "Riya", Age: 20, Course: "CSE", FeesPaid: false},
	}

	printStudents(students)

	fees := map[string]int{
		"Paid":     0,
		"Not Paid": 0,
	}

	for _, student := range students {
		if student.FeesPaid {
			fees["Paid"]++
		} else {
			fees["Not Paid"]++
		}
	}

	fmt.Println("Fees Summary:")
	fmt.Println("Paid:", fees["Paid"])
	fmt.Println("Not Paid:", fees["Not Paid"])
}
