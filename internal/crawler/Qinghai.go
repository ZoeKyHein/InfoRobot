package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type QinghaiCrawler struct{}

func NewQinghaiCrawler() *QinghaiCrawler {
	return &QinghaiCrawler{}
}

func (b *QinghaiCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.QinghaiInfoUrl)
		p.MustWaitStable()
		table := p.MustElements(`[id="tzgg"]`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.QinghaiBaseUrl + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "Qinghai",
		Msgs:   msg,
	}
	return data, nil
}
