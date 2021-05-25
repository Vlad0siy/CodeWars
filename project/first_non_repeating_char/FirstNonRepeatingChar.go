package first_non_repeating_char

import (
  "strings"
  "unicode"
  )

func FirstNonRepeating(str string) string {
  if len(str) == 1 {
    return string(str[0])
  }
  for i := 0; i <= len(str) - 1; i++ {
    if unicode.IsLower(rune(str[i])) {
      if strings.Count(strings.ToLower(str), string(str[i])) == 1 {
        return string(str[i])
      }
    } else {
      if strings.Count(strings.ToUpper(str), string(str[i])) == 1 {
        return string(str[i])
      }
    }
  }
  return ""
}
