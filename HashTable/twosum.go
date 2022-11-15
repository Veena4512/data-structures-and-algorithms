/*
  Write a function that takes in a non-empty array of distinct integers and an
  integer representing a target sum. If any two numbers in the input array sum
  up to the target sum, the function should return them in an array, in any
  order. If no two numbers sum up to the target sum, the function should return
  an empty array.
  Sample Input: [2, 1, 3, -1, 11, 5, 4, 0] Target: 10
  Output: [-1 11]
*/
package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	m := make(map[int]int)
	var result []int
	for i := 0; i < len(array); i++ {
		required := target - array[i]
		if _, ok := m[required]; ok {
			result = append(result, required)
			result = append(result, array[i])
			return result
		} else {
			m[array[i]] = i
		}
	}
	return result
}

func main() {
	arr := []int{2, 1, 3, -1, 11, 5, 4, 0}
	msg := TwoNumberSum(arr, 10)
	fmt.Println(msg)
}