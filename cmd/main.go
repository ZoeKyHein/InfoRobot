package main

import (
	"InfoRobot/browser"
	"InfoRobot/internal"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// 主程序
func main() {
	browser.Today = time.Now().Format("2006-01-02")

	c := cron.New()

	regions := []string{
		"北京",
		"上海",
		"天津",
		"大连",
		"重庆",
		"青岛",
		"宁波",
		"厦门",
		"深圳",
		"内蒙古",
		"安徽",
		"福建",
		"甘肃",
		"广东",
		"广西",
		"贵州",
		"海南",
		"河北",
		"黑龙江",
		"河南",
		"湖北",
		"湖南",
		"江苏",
		"江西",
		"吉林",
		"辽宁",
		"宁夏",
		"青海",
		"陕西",
		"山东",
		"山西",
		"四川",
		"新疆",
		"西藏",
		"云南",
		"浙江",
	}

	// 每天三次任务
	_, err := c.AddFunc("0 9 * * *", func() {
		internal.SpiderAndSend(regions)
	}) // 9:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("0 12 * * *", func() {
		internal.SpiderAndSend(regions)
	}) // 12:00
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	_, err = c.AddFunc("30 17 * * *", func() {
		internal.SpiderAndSend(regions)
	}) // 15:30
	if err != nil {
		log.Fatalf("定时任务添加失败: %v", err)
	}

	c.Start()
	log.Println("爬虫开始运行...")

	// 主程序运行
	select {} // 阻塞主线程
}
