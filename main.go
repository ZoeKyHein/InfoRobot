package main

import (
	"InfoRobot/browser"
	"context"
	"github.com/go-rod/rod"
	"github.com/robfig/cron/v3"
	"log"
)

var InfoURLs = map[string]string{
	"北京": "http://beijing.chinatax.gov.cn/bjswj/c104176/Information.shtml",
}

type MsgSt struct {
	Title string
	Date  string
	Url   string
}

// 爬取网页内容
func FetchData(url string) ([]MsgSt, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	msg := make([]MsgSt, 0)
	defer lxrod.CloseBrowser()
	rod.Try(func() {

		p.MustNavigate(url)
		p.MustWaitStable()
		table := p.MustElement("div.xxgk_tzgg > ul")
		for i, element := range table.MustElements("li") {
			if i > 3 {
				break
			}
			msg = append(msg, MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	return msg, nil

}

// 爬虫任务
func SpiderTask() {
	url := "http://beijing.chinatax.gov.cn/bjswj/c104176/Information.shtml" // 替换为目标网站
	data, err := FetchData(url)
	if err != nil {
		log.Printf("爬取失败: %v", err)
		return
	}

	// 保存数据（可以替换为数据库存储）
	log.Printf("爬取到的数据: %s", data)
}

// 主程序
func main() {
	c := cron.New()

	// 每天三次任务
	_, err := c.AddFunc("0 9 * * *", SpiderTask) // 9:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("0 12 * * *", SpiderTask) // 12:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("54 14 * * *", SpiderTask) // 15:30
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	c.Start()
	log.Println("爬虫开始运行...")

	// 主程序运行
	select {} // 阻塞主线程
}
