package dao

import (
	"fmt"
	"rz/models"
	"testing"
)

func BenchmarkGetDailyQfqByInterval(b *testing.B) {
	// 测试1年周期的内存情况
	var qfqList []*models.DailyQfq

	startDate, endDate := "20200101", "20210101"

	qfqList = GetDailyQfqByInterval(startDate, endDate)

	fmt.Println(cap(qfqList))
}


func BenchmarkGetDailyQfqByDate(b *testing.B) {
	// 测试1年周期的内存情况
	var qfqList []*models.DailyQfq

	calList := GetTradeCalListByInterval("20200101","20210101")

	for _, cal := range calList {
		if cal.IsOpen == 1 {
			qfqList = GetDailyQfqByDate(cal.CalDate)
		}
	}

	fmt.Println(cap(qfqList))
}