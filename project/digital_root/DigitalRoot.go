package digital_root

import (
  "strconv"
  "fmt"
  "os"
)

func DigitalRoot(n int) int {
  var (
  sum = 0
  )
  if len(strconv.FormatInt(int64(n), 10)) > 1 {
    for _, num := range strconv.FormatInt(int64(n), 10) {
      num, _ := strconv.Atoi(string(num))
      sum += num
    }
    return DigitalRoot(sum)
  } else {
    return n
  }
}

func main () {
    fmt.Println(os.Stdout, DigitalRoot(16))
    fmt.Println(os.Stdout, DigitalRoot(195))
    fmt.Println(os.Stdout, DigitalRoot(992))
    fmt.Println(os.Stdout, DigitalRoot(167346))
    fmt.Println(os.Stdout, DigitalRoot(0))
    fmt.Scanln()
}
