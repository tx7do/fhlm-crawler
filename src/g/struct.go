package g

import (
	"fmt"
	"strings"
)

// LotteryIssueData 彩票期数数据
type LotteryIssueData struct {
	Issue string   `json:"issue"`
	Time  string   `json:"time"`
	Code  []string `json:"code"`
}

// CodeToString 开奖号码转字符串
func (d *LotteryIssueData) CodeToString() string {
	return strings.Join(d.Code, ",")
}

// Print 打印彩票信息
func (d *LotteryIssueData) Print() {
	fmt.Printf("期次:%s 开奖时间:%s 开奖号码:%s", d.Issue, d.Time, d.CodeToString())
}

// CrawlerData 采集数据
type CrawlerData struct {
	Title   string             `json:"title"`
	Lottery string             `json:"lottery"`
	Total   int                `json:"total"`
	List    []LotteryIssueData `json:"list"`
}

// GetLatestIssueData 获取最新一期的数据
func (d *CrawlerData) GetLatestIssueData() *LotteryIssueData {
	if d.List == nil {
		return nil
	}
	if len(d.List) == 0 {
		return nil
	}

	return &d.List[0]
}

// Print 打印采集信息
func (d *CrawlerData) Print() {
	fmt.Printf("%s 期次:%s 开奖时间:%s 开奖号码:%s\n", d.Title, d.GetLatestIssueData().Issue, d.GetLatestIssueData().Time, d.GetLatestIssueData().CodeToString())
}
