package main

import (
  "fmt"
)

func main() {
  fmt.Println("Enter text: ")
  num := 0
  fmt.Scanln(&num)
  if num > 0 {
    fmt.Println(rec_fact(num))
    for_fact(num)
  } else {
    fmt.Println("Needed num must be > 0")
  }
}

func for_fact(a int) {
  result_for := 1
  for i := 1; i <= a; i++ {
    result_for *= i
  }
  fmt.Println(result_for)
}

func rec_fact(a int) int {
  if a > 0 {
    return a * rec_fact(a - 1)
  } else {
    return 1
  }
}
