package eventList

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"

	"github.com/sawsdev/github-user-activity/internal/event"
	"github.com/sawsdev/github-user-activity/internal/fetch"
)

type EventList struct {
	Events []event.Event
}

type RepoEventList struct {
	RepoName string
	Events []event.Event
}

const (
	STARRED = "WatchEvent"
	CREATE = "CreateEvent"
	DELETE = "DeleteEvent"
	FORK = "ForkEvent"
	WIKI = "GollumEvent"
	NEWISSUE = "IssuesEvent"
	COMMENTISSUE = "IssueCommentEvent"
	PUSH = "PushEvent"
	RELEASE = "ReleaseEvent"
	PULLREQUEST = "PullRequestEvent"

)

var EventTypes = map[string]string{
	STARRED : "WatchEvent",
	CREATE : "CreateEvent",
	DELETE : "DeleteEvent",
	FORK : "ForkEvent",
	WIKI : "GollumEvent",
	NEWISSUE : "IssuesEvent",
	COMMENTISSUE : "IssueCommentEvent",
	PUSH : "PushEvent",
	RELEASE : "ReleaseEvent",
	PULLREQUEST : "PullRequestEvent",
}


var githubUserBaseUrl = "https://api.github.com/users/"

func GetUserEvents(username string) (EventList, string) {
	if !validateUserName(username) {
		return EventList {
			Events: []event.Event{},
		}, "The username is not valid."
	}
	response := fetch.GetFromUrl(githubUserBaseUrl+username+"/events")
	responseData := []byte(response)
	var events = EventList {
		Events: []event.Event{},
	}
	err := json.Unmarshal(responseData, &events.Events)
	if err != nil {
		log.Fatal(err)
	}
	return events, ""
}

func validateUserName (username string) bool{
 match, err := regexp.MatchString("^[a-zA-Z0-9]([a-zA-Z0-9-]{0,36}[a-zA-Z0-9])?$", username)
 if err != nil {
	log.Fatal(err)
 }
 return match
}

func GroupEventsByRepo (eventList *EventList) (map[string]RepoEventList, map[string]string){

	var groupedEvents = make(map[string]RepoEventList)
	var repoList = make(map[string]string)
	for _, e := range eventList.Events{

			if groupedEvents[e.Repo.Name].RepoName == e.Repo.Name{
				 groupedEvents[e.Repo.Name] = RepoEventList{
					  e.Repo.Name,
					  append(groupedEvents[e.Repo.Name].Events, e),
				 }
			} else {
				 groupedEvents[e.Repo.Name] = RepoEventList{
					  e.Repo.Name,
					  []event.Event{e},
				 }
			}
			repoList[e.Repo.Name] = e.Repo.Name
		}
	return groupedEvents, repoList
}


func CreateActivityLog(repoEvents map[string]RepoEventList, repoList map[string]string, filter string){

	starred,create,delete,fork,wiki,newIssue,commentIssue,push,release,pullRequest := 0,0,0,0,0,0,0,0,0,0
	activityLogText := "Activity log: \n"
	userName := ""
	for _, repo := range repoList {
		if(filter != "all" && filter != repoEvents[repo].Events[0].Type){
			continue
		}
		events := repoEvents[repo].Events
		userName = repoEvents[repo].Events[0].Actor.Login
		for _, e := range events {
			switch e.Type {
			case STARRED:
				starred += 1
			case CREATE:
				create +=1
			case DELETE:
				delete +=1
			case FORK:
				fork++
			case WIKI:
				wiki++
			case NEWISSUE:
				newIssue++
			case COMMENTISSUE:
				commentIssue++
			case PUSH:
				push++
			case RELEASE:
				release++
			case PULLREQUEST:
				pullRequest++
				
			}
		}
		if starred > 0 {
			activityLogText += fmt.Sprintf("%s has starred %s \n", userName, repo)
		}
		if create > 0 {
			activityLogText += fmt.Sprintf("%s has created a new %s\n", userName, repo)
		}
		if delete > 0 {
			activityLogText += fmt.Sprintf("%s has deleted %s\n", userName, repo)
		}
		if fork > 0 {
			activityLogText += fmt.Sprintf("%s has forked the repo %s\n", userName, repo)
		}
		if wiki > 0 {
			activityLogText += fmt.Sprintf("%s has created a new wiki page on %s\n", userName, repo)
		}
		if newIssue > 0 {
			activityLogText += fmt.Sprintf("%s has created a new issue on %s\n", userName, repo)
		}
		if commentIssue > 0 {
			activityLogText += fmt.Sprintf("%s has commented on an issue on %s\n", userName, repo)
		}
		if push > 0 {
			activityLogText += fmt.Sprintf("%s has pushed %d commits to %s\n", userName, push, repo)
		}
		if release > 0 {
			activityLogText += fmt.Sprintf("%s has released a new version of %s\n", userName, repo)
		}

	}
	fmt.Println(activityLogText)
}

func IsValidEventType (eventType string) bool{
	match, err := regexp.MatchString("all|starred|create|delete|fork|wiki|newIssue|commentIssue|push|release|pullRequest", eventType)
	if err != nil {
		log.Fatal(err)
	}
	return match
}

func GetEventType(eventType string) string {
	switch eventType {
	case "starred":
		return STARRED
	case "create":
		return CREATE
	case "delete":
		return DELETE
	case "fork":
		return FORK
	case "wiki":
		return WIKI
	case "newIssue":
		return NEWISSUE
	case "commentIssue":
		return COMMENTISSUE
	case "push":
		return PUSH
	case "release":
		return RELEASE
	case "pullRequest":
		return PULLREQUEST
		
	}
	return "all"
}


