package kata

func FindNb(m int) int {
//   n := 0
//   sum := 0
//   for sum < m {
//     n++
//     sum =  sum + (n * n * n)
    
//   }
//   if sum != m {
//     return -1
//   }
//   return n
  
  //refactoring
  return recursively(m, 1)
}

func recursively(m , n int) int{
  nPart := ((n + 1) * n)/2
  sum := nPart * nPart
  if sum == m { return n }
  if sum > m { return -1}
  return recursively(m, n+1)
}
