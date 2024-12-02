package main

import (
	"InfoRobot/browser"
	"context"
	"fmt"
	"github.com/ZoeKyHein/bot/wx"
	"github.com/go-rod/rod"
	"github.com/robfig/cron/v3"
	"log"
)

const BaseURL = `beijing.chinatax.gov.cn`

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
			if i > 4 {
				break
			}
			msg = append(msg, MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   BaseURL + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	return msg, nil

}

// 爬虫任务
func SpiderTask() []MsgSt {
	url := "http://beijing.chinatax.gov.cn/bjswj/c104176/Information.shtml" // 替换为目标网站
	data, err := FetchData(url)
	if err != nil {
		log.Printf("爬取失败: %v", err)
		return nil
	}

	// 保存数据（可以替换为数据库存储）
	log.Printf("爬取到的数据: %s", data)
	return data
}

// 主程序
func main() {
	c := cron.New()

	// 每天三次任务
	_, err := c.AddFunc("0 9 * * *", SpiderAndMessage) // 9:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("0 12 * * *", SpiderAndMessage) // 12:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("54 14 * * *", SpiderAndMessage) // 15:30
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	c.Start()
	log.Println("爬虫开始运行...")

	// 主程序运行
	select {} // 阻塞主线程
}

func SpiderAndMessage() {
	msgSts := SpiderTask()
	SendMessage(msgSts)

}

func SendMessage(msgSts []MsgSt) {
	botClient := wx.BotClient{
		Key:    "ee556a46-a3a7-4978-a186-7e3181f29da9",
		Logger: nil,
	}
	msg := ""
	for i, st := range msgSts {
		msg += fmt.Sprintf("%d. 「%s」[%s](%s)\n", i+1, st.Date, st.Title, st.Url)
	}
	err := botClient.SendMarkDown(context.Background(), msg)
	if err != nil {
		log.Printf("发送消息失败: %v", err)
	}
}
