package solution

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, val := range nums {
		complement := target - val
		if value, ok := hashmap[complement]; ok {
			return []int{value, i}
		}
		hashmap[val] = i
	}
	return []int{}
}
