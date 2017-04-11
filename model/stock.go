package model

type stock struct {
	code             string  //代码
	name             string  //名称
	industry         string  //所属行业
	area             string  //地区
	pe               float64 //市盈率
	outstanding      float64 //流通股本(亿)
	totals           float64 //总股本(亿)
	totalAssets      float64 //总资产(万)
	liquidAssets     float64 //流动资产
	fixedAssets      float64 //固定资产
	reserved         float64 //公积金
	reservedPerShare float64 //每股公积金
	esp              float64 //每股收益
	bvps             float64 //每股净资
	pb               float64 //市净率
	timeToMarket     string  //上市日期
	undp             float64 //未分利润
	perundp          float64 // 每股未分配
	rev              float64 //收入同比(%)
	profit           float64 //利润同比(%)
	gpr              float64 //毛利率(%)
	npr              float64 //净利润率(%)
	holders          float64 //股东人数
}

func (s *stock) timeShare(key string) {

}

// 业绩报告
func (s *stock) Report() {

}
