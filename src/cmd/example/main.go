package main

import "fmt"

func updateValue(someVal *int) {
  *someVal = *someVal + 100
}

func main() {
  val := 1000
  fmt.Println("Value before: ", val)
  updateValue(&val)
  fmt.Println("Value after updateValue: ", val)
}
