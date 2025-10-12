package main

import (
	"fmt"
)

// const s string = "constant"

func main() {

	// arrays
	var a [5]int
	fmt.Println("emp:", a)

	

	// // // switch statements
	// i := 1
	// switch i {
	// case 1:
	// 	fmt.Println("number is one")
	// case 2:
	// 	fmt.Println("number is two")
	// default:
	// 	fmt.Println("number is greater than 2")
	// }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Today is weekend")
	// default:
	// 	fmt.Println("Today is weekday")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("good morning")
	// default:
	// 	fmt.Println("afternoon")
	// }

	//  whatAmI := func(i interface{}) {
	//      switch t := i.(type) {
	//      case bool:
	//          fmt.Println("I'm a bool")
	//      case int:
	//          fmt.Println("I'm an int")
	//      default:
	//          fmt.Printf("Don't know type %T\n", t)
	//      }
	//  }
	//  whatAmI(true)
	//  whatAmI(1)
	//  whatAmI("hey")

	// // if else if - conditions
	// if num:=9; num < 0 {
	// 	fmt.Println("number is negative")
	// } else if num < 10 {
	// 	fmt.Println("number is one digit")
	// } else {
	// 	fmt.Println("number is mutlidigits")
	// }

	// // for - loops
	// // simple loops
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// // classic - initialize, condition, and after
	// for j := 1; j <= 10; j++ {
	// 	fmt.Println(j)
	// }
	// // range - do this N times” iteration is range over an integer
	// for i := range 3 {
	// 	fmt.Println("range", i)
	// }

	// // simple range for - to print even number ina rnage upto 6
	// for n := range 6 {
	// 	if n == 0 {
	// 		continue
	// 	}
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// // using constants in go
	// fmt.Println(s)
	// const n = 5000000000
	// const d = 3e20 / n
	// fmt.Println(d)
	// fmt.Println(int64(d))
	// fmt.Println(math.Sin(n))

	// // printing - using fmt
	// fmt.Println(1 + 2)
	// fmt.Println( true && false)
	// fmt.Println(false || true)
	// fmt.Println("go" + "lang")

	// // vairable declaration
	// var num, num2 int = 1, 3
	// fmt.Println(num + num2)
	// var d, c, e = true, false, true
	// fmt.Println(d || c || e)
	// var z int
	// fmt.Println(z)
	// f := 2
	// fmt.Println(f)

	// constants in go
}
