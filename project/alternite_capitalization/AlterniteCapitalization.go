package alternite_capitalization

import "strings"

func Capitalize(st string) []string {
  var (
    str1 = ""
    str2 = ""
  )
  for i := 0; i < len(st); i++ {
    if i % 2 == 0 {
      str1 = str1 + strings.ToUpper(string(st[i]))
      str2 = str2 + string(st[i])
    } else {
      str2 = str2 + strings.ToUpper(string(st[i]))
      str1 = str1 + string(st[i])
    }
  }
  return []string{str1, str2}
}
