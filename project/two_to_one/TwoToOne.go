package two_to_one

import (
  "strings"
  "sort"
)
func TwoToOne(s1 string, s2 string) string {
  main_str := ""
  result := ""
  str := []string{}
  if s1 != s2 {
    main_str = s1 + s2
  } else {
    main_str = s1
  }
  for i := 0; i <= len(main_str) - 1; i++ {
    if !strings.ContainsRune(result, rune(main_str[i])) {
      str = append(str, string(main_str[i]))
      result = result + string(main_str[i])
    }
  }
  sort.Strings(str)
  result = strings.Join(str, "")
  return result
}
