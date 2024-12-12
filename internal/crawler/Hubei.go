package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type HubeiCrawler struct{}

func NewHubeiCrawler() *HubeiCrawler {
	return &HubeiCrawler{}
}

func (b *HubeiCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.HubeiInfoUrl)
		p.MustWaitStable()
		table := p.MustElements(`[class="tab-box__list newsList tabcontent"]`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement(`span[class="newsTime fl"]`).MustText(),
				Url:   browser.HubeiBaseUrl + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "Hubei",
		Msgs:   msg,
	}
	return data, nil
}