package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var SearchAnime = &cobra.Command{
	Use:   "find",
	Short: "Search for a list of animes that are releated to the user-specified keyword.",
	Long:  "anime-cli find <ANIME_NAME>",
	Run: func(cmd *cobra.Command, args []string) {
		var queryString string

		// Get the query string from command line arguments.
		if len(args) == 1 {
			queryString = args[0]
		} else {
			for i, v := range args {
				if i == len(args)-1 {
					queryString += args[i]
				} else {
					queryString += v + "-"
				}
			}
		}

		// Get the list of animes
		DisplayList := searchAnime(queryString)
		if len(DisplayList) == 0 {
			fmt.Println("Cannot find anything!")
			os.Exit(1)
		}
		AnimeList := []Anime{}
		count := 0
		for _, v := range DisplayList {
			fmt.Println("[" + strconv.Itoa(count+1) + "] " + v.Title)
			AnimeList = append(AnimeList, v)
			count++
		}

		// TODO: Get input from stdin
		var input string
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter number: ")
		scanner.Scan()
		input = scanner.Text()
		index, err := strconv.Atoi(input)
		for err != nil {
			fmt.Println("Not a number!")
			fmt.Print("\nEnter number: ")
			scanner.Scan()
			input = scanner.Text()
			index, err = strconv.Atoi(input)
		}
		selected := AnimeList[index-1]
		s := spinner.New(spinner.CharSets[33], 50*time.Millisecond)
		s.Prefix = "🔎 Searching for the anime: \n"
		s.FinalMSG = color.GreenString("✔️Found!\n")

		go catchInterrupt(s)

		s.Start()
		s.Stop()
		number := getNumEps(selected)
		fmt.Printf("Choose episode [1-%d]: ", number)
		scanner.Scan()
		input = scanner.Text()
		current, er := strconv.Atoi(input)
		for er != nil {
			fmt.Println("Not a number!")
			fmt.Printf("Choose episode [1-%d]: ", number)
			scanner.Scan()
			input = scanner.Text()
			current, er = strconv.Atoi(input)
		}
		fmt.Printf("Getting data for episode %d\n", number)

		watchAnime(selected, input)

		for input != "q" {
			fmt.Printf("\nCurrently playing %s episode %d/%d\n", selected.Title, current, number)
			fmt.Print("[n] next episode\n")
			fmt.Print("[p] previous episode\n")
			fmt.Print("[q] exit\n")
			fmt.Print("Enter choice: ")
			scanner.Scan()
			input = scanner.Text()
			if input == "n" {
				current += 1
				watchAnime(selected, strconv.Itoa(current))
			} else if input == "p" {
				current -= 1
				watchAnime(selected, strconv.Itoa(current))
			} else if input == "q" {
				os.Exit(1)
			}
		}

	},
}

// lauch all commands
func AddCommands() {
	RootCmd.AddCommand(SearchAnime)
}
