package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
)

type AllSteamApps struct {
  Applist Applist
}

type Applist struct {
  Apps []App
}

type App struct {
  Appid int
  Name  string
}

// GetAllSteamApps returns a integer list of the appids of all steam apps
func GetAllSteamApps() []int {
  apiKey := os.Getenv("STEAM_API_KEY")

  url := fmt.Sprintf("http://api.steampowered.com/ISteamApps/GetAppList/v002/?key=%s&format=json", apiKey)

  resp := getRequest(url)

  return parseAllAppsJson(resp)
}

func parseAllAppsJson(jsonString string) []int {
  var allSteamApps AllSteamApps

  json.Unmarshal([]byte(jsonString), &allSteamApps)

  apps := allSteamApps.Applist.Apps

  appids := make([]int, len(apps))

  for i, app := range apps {
    appids[i] = app.Appid
  }

  return appids
}

// GetSteamAppHTML returns the html for the steam store page of the given appid
func GetSteamAppHTML(appid int) string {
  return getRequest(fmt.Sprintf("https://store.steampowered.com/app/%d", appid))
}

func getRequest(url string) string {
  resp, err := http.Get(url)

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)

  return string(body[:])
}
