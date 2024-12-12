package main

import (
	"InfoRobot/internal"
	"testing"
)

func TestSendMultiRegions(t *testing.T) {
	regions := []string{
		//"北京",
		//"上海",
		//"天津",
		//"大连",
		//"重庆",
		//"青岛",
		//"宁波",
		//"厦门",
		//"深圳",
		//"内蒙古",
		//"安徽",
		//"福建",
		//"甘肃",
		"广东",
	}
	internal.SpiderAndSend(regions)
}
