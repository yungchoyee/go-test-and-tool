package utime

import "time"

var MonthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

//标准时间戳转成200601格式日期
func Unix2YearMonth(timestamp int64) int32 {
	tm := time.Unix(timestamp, 0)
	y, M, _ := tm.Date()
	m := MonthMap[M.String()]
	return int32(y)*100 + int32(m)
}

//获取两个时间间的年月日期格式列表
func Unix2YearMonthList(st, et int64) (list []int32) {
	if st > et {
		et, st = st, et
	}
	endMonth := Unix2YearMonth(et)
	startMonth := Unix2YearMonth(st)
	list = make([]int32, 0)
	list = append(list, startMonth)
	for startMonth != endMonth {
		if startMonth%100 == 12 {
			//12月份特殊处理，年份加1，月份-11变成1月份
			startMonth = startMonth + 100 - 11
		} else {
			//其余月份加1
			startMonth = startMonth + 1
		}
		list = append(list, startMonth)
	}
	return
}
