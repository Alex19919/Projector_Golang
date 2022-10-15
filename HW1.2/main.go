// 2. Імітація лікаря. У вигляді опитування. Кожна відповідь веде до наступного запитання в терміналі

package main

import 	(
  "fmt"
)

  func main() {

fmt.Println("Hello?: ")

var greeting string
fmt.Scanln(&greeting)  //hi

fmt.Println("What is your name?: ")

var name string
fmt.Scanln(&name) //Alex

fmt.Println("What is your age?: ")

var age int
fmt.Scanln(&age) //31

fmt.Println("What is your weight?: ")

var weight float64
fmt.Scanln(&weight)  //100

fmt.Println("What health problems do you have?: ")

var problems string
fmt.Scanln(&problems)  //headache

fmt.Println("I will write you a prescription for medicine: ")

var prescriptionForMedicine string
fmt.Scanln(&prescriptionForMedicine)   //ok
}
