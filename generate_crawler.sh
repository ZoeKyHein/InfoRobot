#!/bin/bash

# 定义地区列表，从广西开始
regions=("Guangxi" "Xizang" "Ningxia" "Xinjiang" "Hebei" "Shanxi" "Liaoning" "Jilin" "Heilongjiang" "Jiangsu" "Zhejiang" "Anhui" "Fujian" "Jiangxi" "Shandong" "Henan" "Hubei" "Hunan" "Guangdong" "Hainan" "Sichuan" "Guizhou" "Yunnan" "Shaanxi" "Gansu" "Qinghai")

# 生成文件夹
mkdir -p crawlers

# 遍历地区列表并生成对应的 .go 文件
for region in "${regions[@]}"; do
    file="crawlers/${region}.go"
    touch "$file"
    cat <<EOL >"$file"
package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type ${region}Crawler struct{}

func New${region}Crawler() *${region}Crawler {
	return &${region}Crawler{}
}

func (b *${region}Crawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.${region}InfoUrl)
		p.MustWaitStable()
		table := p.MustElements(\`[class="common-list-items"]\`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.${region}BaseUrl + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "${region}",
		Msgs:   msg,
	}
	return data, nil
}
EOL
    echo "$file created."
done

echo "All crawler files generated in the 'crawlers' directory."