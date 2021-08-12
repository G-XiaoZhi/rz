package dao

import (
	"fmt"
	"rz/models"
)

func GetTradeCal(calDate string) (string, int) {
	var cal models.TradeCal
	DB.Table("trade_cal").Where("cal_date = ?", calDate).First(&cal)
	fmt.Println(cal.ID, cal.Exchange, cal.CalDate, cal.IsOpen, cal.PreTradeDate)
	return cal.CalDate, cal.IsOpen
}

func GetTradeCalListByInterval(startDate, endDate string) (calList []*models.TradeCal) {
	DB.Table("trade_cal").Where("cal_date >= ? and cal_date <= ? "+
		"and is_open = 1", startDate, endDate).Find(&calList)
	return calList
}
