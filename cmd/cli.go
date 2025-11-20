package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sawsdev/github-user-activity/internal/fetch"
)


var RootCmd = &cobra.Command{
	Use: "github-activity <username>",
	Short: "A github activity fetcher api",
	Long: `github-activity is a CLI tool for fetching github user activities from github api.

	You can provide the username as argument via stdin.`,
	Args: cobra.ArbitraryArgs,
	Run: fetchGithubUserActivity,

}

var githubUserBaseUrl = "https://api.github.com/users/"




func fetchGithubUserActivity( cmd *cobra.Command, args [] string){
	fmt.Println(args)
	response := fetch.GetFromUrl(githubUserBaseUrl+"midudev/events")
	fmt.Println(response)
}
