package leecode

import "strings"

//290单词规律
func WordPattern(pattern string, s string) (res bool) {
	//获取两个字符串数组
	wordArr := strings.Split(pattern, "")
	word1Arr := strings.Split(s, " ")
	if len(wordArr) != len(word1Arr) {
		return
	}
	//创建两个hashMap用来存储值
	wordMap := make(map[string]string)
	word1Map := make(map[string]string)

	//遍历规则字符串数组1
	for key, val := range wordArr {
		if sVal, ok := wordMap[val]; ok {
			//如果map1中存在该值
			pVal, ok1 := word1Map[sVal]
			//检查map中对应是否存在该值，不存在则不符合，且word1Map中该值等于遍历到的此时的值，wordMap中该值等于数组word1Arr的值
			if !ok1 || pVal != val || sVal != word1Arr[key] {
				return
			}
		} else {
			//不存在则加入hashMap
			wordMap[val] = word1Arr[key]
			//map1不存在，则map2是一一对应 关系则也不能存在
			if _, ok := word1Map[word1Arr[key]]; ok {
				return
			}
			word1Map[word1Arr[key]] = val
		}
	}
	return true
}

//714 买卖股票的最佳时机含手续费
func MaxProfit(prices []int, fee int) int {
	sell, buy := 0, -prices[0]
	for _, val := range prices {
		sell = MaxInt(sell, buy+val-fee)
		buy = MaxInt(buy, sell-val)
	}
	return sell
}
//prices := []int{1, 3, 2, 8, 4, 9,14,44,32,12,33,34,234}
func MaxProfit1(prices []int, fee int) (profit int) {
	if len(prices) < 2{
		return
	}
	min := prices[0]
	for _,val := range prices[1:]{
		if val < min{
			min = val
		}else if val > min + fee{
			profit += val - (min+fee)
			min = val - fee
		}
	}
	return
}

//maxInt return int
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//389 找不同   位运算 哈希表
func FindTheDifference(s string, t string) (difference byte) {
	for k := range s{
		difference ^= s[k] ^ t[k]
	}

	return difference ^ t[len(t) - 1]
}