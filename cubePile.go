package main

/*
  DESCRIPTION:
Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of the cube above will have volume of
  and so on until the top which will have a volume of
You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?

The parameter of the function findNb (find_nb, find-nb, findNb, ...) will be an integer m and you have to return the integern
Examples:
findNb(1071225) --> 45

findNb(91716553919377) --> -1
*/

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

func recursively(m, n int) int {
	nPart := ((n + 1) * n) / 2
	sum := nPart * nPart
	if sum == m {
		return n
	}
	if sum > m {
		return -1
	}
	return recursively(m, n+1)
}
