package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type JilinCrawler struct{}

func NewJilinCrawler() *JilinCrawler {
	return &JilinCrawler{}
}

func (b *JilinCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.JilinInfoUrl)
		p.MustWaitStable()
		table := p.MustElements(`[class="common-list-items"]`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "Jilin",
		Msgs:   msg,
	}
	return data, nil
}