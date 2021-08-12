package dao

import (
	"rz/models"
)

// 基于股票和日期查询db获取数据
func GetDailyBasicByCodeDate(tsCode, calDate string) (basic models.DailyBasic){
	//var qfq models.DailyQfq
	DB.Table("daily_basic").Where("ts_code = ? and trade_date = ?", tsCode, calDate).First(&basic)
	return
}

// 基于日期获取当天的数据
func GetDailyBasicByDate(calDate string) (basicList []*models.DailyBasic){
	DB.Table("daily_basic").Where("trade_date = ?", calDate).Find(&basicList)
	return
}