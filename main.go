package main

import (
  "github.com/Vlad0siy/CodeWars/project/digital_root"
  "github.com/Vlad0siy/CodeWars/project/bit_counting"
  "github.com/Vlad0siy/CodeWars/project/first_non_repeating_char"
  "github.com/Vlad0siy/CodeWars/project/roman_numerals_decoder"
  "github.com/Vlad0siy/CodeWars/project/duplicate_encoder"
  "github.com/Vlad0siy/CodeWars/project/find_odd_init"
  "github.com/Vlad0siy/CodeWars/project/two_to_one"
  "github.com/Vlad0siy/CodeWars/project/arithmetic"
  "github.com/Vlad0siy/CodeWars/project/alternite_capitalization"
  "fmt"
  "os"
)

func main() {
  fmt.Println("---DigitalRoot---")
  fmt.Println(os.Stdout, digital_root.DigitalRoot(16))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(992))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(167346))
  fmt.Println(os.Stdout, digital_root.DigitalRoot(0))

  fmt.Println("---BitCounting---")
  fmt.Println(os.Stdout, bit_counting.CountBits(10))
  fmt.Println(os.Stdout, bit_counting.CountBits(5))
  fmt.Println(os.Stdout, bit_counting.CountBits(3))
  fmt.Println(os.Stdout, bit_counting.CountBits(100))

  fmt.Println("---FirstNonRepeating---")
  fmt.Println(os.Stdout, first_non_repeating_char.FirstNonRepeating("should handle simple tests"))
  fmt.Println(os.Stdout, first_non_repeating_char.FirstNonRepeating("stress"))
  fmt.Println(os.Stdout, first_non_repeating_char.FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"))
  fmt.Println(os.Stdout, first_non_repeating_char.FirstNonRepeating("~><#~><"))

  fmt.Println("---RomanNumeralsDecoder---")
  fmt.Println(os.Stdout, roman_numerals_decoder.Decode("MMVIII"))
  fmt.Println(os.Stdout, roman_numerals_decoder.Decode("MDCLXVI"))
  fmt.Println(os.Stdout, roman_numerals_decoder.Decode("I"))
  fmt.Println(os.Stdout, roman_numerals_decoder.Decode("IV"))

  fmt.Println("---DuplicateEncoder---")
  fmt.Println(os.Stdout, duplicate_encoder.DuplicateEncode("din"))
  fmt.Println(os.Stdout, duplicate_encoder.DuplicateEncode("recede"))
  fmt.Println(os.Stdout, duplicate_encoder.DuplicateEncode("(( @"))
  fmt.Println(os.Stdout, duplicate_encoder.DuplicateEncode("Success"))

  fmt.Println("---FindOddInit---")
  fmt.Println(os.Stdout, find_odd_init.FindOdd([]int{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5}))
  fmt.Println(os.Stdout, find_odd_init.FindOdd([]int{1,1,2,-2,5,2,4,4,-1,-2,5}))
  fmt.Println(os.Stdout, find_odd_init.FindOdd([]int{20,1,1,2,2,3,3,5,5,4,20,4,5}))
  fmt.Println(os.Stdout, find_odd_init.FindOdd([]int{1,1,1,1,1,1,10,1,1,1,1}))

  fmt.Println("---TwoToOne---")
  fmt.Println(os.Stdout, two_to_one.TwoToOne("aretheyhere", "yestheyarehere"))
  fmt.Println(os.Stdout, two_to_one.TwoToOne("loopingisfunbutdangerous", "lessdangerousthancoding"))
  fmt.Println(os.Stdout, two_to_one.TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
  fmt.Println(os.Stdout, two_to_one.TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"))

  fmt.Println("---Arithmetic---")
  fmt.Println(os.Stdout, arithmetic.Arithmetic(8, 2, "add"))
  fmt.Println(os.Stdout, arithmetic.Arithmetic(8, 2, "substract"))
  fmt.Println(os.Stdout, arithmetic.Arithmetic(8, 2, "multiply"))
  fmt.Println(os.Stdout, arithmetic.Arithmetic(8, 2, "divide"))

  fmt.Println("---AlterniteCapitalization---")
  fmt.Println(os.Stdout, alternite_capitalization.Capitalize("AbCdEf"))
  fmt.Println(os.Stdout, alternite_capitalization.Capitalize("CoDeWaRrIoRs"))
  fmt.Println(os.Stdout, alternite_capitalization.Capitalize("InDeXiNgLeSsOnS"))
  fmt.Println(os.Stdout, alternite_capitalization.Capitalize("CoDiNgIsAfUnAcTiViTy"))
}
