package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	r, e := http.Get("https://api.github.com/users/prateesh94/events")
	var ev []Event
	var msg []string
	if e != nil {
		fmt.Println(e)
	}
	rdat, er := io.ReadAll(r.Body)
	if er != nil {
		fmt.Println(er)
	} else {
		er := json.Unmarshal(rdat, &ev)
		for i := range ev {
			switch ev[i].Type {
			case "CreateEvent":
				if ev[i].Payload.Ref_type == "repository" {
					msg = append(msg, "Created a "+ev[i].Payload.Ref_type+" named "+ev[i].Repo.Name+"\n")

				} else {
					msg = append(msg, "Created a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")

				}
				break
			case "CommitCommentEvent":
				if ev[i].Payload.Comment.Reason == "valid" {
					msg = append(msg, "Created a valid commit "+ev[i].Payload.Comment.Payload+" with signature "+ev[i].Payload.Comment.Signature+" in "+ev[i].Repo.Name+"\n")

				}
				break
			case "DeleteEvent":
				msg = append(msg, "Deleted a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")

				break
			case "GollumEvent":
				for pages := range ev[i].Payload.Pages {
					if ev[i].Payload.Pages[pages].Action == "created" {
						msg = append(msg, "Created page "+ev[i].Payload.Pages[pages].Page_name+" in "+ev[i].Repo.Name+"\n")

					} else {
						msg = append(msg, "Edited page "+ev[i].Payload.Pages[pages].Page_name+" last commit -"+ev[i].Payload.Pages[pages].Sha+" in "+ev[i].Repo.Name+"\n")

					}
				}

				break
			case "PushEvent":
				msg = append(msg, "Pushed "+strconv.Itoa(ev[i].Payload.Size)+" commits in "+ev[i].Repo.Name+"\n")

				break
			case "WatchEvent":
				msg = append(msg, "Starred "+ev[i].Repo.Name+"\n")

				break
			case "IssuesEvent":
				if ev[i].Payload.Action == "opened" {
					msg = append(msg, "Opened a new issue in "+ev[i].Repo.Name+"\n")

				} else if ev[i].Payload.Action == "edited" {
					msg = append(msg, "Edited an issue in "+ev[i].Repo.Name+"\n")

				} else if ev[i].Payload.Action == "closed" {
					msg = append(msg, "Closed an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "reopened" {
					msg = append(msg, "Reopend an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "assigned" {
					msg = append(msg, "Assigned an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unassigned" {
					msg = append(msg, "Unassigned an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "labeled" {
					msg = append(msg, "Labeled an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unlabeled" {
					msg = append(msg, "Unlabeled an issue in "+ev[i].Repo.Name+"\n")
				}
				break
			case "MemberEvent":
				if ev[i].Payload.Action == "added" {
					msg = append(msg, "Added a member to "+ev[i].Repo.Name+"\n")
				}
				break
			case "PublicEvent":
				msg = append(msg, "Made repo "+ev[i].Repo.Name+" public\n")
				break
			case "PullRequestEvent":
				if ev[i].Payload.Action == "opened" {
					msg = append(msg, "Opened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")

				} else if ev[i].Payload.Action == "edited" {
					msg = append(msg, "Edited pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")

				} else if ev[i].Payload.Action == "closed" {
					msg = append(msg, "Closed pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "reopened" {
					msg = append(msg, "Reopened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "assigned" {
					msg = append(msg, "Assigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unassigned" {
					msg = append(msg, "Unassigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "review_requested" {
					msg = append(msg, "Review requested on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "review_request_removed" {
					msg = append(msg, "Review request removed on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "labeled" {
					msg = append(msg, "Labeled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unlabeled" {
					msg = append(msg, "Unlabled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "synchronize" {
					msg = append(msg, "Synchronize on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				}
				break
			case "PullRequestReviewEvent":
				msg = append(msg, "Created pull request to review in "+ev[i].Repo.Name+"\n")
				break
			case "PullRequestReviewCommentEvent":
				if ev[i].Payload.Action == "created" {
					msg = append(msg, "Created comment on a pull request in "+ev[i].Repo.Name+"\n")
				} else {
					msg = append(msg, "Edited comment on a pull request in "+ev[i].Repo.Name+"\n")
				}
				break
			case "PullRequestReviewThreadEvent":
				if ev[i].Payload.Action == "resolved" {
					msg = append(msg, "Comment thread marked resolved in "+ev[i].Repo.Name+"\n")
				} else {
					msg = append(msg, "Comment thread marked unresolved in "+ev[i].Repo.Name+"\n")
				}
				break
			default:
				break

			}
		}

		if er != nil {
			fmt.Println(er)
		}
		for l := range msg {
			fmt.Println(msg[l])
		}
	}
}
