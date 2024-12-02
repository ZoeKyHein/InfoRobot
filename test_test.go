package main

import (
	"InfoRobot/internal"
	"testing"
)

func TestSendMultiRegions(t *testing.T) {
	regions := []string{
		"北京",
		"上海",
	}
	internal.SpiderAndSend(regions)
}
