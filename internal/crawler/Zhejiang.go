package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"strings"
	"time"
)

type ZhejiangCrawler struct{}

func NewZhejiangCrawler() *ZhejiangCrawler {
	return &ZhejiangCrawler{}
}

func (b *ZhejiangCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.ZhejiangInfoUrl)
		p.MustWaitStable()
		table := p.MustElements(`[class="common-list-items"]`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			tmpUrl := *element.MustElement("a").MustAttribute("href")
			if !strings.Contains(tmpUrl, "http") {
				tmpUrl = browser.ZhejiangBaseUrl + tmpUrl
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   tmpUrl,
			})
		}
	})
	data := models.Data{
		Region: "Zhejiang",
		Msgs:   msg,
	}
	return data, nil
}
