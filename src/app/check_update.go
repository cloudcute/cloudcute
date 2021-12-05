package app

//type GitHubRelease struct {
//	URL  string `json:"html_url"`
//	Name string `json:"name"`
//	Tag  string `json:"tag_name"`
//}

func checkUpdate()  {
	// TODO 检查更新
	//client := request.NewClient()
	//res, err := client.Request("GET", "https://api.github.com/repos/cloudcute/cloudcute/releases", nil).GetResponse()
	//if err != nil {
	//	util.Log().Warning("更新检查失败, %s", err)
	//	return
	//}
	//
	//var list []GitHubRelease
	//if err := json.Unmarshal([]byte(res), &list); err != nil {
	//	util.Log().Warning("更新检查失败, %s", err)
	//	return
	//}
	//
	//if len(list) > 0 {
	//	present, err1 := version.NewVersion(conf.BackendVersion)
	//	latest, err2 := version.NewVersion(list[0].Tag)
	//	if err1 == nil && err2 == nil && latest.GreaterThan(present) {
	//		util.Log().Info("有新的版本 [%s] 可用，下载：%s", list[0].Name, list[0].URL)
	//	}
	//}
}
