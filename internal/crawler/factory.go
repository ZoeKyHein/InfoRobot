package crawler

import (
	"fmt"
)

func GetCrawler(region string) (Crawler, error) {
	switch region {
	case "北京":
		return NewBeijingCrawler(), nil
	case "上海":
		return NewShanghaiCrawler(), nil
	case "天津":
		return NewTianjinCrawler(), nil
	case "大连":
		return NewDalianCrawler(), nil
	case "重庆":
		return NewChongqingCrawler(), nil
	default:
		return nil, fmt.Errorf("unsupported region: %s", region)
	}
}
