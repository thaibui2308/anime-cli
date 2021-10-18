package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const baseUrl = "https://gogoanime.pe//search.html?keyword="
const watchUrl = "https://goload.one/videos/"

func searchAnime(query string) map[string]Anime {
	List := map[string]Anime{}

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !strings.HasPrefix(link, "/category/") {
			return
		}
		temp := Anime{}
		temp.Link = "https://gogoanime.pe" + link
		temp.Title = e.Attr("title")
		temp.Query = generateQuery(strings.Split(link, "/")[2])
		_, ok := List[temp.Title]
		if !ok && len(List) <= 20 {
			List[temp.Title] = temp
		}
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit(baseUrl + query)
	return List
}

// Get number of episode
func getNumEps(a Anime) int {
	counter := 0
	c := colly.NewCollector()
	visited := map[string]bool{}

	// Find and visit all links
	c.OnHTML("a[ep_end]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		_, ok := visited[link]
		if !ok {
			visited[link] = true
			counter, _ = strconv.Atoi(e.Attr("ep_end"))
		}
	})
	c.Visit(a.Link)
	return counter
}

func watchAnime(anime Anime, episode string) {
	PageUrl := watchUrl + anime.Query + "-episode-" + episode

	var url string

	c := colly.NewCollector()
	c.OnHTML("iframe[src]", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		if !strings.HasPrefix(link, "//goload.one/streaming.php") {
			return
		}
		url = link
	})
	c.Visit(PageUrl)
	url = "https:" + url
	anime.ChangeWatchUrl("https:" + watchUrl)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func generateQuery(str string) string {
	res := strings.ReplaceAll(str, "/anime/", "")
	return res
}
