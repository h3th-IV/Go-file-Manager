package main

import (
	"fmt"
	"os"
	"time"
)

func test() {
	var (
		Name       string = "Samuel Bonu"
		Address    string = "Block 10, Oremeji Street, HappyLand Estate Sangotede, Ajah. Lagos"
		Str_un     string = "string 1"
		Str_deu    string = "string 2"
		Str_tri    string = "string 3"
		Str_quatre string = "string 4"
	)

	fmt.Printf("Name: %s\n", Name)
	fmt.Printf("Address: %s\n", Address)
	fmt.Printf("some strings, %s, %s, %s, %s\n", Str_un, Str_deu, Str_tri, Str_quatre)
}

func test2() {
	var (
		firstName string
		LuckyNum  int
	)
	fmt.Print("Enter your first Name: ")
	fmt.Scanf("%s", &firstName)
	fmt.Printf("Your name is %s\n", firstName)

	fmt.Print("Enter your Lucky Number: ")
	fmt.Scanf("%d", &LuckyNum)
	fmt.Printf("You have selected %d as you Lucky Number...best of luck\n", LuckyNum)
	fmt.Println("Calculating....")
	time.Sleep(3 * time.Second)

	if LuckyNum >= 1 && LuckyNum <= 10 {
		fmt.Printf("Congrats you win with NUmber %d\n", LuckyNum)
	} else {
		fmt.Printf("YOU LOOSE :( %d\n", LuckyNum)
	}

}

func Age() {
	var (
		today = time.Now().Year()
		DOB   int
		age   int
	)
	fmt.Print("Year of Birth: ")
	fmt.Scanf("%d", &DOB)
	fmt.Println("Calculating...")
	time.Sleep(3 * time.Second)
	age = today - DOB
	fmt.Printf("you are %d years old", age)
}

func Weight() {
	var (
		ClassWeights = []float64{10, 20, 30, 40, 50}

		sum   = 0.0
		count = 0
	)

	for _, W := range ClassWeights {
		sum += W
		count++
	}

	// Calculate the average
	if count > 0 {
		average := sum / float64(count)
		fmt.Printf("Average: %.2f\n", average)
	} else {
		fmt.Println("No numbers to calculate the average.")
	}
}

// A local variable is said to be available only within the scope to which it is defined e.g a Function
// A Global variable is available and visible throughout the program
// You can make a variable global by using the Capital Notation i.e it's first Letter is a CAPITAL LETTER...
//...and also by declaring it in the global scope

func creatArray() {
	myArrayInt := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print(myArrayInt)

	myArrayString := [6]string{"Nat Love", "Mrs Mary", "Bass Revees", "Jim Beckworth", "Cufee", "Bill Pickett"}
	fmt.Println(myArrayString)
}

// for loops can exist in for loops, they are called nested loops

func counter() {
	for i := 1; i <= 10; i++ {
		print(i)
	}
}

// the range allwow us to iterate over items stored in a container such as slices,e.t.c
/*   the line: index, elements := range a iterate over both elemets and thier indexes and uses both values in the succeding lines
of code

	the line: _, elemnts :=..... uses a blank identifier inplace of indeces as the index may not or will not be used
*/

func divider() {
	var x int
	var result int
	fmt.Println("Enter a value for x: ")
	fmt.Scanf("%d", &x)

	if x >= 1 {
		result = x / 2
		fmt.Println(result)
	} else {
		fmt.Println("Zero Division error")
	}
}

//yes if statements can be Nested

// in go we do not have a keyword for while loops...
// but while loops run if a specified condition is true and for loops is used to do something for a
// specified number of times

type house struct {
	noRooms int
	price   int
	city    string
}

// a struct is mainly native and used for procedural programming while Classes are used in the concept of OOP
// a map is an object containing data in key value pairs
//no maps are not ordered

// a pointer is a variable that stores the location of a value in memory used especially in linked list

func filer() {

	filePath := "hercury/Documents/nano.txt"

	_, err := os.Stat(filePath) // get file info

	if err == nil {
		fmt.Printf("File '%s' exists.\n", filePath)
	} else if os.IsNotExist(err) {
		fmt.Printf("File '%s' does not exist.\n", filePath)
	} else {
		fmt.Printf("Error checking file: %v\n", err)
	}
}

func main() {
	test()
	test2()
	Age()
	Weight()
	divider()
	creatArray()
	filer()
}

/* i skipped questions i'm not familiar with :(
 */