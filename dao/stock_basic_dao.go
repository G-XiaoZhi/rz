package dao

// 基于股票和日期查询db获取数据
func GetAllTsCode() (codeList []string){
	//var qfq models.DailyQfq
	DB.Table("stock_basic").Select("ts_code").Find(&codeList)
	return
}

// 基于日期获取当天的数据
func GetZbTsCode() (codeList []string){
	DB.Table("stock_basic").Select("ts_code").Where("market = ?", "主板").Find(&codeList)
	return
}