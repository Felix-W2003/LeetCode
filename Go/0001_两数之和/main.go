package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。

*/

/*
方法一：暴力枚举
复杂度分析:
时间复杂度：O(N*N)
空间复杂度: O(1)
*/
func twoSum_1(nums []int, target int) []int {
	var arryList = []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				arryList = append(arryList, i, j)
			}
		}
	}
	return arryList

}

/*
方法二：哈希表
方法一的时间复杂度较高的原因是寻找target-x的时间复杂度过高，因此我们需要能够
快速寻找数组中是否存在目标元素，如果存在，我们需要找出它的索引
使用哈希表，可以将寻找target-x的时间复杂度从O(N)降到O(1),
这样我们创建一个哈希表，对于每一个 x，我们首先查询哈希表中是否存在 target - x，
然后将 x 插入到哈希表中，即可保证不会让 x 和自己匹配。
*/
func twoSum_2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
func main() {
	test1 := []int{1, 1, 1, 4, 5, 6, 7, 8}
	test1_target := 2
	result1 := twoSum_1(test1, test1_target)
	if len(result1) == 0 {
		fmt.Println("没有找到")
	} else {
		fmt.Println(result1 == nil)
		for _, j := range result1 {
			fmt.Println(j, test1[j])
		}
	}
}
