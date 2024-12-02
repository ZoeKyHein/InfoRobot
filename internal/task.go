package internal

import (
	"InfoRobot/internal/crawler"
	"InfoRobot/internal/sender"
	"InfoRobot/models"
	"fmt"
	"log"
	"sync"
)

func SpiderAndMessage(region string) models.Data {

	cr, err := crawler.GetCrawler(region)
	if err != nil {
		log.Printf("获取爬虫失败: %v", err)
		return models.Data{}
	}

	data, err := cr.FetchData()
	if err != nil {
		log.Printf("爬取失败: %v", err)
		return models.Data{}
	}

	msg := fmt.Sprintf("「%s」\n", data.Region)
	for i, st := range data.Msgs {
		msg += fmt.Sprintf("%d. 「%s」[%s](%s)\n", i+1, st.Date, st.Title, st.Url)
	}

	return data
}

func SpiderAndSend(regions []string) {
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	msgMap := make(map[string]models.Data)

	for _, region := range regions {
		wg.Add(1)
		go func(region string) {
			defer wg.Done()
			msgs := SpiderAndMessage(region)

			mu.Lock()
			msgMap[region] = msgs
			mu.Unlock()
		}(region)
	}

	wg.Wait()

	finalMessage := ""
	for region, msgs := range msgMap {
		finalMessage += fmt.Sprintf("**『%s』**\n", region)
		for i, st := range msgs.Msgs {
			finalMessage += fmt.Sprintf("%d. **[%s]** [%s](%s)\n", i+1, st.Date, st.Title, st.Url)
		}
		finalMessage += "========================\n"
	}
	sender.SendMessage(finalMessage)
}
