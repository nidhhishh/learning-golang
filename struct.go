package main

import "fmt"

// student info struct

// type studentInfo struct {
// 	Name     string
// 	Age      int
// 	Course   string
// 	FeesPaid bool
// }

// //-------------------------------------------------------------------------------------------------------------

// func studentt(name string, age int, course string) *studentInfo {
// 	stu := studentInfo{
// 		Name:   name,
// 		Age:    age,
// 		Course: course,
// 	}
// 	return &stu
// }

// this func is a reciever type, to change the value using func, pointers used

// func (s *studentInfo) namechange(Name string) {
// 	s.Name = Name
// }

//-------------------------------------------------------------------------------------------------------------

//  to get the value of a specific field using func
// func (s studentInfo) courseget() string {
// 	return s.Course
// }

func main() {

	// 	stu1 := studentt("kaka", 13, "bca")
	// 	fmt.Println(stu1.Name)

	// student1 := studentInfo{
	// 	Name:     " daksh",
	// 	Age:      17,
	// 	Course:   "B.Tech",
	// 	FeesPaid: true,
	// }

	// student1.namechange("bauna")

	// student2 := studentInfo{
	// 	Name:     "Rambo",
	// 	Age:      19,
	// 	Course:   "bba",
	// 	FeesPaid: false,
	// }

	// student1.Age = 19
	// fmt.Println("studentInfo struct", student1)
	// fmt.Println("studentInfo struct", student2)

	//-------------------------------------------------------------------------------------------------------------
	// one time stuct

	ludo := struct {
		player1 string
		id      int
		win     bool
	}{"kuku", 234, true}

	fmt.Println(ludo)

}
