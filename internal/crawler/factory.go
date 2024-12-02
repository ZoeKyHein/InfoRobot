package crawler

import "fmt"

func GetCrawler(region string) (Crawler, error) {
	switch region {
	case "北京":
		return NewBeijingCrawler(), nil
	case "上海":
		return NewShanghaiCrawler(), nil
	default:
		return nil, fmt.Errorf("unsupported region: %s", region)
	}
}
