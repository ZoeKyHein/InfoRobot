package internal

import (
	"InfoRobot/browser"
	"InfoRobot/internal/crawler"
	"InfoRobot/internal/sender"
	"InfoRobot/models"
	"fmt"
	"log"
	"strings"
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

	gr := groupRegions(regions, 4)
	for _, group := range gr {
		finalMessage := ""
		for _, region := range group {
			msgs := msgMap[region]
			finalMessage += fmt.Sprintf("**『%s』**\n", region)
			for i, st := range msgs.Msgs {
				if st.Date == browser.Today {
					finalMessage += fmt.Sprintf("%d. **<font color=\"warning\">「%s」</font>** [%s](%s)\n", i+1, st.Date, st.Title, st.Url)
				} else {
					finalMessage += fmt.Sprintf("%d. **「%s」** [%s](%s)\n", i+1, st.Date, st.Title, st.Url)
				}
			}
			finalMessage += "========================\n"
		}

		// 删除"国家税务总局",以节省消息长度
		finalMessage = strings.ReplaceAll(finalMessage, "国家税务总局", "")
		sender.SendMessage(finalMessage)
	}

}

// groupRegions 按指定数量分组地区
func groupRegions(regions []string, size int) [][]string {
	var groups [][]string
	for i := 0; i < len(regions); i += size {
		end := i + size
		if end > len(regions) {
			end = len(regions)
		}
		groups = append(groups, regions[i:end])
	}
	return groups
}
