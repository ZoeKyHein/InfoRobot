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

type ShanghaiCrawler struct{}

func NewShanghaiCrawler() *ShanghaiCrawler {
	return &ShanghaiCrawler{}
}

func (b *ShanghaiCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.ShanghaiInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`[class="infoList js_infoList"]`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span.time").MustText(),
				Url:   browser.ShanghaiBaseUrl + strings.Replace(*element.MustElement("a").MustAttribute("href"), ".", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "上海",
		Msgs:   msg,
	}
	return data, nil
}
