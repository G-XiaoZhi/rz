package dao

import (
	"rz/models"
)

// 基于开始日期和结束日期查询这段时间的daily数据
func GetDailyQfqByInterval(startDate, endDate string) (qfqList []*models.DailyQfq){
	DB.Table("daily_qfq").Where("trade_date >= ? and trade_date <= ?", startDate, endDate).Find(&qfqList)
	return
}

// 基于日期获取当天的数据
func GetDailyQfqByDate(calDate string) (qfqList []*models.DailyQfq){
	DB.Table("daily_qfq").Where("trade_date = ?", calDate).Find(&qfqList)
	return
}