package find_odd_init

func FindOdd(seq []int) int {
    counter := make(map[int] int)
    result := 0
    for _, val := range seq {
      _, bol := counter[val]
      if bol {
        counter[val]++
      } else {
        counter[val] = 1
      }
    }
    for key, val := range counter {
      if val % 2 != 0 {
        result = key
      }
    }
  return result
}
