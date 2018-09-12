package main

import (
	"fmt"
	"time"
)

func main() {
	var c = make(chan Tag)

	appids := GetAllSteamApps()

	go DatabaseProcess(c)

	for index, appid := range appids {
		if index % 50 == 0 {
			time.Sleep(time.Millisecond * 1000)
		}
		go getTagsForAppProcess(c, appid)
	}

	var input string
	fmt.Scanln(&input)
}

func getTagsForAppProcess(c chan Tag, appid int) {
	tags := getTags(appid)

	for _, tag := range tags {
		c <- Tag{ appid, tag }
	}
}

func getTags(appid int) []string {
	html := GetSteamAppHTML(appid)
	return ExtractTags(html)
}