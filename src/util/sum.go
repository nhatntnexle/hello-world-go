package util 

func Sum(numbers []int) int {
  sum := 0
  // This bug is intentional
  for n := range numbers {
    sum += n
  }
  return sum
}
