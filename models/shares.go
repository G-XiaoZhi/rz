package models

// 定义交易日期的ORM struct
type TradeCal struct {
	ID           uint   `gorm:"primaryKey"`
	Exchange     string `gorm:"->"` // 只读（除非有自定义配置，否则禁止写）
	CalDate      string `gorm:"->"`
	IsOpen       int    `gorm:"->"`
	PreTradeDate string `gorm:"->"`
}

// 定义前复权数据的ORM struct
type DailyQfq struct {
	ID           uint    `gorm:"primaryKey"`
	TsCode       string  `gorm:"->"` // 只读（除非有自定义配置，否则禁止写）
	TradeDate    string  `gorm:"->"`
	Open         float64 `gorm:"->"`
	High         float64 `gorm:"->"`
	Low          float64 `gorm:"->"`
	Close        float64 `gorm:"->"`
	PreClose     float64 `gorm:"->"`
	Change       float64 `gorm:"->"`
	PctChg       float64 `gorm:"->"`
	Vol          float64 `gorm:"->"`
	Amount       float64 `gorm:"->"`
	TurnoverRate float64 `gorm:"->"`
	VolumeRatio  float64 `gorm:"->"`
	Ma5          float64 `gorm:"->;column:ma5"`
	MaV5         float64 `gorm:"->;column:ma_v_5"`
	Ma10         float64 `gorm:"->;column:ma10"`
	MaV10        float64 `gorm:"->;column:ma_v_10"`
	Ma20         float64 `gorm:"->;column:ma20"`
	MaV20        float64 `gorm:"->;column:ma_v_20"`
	Ma30         float64 `gorm:"->;column:ma30"`
	MaV30        float64 `gorm:"->;column:ma_v_30"`
	Ma60         float64 `gorm:"->;column:ma60"`
	MaV60        float64 `gorm:"->;column:ma_v_60"`
	Ma120        float64 `gorm:"->;column:ma120"`
	MaV120       float64 `gorm:"->;column:ma_v_120"`
	Ma180        float64 `gorm:"->;column:ma180"`
	MaV180       float64 `gorm:"->;column:ma_v_180"`
	Ma240        float64 `gorm:"->;column:ma240"`
	MaV240       float64 `gorm:"->;column:ma_v_240"`
}

// 定义基础股票数据的ORM struct
type StockBasic struct {
	ID         uint   `gorm:"primaryKey"`
	TsCode     string `gorm:"->"` // 只读（除非有自定义配置，否则禁止写）
	Symbol     string `gorm:"->"`
	Name       string `gorm:"->"`
	Area       string `gorm:"->"`
	Industry   string `gorm:"->"`
	Fullname   string `gorm:"->"`
	Enname     string `gorm:"->"`
	Market     string `gorm:"->"`
	Exchange   string `gorm:"->"`
	CurrType   string `gorm:"->"`
	ListStatus string `gorm:"->"`
	ListDate   string `gorm:"->"`
	DelistDate string `gorm:"->"`
	IsHs       string `gorm:"->"`
}

// 定义基础股票数据的ORM struct
type DailyBasic struct {
	ID            uint    `gorm:"primaryKey"`
	TsCode        string  `gorm:"->"` // 只读（除非有自定义配置，否则禁止写）
	TradeDate     string  `gorm:"->"`
	Close         float64 `gorm:"->"`
	TurnoverRate  float64 `gorm:"->"`
	TurnoverRateF float64 `gorm:"->"`
	VolumeRatio   float64 `gorm:"->"`
	Pe            float64 `gorm:"->"`
	PeTtm         float64 `gorm:"->"`
	Pb            float64 `gorm:"->"`
	Ps            float64 `gorm:"->"`
	PsTtm         float64 `gorm:"->"`
	DvRatio       float64 `gorm:"->"`
	DvTtm         float64 `gorm:"->"`
	TotalShare    float64 `gorm:"->"`
	FloatShare    float64 `gorm:"->"`
	FreeShare     float64 `gorm:"->"`
	TotalMv       float64 `gorm:"->"`
	CircMv        float64 `gorm:"->"`
}
