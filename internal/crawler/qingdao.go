package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"strings"
)

type QingdaoCrawler struct{}

func NewQingdaoCrawler() *QingdaoCrawler {
	return &QingdaoCrawler{}
}

func (b *QingdaoCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.MustNavigate(browser.QingdaoInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`#tongzhi_con1`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.QingdaoInfoUrl + strings.Replace(*element.MustElement("a").MustAttribute("href"), "./", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "青岛",
		Msgs:   msg,
	}
	return data, nil
}
