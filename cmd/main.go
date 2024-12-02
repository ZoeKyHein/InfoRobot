package main

import (
	"InfoRobot/internal"
	"github.com/robfig/cron/v3"
	"log"
)

// 主程序
func main() {
	c := cron.New()

	regions := []string{
		"北京", "上海",
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
