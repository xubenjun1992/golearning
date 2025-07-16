package main

import (
	"fmt"
	"sort"
	"strconv"

	"main.go/gorm/blog"
)

/*
*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
*/
func getOnceNum(nums []int64) []int64 {
	numCount := make(map[int64]int64)
	for _, num := range nums {
		numCount[num]++
	}

	result := []int64{}

	for num, count := range numCount {
		if count == 1 {
			result = append(result, num)
		}
	}
	if len(result) == 0 {
		panic("no unique element found")
	}
	return result
}

/*
*
125. 验证回文串：给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，忽略字母的大小写。
*/
func palindrome(num int) bool {
	original := strconv.FormatInt(int64(num), 10)
	for i, j := 0, len(original)-1; i < j; i, j = i+1, j-1 {
		if original[i] != original[j] {
			return false
		}
	}
	return true
}

/*
*
20. 有效的括号：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
*/
func stackApp(str string) bool {
	stack := []rune{}
	mapping := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if len(stack) > 0 && stack[len(stack)-1] == mapping[char] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

// 最长公前缀
func sameLongest(str []string) string {
	s := str[0]
	if len(s) == 0 {
		return ""
	}
	for _, v := range str {
		for i := 0; i < len(s); i++ {
			if i < len(v) && s[i] != v[i] {
				s = s[:i]
				break
			}
		}
	}
	return s
}

// 交换数字
func swapData(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

// 交换数字
func swapData2(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// 给定一个数字，求多少个星期多少天
func getData(num int) (int, int) {
	week := num / 7
	day := num % 7
	return week, day
}

// 华氏温度转摄氏度
func temperatureConv(num float64) float64 {
	celsius := (num - 32) / 1.8
	return celsius
}

// 获取0-50之间整除2的数
func getEvenNumber() {
	for i := 0; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// 累加
func accrual(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

// 整除9的数之和及个数
func multipleFor9AndSum() (int, int) {
	sum := 0
	count := 0
	for i := 1; i <= 100; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	return count, sum
}

// 阶乘
func factorial(num int) int {
	factorial := 1
	for i := num + 1; i > 0; i-- {
		if i == 1 {
			return factorial
		}
		factorial *= i - 1
	}
	return factorial
}

// 打印矩形
func rectangle(hight int, width int) {
	for i := 0; i < hight; i++ {
		for j := 0; j < width; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
*
*删除排序数组中的重复项
 */
func removeDuplicates(num []int) (int, []int) {
	slow := 0
	for fast := 0; fast < len(num); fast++ {
		if num[fast] != num[slow] {
			slow++
			num[slow] = num[fast]
		}
	}
	resultNum := slow + 1
	return resultNum, num[:resultNum]
}

// 加1
func plusOne(num []int) []int {
	for i := len(num) - 1; i >= 0; i-- {
		if num[i] < 9 {
			num[i] += 1
			return num
		}
		num[i] = 0
	}

	//跳出循环说明其他位都是9
	newNum := make([]int, len(num)+1)
	newNum[0] = 1
	return newNum
}

/*
* 合并区间
 */
func mergeIntervals(num [][]int) [][]int {
	if len(num) == 1 {
		return num
	}
	sort.Slice(num, func(i, j int) bool {
		return num[i][0] < num[j][0]
	})

	result := [][]int{}
	current := num[0]
	for i := 1; i < len(num); i++ {
		next := num[i]
		if current[1] >= next[0] {
			current[1] = max(current[1], next[1])
		} else {
			result = append(result, current)
			current = next
		}
	}
	result = append(result, current)
	return result
}

/*
* 两数相加
 */
func twoSum(num []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if num[i]+num[j] == target {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func main() {
	// num := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(twoSum(num, 9))
	// samplesql.TransferAccounts(1, 2, 100)
	// blog.GenerateTable()
	blog.QueryPostsAndCommentsByUserId(1)
	blog.QueryMostCommentsPost()
}
