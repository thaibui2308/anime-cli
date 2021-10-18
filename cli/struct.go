package cli

type Anime struct {
	Link     string
	Title    string
	WatchUrl string
	Query    string
}

func (a *Anime) ChangeWatchUrl(url string) {
	a.WatchUrl = url
}
