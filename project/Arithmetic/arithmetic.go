package arithmetic

func Arithmetic(a int, b int, operator string) int{
  result := 0
  switch operator {
    case "add": result = a + b
    case "subtract": result = a - b
    case "multiply": result = a * b
    case "divide": result = a / b
    default: result = 0
  }
  return result
}
