package duplicate_encoder

import "strings"

func DuplicateEncode(word string) string {
  var (
    result = ""
  )
  for i := 0; i <= len(word) - 1; i++ {
    if strings.Count(strings.ToLower(word), strings.ToLower(string(word[i]))) == 1 {
      result = result + "("
    } else {
      result = result + ")"
    }
  }
  return result
}
