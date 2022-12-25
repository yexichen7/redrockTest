package main

import "fmt"

func main() {
	var quality [][]int64
	var time []int64
	var frequency []int64
	var n int
	var term, t, q int64

	fmt.Scanf("%d ", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf(" %d , %d ", &t, &q)
		time = append(time, t)
		frequency = append(frequency, q)
		sl := make([]int64, 0, q)
		var j int64
		for j = 0; j < q; j++ {
			fmt.Scanf(" %d ", &term)
			sl = append(sl, term)
		}
		quality = append(quality, sl)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			re := RemoveRepByLoop(quality[i])
			fmt.Println(len(re))
		} else {
			if time[i]-time[i-1] < 86400 {
				var temp []int64
				var j int64
				for j = 0; j < frequency[i]; j++ {
					temp = append(temp, quality[i][j])
				}
				for j = 0; j < frequency[i-1]; j++ {
					temp = append(temp, quality[i-1][j])
				}
				re := RemoveRepByLoop(temp)
				fmt.Println(len(re))
			} else {
				var temp []int64
				var j int64
				for j = 0; j < frequency[i]; j++ {
					temp = append(temp, quality[i][j])
				}
			}
		}
	}

}

// 通过两重循环过滤重复元素,求出有多少种品质
func RemoveRepByLoop(slc []int64) []int64 {
	result := []int64{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
