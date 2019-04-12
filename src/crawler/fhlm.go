package fhlm

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"

	"../db"
	"../g"
)

// come from https://www.fhlm.com/lotteries.html
// https://www.fhlm.com/_data/jlffc.json
const (
	baseURL = "https://www.fhlm.com/_data/"
)

// StartCrawler 启动采集
func StartCrawler(lotteryCode string) {

	url := fmt.Sprintf("%s%s.json", baseURL, lotteryCode)

	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		//fmt.Println("Visiting", r.Request.URL)

		data := &g.CrawlerData{}
		err := json.Unmarshal(r.Body, data)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//fmt.Println(data)
		processCrawlData(data)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 5 * time.Second,
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal("Visit Error: ", err)
	}

	c.Wait()
}

// processCrawlData 处理采集数据
func processCrawlData(data *g.CrawlerData) {
	// 输出结果
	data.Print()

	db.WriteToFile(data)
}
