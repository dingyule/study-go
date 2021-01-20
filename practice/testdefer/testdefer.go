package main

import (
	"fmt"
	//"github.com/sirupsen/logrus"
	"io"
)

func main() {
	//function1()
	//doDBOperations()
	func1("Go")
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")

}

func function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

func func1(s string) (n int, err error) {

	defer func() {
		fmt.Printf("func1(%q) = %d, %v", s, n, err)
		//logrus.Infof("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF

}
