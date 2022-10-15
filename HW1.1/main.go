// 1. Конфігурування свого авто. Ви вказуєте конкретні параметри, а в результаті виконання програми отримуєте, наприклад, детальний опис автівки
package main

import 	"fmt"

type Car struct {
	Brand string
	Model string
	Year int
	Engine float64
	color string
}

func main () { 
	myCar := Car {"Ford", "Escape", 2015, 2.5, "Grey"}
		
	if myCar.Brand == "Ford" {
		fmt.Println(myCar)
	} else {
		fmt.Println("It is not you car")
	}
}	