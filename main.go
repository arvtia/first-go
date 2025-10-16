package main

import (
	"fmt"
	"math"
)

// const s string = "constant"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// func vals() (int, int) {
// 	return 3, 7
// }

func sum(num ...int) {
	fmt.Println(num, " ")
	total := 0

	for _, num := range num {
		total += num
	}
	fmt.Println(total)
}

// // closures
// func intSeq() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

type Person struct {
	Name  string
	Age   int
	Email string
}

func newPerson(name string) *Person {

	p := Person{Name: name}
	p.Age = 22
	return &p
}

type rect struct {
	width, height float64
}

// func (r *rect) area() int {
// 	return r.width * r.height
// }

// func (r rect) perim() int {
// 	return 2*r.width + 2*r.height
// }

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCirlce(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

type serverState int

const (
	StateIdle serverState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[serverState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss serverState) String() string {
	return stateName[ss]
}

func transition(s serverState) serverState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected:
		return StateError
	case StateError:
		return StateRetrying
	case StateRetrying:
		return StateIdle
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {

	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	// r := rect{width: 3, height: 8}
	// c := circle{radius: 5}
	// measure(r)
	// measure(c)

	// detectCirlce(c)
	// detectCirlce(r)

	// interfaces in go -
	//  interfaces are named collection of methods signatures

	// r := rect{width: 10, height: 20}
	// fmt.Println("area: ", r.area())

	// fmt.Println("perim: ", r.perim())
	//  methods

	// fmt.Println(Person{"bob", 29, "bob@gmail.com"})
	// fmt.Println(Person{"italic", 16, "itlaic@gmail.com"})

	// const s = "こんにちは"

	// fmt.Println("Len:", len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%x ", s[i])
	// }
	// fmt.Println()

	// // pointers
	// i := 1
	// fmt.Println("inital", i)

	// zeroval(i)
	// fmt.Println("zeroval", i)

	// zeroptr(&i)
	// fmt.Println("zeroptr", i)

	// fmt.Println("pointer:", &i)

	// range over built in types

	// nums := []int{2, 3, 4}
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println("sum:", sum)

	// for i, num := range nums {
	// 	if num == 3 {
	// 		fmt.Println("index: ", i)
	// 	}
	// }

	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	// 	fmt.Printf("%s -> %s\n", k, v)
	// }

	// for k := range kvs {
	// 	fmt.Println("key", k)

	// }

	// // variadic functions
	// sum(11, 44)
	// sum(3, 4, 5)

	// nums := []int{1, 2, 3, 4}
	// sum(nums...)

	// nextInt := intSeq()

	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// // multliple values return
	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)

	// d, _ := vals()
	// fmt.Println(d)

	// // functions

	// res := plus(1, 2)
	// fmt.Println(res)

	// res = plusPlus(1, 3, 55)
	// fmt.Println(res)

	//

	// // maps - in go

	// m := make(map[string]int)
	// m["k1"] = 7
	// m["k2"] = 13

	// fmt.Println("map:", m)
	// v1 := m["k1"]
	// fmt.Println("v1:", v1)

	// v3 := m["k3"]
	// fmt.Println("v3:", v3)

	// fmt.Println("len:", len(m))

	// delete(m, "k2")
	// fmt.Println(m)
	// m["k4"] = 22
	// fmt.Println("len:", len(m))
	// fmt.Println(m)

	// clear(m)

	// _, prs := m["k2"]
	// fmt.Println("prs:", prs)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)

	// n2 := map[string]int{"Foo": 1, "bar": 2}

	// if maps.Equal(n, n2) {
	// 	fmt.Println("n==n2")
	// }

	// // slices
	// var s []string
	// fmt.Println("initial:", s, s == nil, len(s) == 0)

	// s = make([]string, 4)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// s[3] = "d"
	// fmt.Println("Set:", s)
	// fmt.Println("Get:", s[2])
	// fmt.Println("Len:", len(s))

	// s = append(s, "e", "f")
	// fmt.Println("Append: ", s)

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy:", c)

	// l := s[2:4]
	// fmt.Println("sl1:", l)
	// l = s[:5] // everying excluding from 5 and above
	// fmt.Println("sl2:", l)
	// l = s[2:] // everying excluding 0 - 2
	// fmt.Println("sl3:", l)

	// t := []string{"g", "h", "i"}
	// fmt.Println("dcl", t)
	// t2 := []string{"j", "k", "l", "m", "n"}
	// if slices.Equal(t, t2) {
	// 	fmt.Println("t == t2")
	// } else {
	// 	fmt.Println("t != t2")
	// }

	// twoD := make([][]int, 5)
	// for i := range 5 {
	// 	innerlen := i + 1
	// 	twoD[i] = make([]int, innerlen)
	// 	for j := range innerlen {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	// // arrays
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)    // set the value of index 4 - also printing out the value of whole array
	// fmt.Println("get:", a[4]) // get the value of index 4
	// fmt.Println("len:", len(a))

	// b := [5]int{1, 3, 5, 7, 9}
	// fmt.Println("dlc:", b)

	// c := [...]int{1, 2, 4, 5, 6}
	// fmt.Println("idx:", c)

	// v := [...]int{3, 4: 500, 599}
	// fmt.Println(v)

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
