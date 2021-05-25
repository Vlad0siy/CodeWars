package roman_numerals_decoder

func Decode(roman string) int {
  var (
       result = 0
       new_val = 0
       old_val = 0
      )
  for i := 0; i < len(roman); i++ {
    if roman[i] == 'I' {
        new_val = 1
    } else if roman[i] == 'V' {
        new_val = 5
    } else if roman[i] == 'X' {
        new_val = 10
    } else if roman[i] == 'L' {
        new_val = 50
    } else if roman[i] == 'C' {
         new_val = 100
    } else if roman[i] == 'D' {
       new_val = 500
    } else if roman[i] == 'M' {
        new_val = 1000
    }
    if new_val > old_val {
      result = result + new_val - 2 * old_val
    } else {
      result = result + new_val
    }
    old_val = new_val
  }
return result
}
