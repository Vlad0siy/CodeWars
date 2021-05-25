package main

import (
  "github.com/Vlad0siy/CodeWars/project/digital_root"
  "github.com/Vlad0siy/CodeWars/project/bit_counting"
  "github.com/Vlad0siy/CodeWars/project/first_non_repeating_char"
  "fmt"
  "os"
)

func main() {
  fmt.Println("---DigitalRoot---")
  fmt.Println(os.Stdout, digital_root.DigitalRoot(16))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(195))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(992))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(167346))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(0))

  fmt.Println("---BitCounting---")
  fmt.Println(os.Stdout, bit_counting.CountBits(10))
  fmt.Println(os.Stdout, bit_counting.CountBits(5))
  fmt.Println(os.Stdout, bit_counting.CountBits(3))
  fmt.Println(os.Stdout, bit_counting.CountBits(100))
}
