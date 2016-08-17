package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you call me on %s\n", h.name, h.phone)
}

func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

// Employee overload SayHi
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s.  Call me on %s\n", e.name, e.company, e.phone)
}

func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and ...)
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!
}

// define interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
