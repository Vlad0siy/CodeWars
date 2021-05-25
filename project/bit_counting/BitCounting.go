package bit_counting

import (
	"fmt"
	"strconv"
  "os"
)

func CountBIN (ABin string) int {
	result := 0
	num := 0
	for i := 0; i <= len(ABin) - 1; i++ {
		if ABin[i] != '0' {
			num, _ = strconv.Atoi(string(ABin[i]))
			result += num
		}
	}
	return result
}

func CountBits(ANum uint) int {
	result := make([]string, 0)
	str := ""
	for true {
		if ANum <= 1 {
			result = append(result, fmt.Sprint(ANum))
			break;
		} else {
			result = append(result, fmt.Sprint(ANum % 2))
			ANum = ANum / 2
		}
	}

	for i := len(result) - 1; i >= 0 ; i-- {
		str += result[i]
	}
  fmt.Println(os.Stdout, str)
	return CountBIN(str)
}

func main() {
  fmt.Println(os.Stdout, CountBits(10))
  fmt.Println(os.Stdout, CountBits(5))
  fmt.Println(os.Stdout, CountBits(3))
  fmt.Println(os.Stdout, CountBits(100))
  fmt.Scanln()
}
