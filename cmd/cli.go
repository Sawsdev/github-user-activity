package cmd

import (
	"fmt"
	"github.com/sawsdev/github-user-activity/internal/eventList"
	"github.com/spf13/cobra"
)


var RootCmd = &cobra.Command{
	Use: "github-activity <username>",
	Short: "A github activity fetcher api",
	Long: `github-activity is a CLI tool for fetching github user activities from github api.

	You can provide the username as argument via stdin.`,
	Args: cobra.ArbitraryArgs,
	Run: fetchGithubUserActivity,

}

var filterByEventType string

func init(){
	RootCmd.Flags().StringVarP(&filterByEventType, "eventtype", "e", "all", "event type to filter by: all, starred, create, delete, fork, wiki, newIssue, commentIssue, push, release, pullRequest")
}


func fetchGithubUserActivity( cmd *cobra.Command, args [] string){
	if !eventList.IsValidEventType(filterByEventType) {
		fmt.Println("The event type is not valid")
		return
	}
	events, err :=eventList.GetUserEvents(args[0])
	if err != ""{
		fmt.Println("has occurred an error" + err)
	}
	repoEvents, repoList := eventList.GroupEventsByRepo(&events)


	eventList.CreateActivityLog(repoEvents, repoList, eventList.GetEventType(filterByEventType))

}
