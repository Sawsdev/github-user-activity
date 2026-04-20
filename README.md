# Github user public activity
---

Github user activity project created for [roadmap.sh](https://roadmap.sh/) project [task tracker](https://roadmap.sh/projects/github-user-activity) the details of the project can be found there

---
## How to run the project
Clone the project and move into the project folder
~~~
https://github.com/Sawsdev/github-user-activity.git
cd github-user-activity
~~~

Build the binary an run the program
~~~
go build -o github-user-activity
./github-user-activity <username>
~~~

the program actually does a single execution so, you'll need to add a username to ensure it returns the activity of the given user

you can filter the type of the events using the --eventtype=<EVENTTYPE> or -e <EVENTTYPE>

### List of event types:

|Event type filter|description|
|-----------------|-----------|
|starred|Shows only events of starring a repo|
|create|Shows all type of creation events: branches, repos, etc.|
|delete|Shows all type of deletion events|
|fork|Show forked repos events|
|wiki|Show all related wiki action events|
|newissue|Show all issue creation events|
|commentissue|Show all comment issue events|
|push|Show a resume of all of the pushed commits into a repo events|
|release|Show all release events|
|pullrequest|Show all pull request events|


