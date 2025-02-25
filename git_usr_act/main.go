package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	arg := os.Args[1:]
	r, e := http.Get("https://api.github.com/users/" + arg[0] + "/events")
	var ev []Event
	var msg []string
	var cr_event, comm_event, del_event, gol_event, push_event, wtc_event, iss_event, mem_event, pub_event, pullrq_event, pullre_rev_event, pullrec_rev_com_event, pullrec_rev_thread_event []string
	if e != nil {
		fmt.Println(e)
	}
	rdat, er := io.ReadAll(r.Body)

	if er != nil || len(rdat) == 0 {
		fmt.Println("User Not Found...Enter valid username")
		return
	} else {
		er := json.Unmarshal(rdat, &ev)
		for i := range ev {
			switch ev[i].Type {
			case "CreateEvent":
				if ev[i].Payload.Ref_type == "repository" {
					msg = append(msg, "Created a "+ev[i].Payload.Ref_type+" named "+ev[i].Repo.Name+"\n")
					cr_event = append(cr_event, "Created a "+ev[i].Payload.Ref_type+" named "+ev[i].Repo.Name+"\n")
				} else {
					msg = append(msg, "Created a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")
					cr_event = append(cr_event, "Created a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")
				}
				break
			case "CommitCommentEvent":
				if ev[i].Payload.Comment.Reason == "valid" {
					msg = append(msg, "Created a valid commit "+ev[i].Payload.Comment.Payload+" with signature "+ev[i].Payload.Comment.Signature+" in "+ev[i].Repo.Name+"\n")
					comm_event = append(comm_event, "Created a valid commit "+ev[i].Payload.Comment.Payload+" with signature "+ev[i].Payload.Comment.Signature+" in "+ev[i].Repo.Name+"\n")
				}
				break
			case "DeleteEvent":
				msg = append(msg, "Deleted a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")
				del_event = append(del_event, "Deleted a "+ev[i].Payload.Ref_type+" "+ev[i].Payload.Ref+" in "+ev[i].Repo.Name+"\n")
				break
			case "GollumEvent":
				for pages := range ev[i].Payload.Pages {
					if ev[i].Payload.Pages[pages].Action == "created" {
						msg = append(msg, "Created page "+ev[i].Payload.Pages[pages].Page_name+" in "+ev[i].Repo.Name+"\n")
						gol_event = append(gol_event, "Created page "+ev[i].Payload.Pages[pages].Page_name+" in "+ev[i].Repo.Name+"\n")
					} else {
						msg = append(msg, "Edited page "+ev[i].Payload.Pages[pages].Page_name+" last commit -"+ev[i].Payload.Pages[pages].Sha+" in "+ev[i].Repo.Name+"\n")
						gol_event = append(gol_event, "Edited page "+ev[i].Payload.Pages[pages].Page_name+" last commit -"+ev[i].Payload.Pages[pages].Sha+" in "+ev[i].Repo.Name+"\n")
					}
				}

				break
			case "PushEvent":
				msg = append(msg, "Pushed "+strconv.Itoa(ev[i].Payload.Size)+" commits in "+ev[i].Repo.Name+"\n")
				push_event = append(push_event, "Pushed "+strconv.Itoa(ev[i].Payload.Size)+" commits in "+ev[i].Repo.Name+"\n")
				break
			case "WatchEvent":
				msg = append(msg, "Starred "+ev[i].Repo.Name+"\n")
				wtc_event = append(wtc_event, "Starred "+ev[i].Repo.Name+"\n")
				break
			case "IssuesEvent":
				if ev[i].Payload.Action == "opened" {
					msg = append(msg, "Opened a new issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Opened a new issue in "+ev[i].Repo.Name+"\n")

				} else if ev[i].Payload.Action == "edited" {
					msg = append(msg, "Edited an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Edited an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "closed" {
					msg = append(msg, "Closed an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Closed an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "reopened" {
					msg = append(msg, "Reopend an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Reopend an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "assigned" {
					msg = append(msg, "Assigned an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Assigned an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unassigned" {
					msg = append(msg, "Unassigned an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Unassigned an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "labeled" {
					msg = append(msg, "Labeled an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Labeled an issue in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unlabeled" {
					msg = append(msg, "Unlabeled an issue in "+ev[i].Repo.Name+"\n")
					iss_event = append(iss_event, "Unlabeled an issue in "+ev[i].Repo.Name+"\n")
				}
				break
			case "MemberEvent":
				if ev[i].Payload.Action == "added" {
					msg = append(msg, "Added a member to "+ev[i].Repo.Name+"\n")
					mem_event = append(mem_event, "Added a member to "+ev[i].Repo.Name+"\n")
				}
				break
			case "PublicEvent":
				msg = append(msg, "Made repo "+ev[i].Repo.Name+" public\n")
				pub_event = append(pub_event, "Made repo "+ev[i].Repo.Name+" public\n")
				break
			case "PullRequestEvent":
				if ev[i].Payload.Action == "opened" {
					msg = append(msg, "Opened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Opened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "edited" {
					msg = append(msg, "Edited pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Edited pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "closed" {
					msg = append(msg, "Closed pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Closed pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "reopened" {
					msg = append(msg, "Reopened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Reopened pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "assigned" {
					msg = append(msg, "Assigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Assigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unassigned" {
					msg = append(msg, "Unassigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Unassigned pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "review_requested" {
					msg = append(msg, "Review requested on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Review requested on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "review_request_removed" {
					msg = append(msg, "Review request removed on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Review request removed on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "labeled" {
					msg = append(msg, "Labeled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Labeled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "unlabeled" {
					msg = append(msg, "Unlabled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Unlabled pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				} else if ev[i].Payload.Action == "synchronize" {
					msg = append(msg, "Synchronize on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
					pullrq_event = append(pullrq_event, "Synchronize on pull request no. "+strconv.Itoa(ev[i].Payload.Number)+" in "+ev[i].Repo.Name+"\n")
				}
				break
			case "PullRequestReviewEvent":
				msg = append(msg, "Created pull request to review in "+ev[i].Repo.Name+"\n")
				pullre_rev_event = append(pullre_rev_event, "Created pull request to review in "+ev[i].Repo.Name+"\n")
				break
			case "PullRequestReviewCommentEvent":
				if ev[i].Payload.Action == "created" {
					msg = append(msg, "Created comment on a pull request in "+ev[i].Repo.Name+"\n")
					pullrec_rev_com_event = append(pullrec_rev_com_event, "Created comment on a pull request in "+ev[i].Repo.Name+"\n")
				} else {
					msg = append(msg, "Edited comment on a pull request in "+ev[i].Repo.Name+"\n")
					pullrec_rev_com_event = append(pullrec_rev_com_event, "Edited comment on a pull request in "+ev[i].Repo.Name+"\n")
				}
				break
			case "PullRequestReviewThreadEvent":
				if ev[i].Payload.Action == "resolved" {
					msg = append(msg, "Comment thread marked resolved in "+ev[i].Repo.Name+"\n")
					pullrec_rev_thread_event = append(pullrec_rev_thread_event, "Comment thread marked resolved in "+ev[i].Repo.Name+"\n")
				} else {
					msg = append(msg, "Comment thread marked unresolved in "+ev[i].Repo.Name+"\n")
					pullrec_rev_thread_event = append(pullrec_rev_thread_event, "Comment thread marked unresolved in "+ev[i].Repo.Name+"\n")
				}
				break
			default:
				break

			}
		}

		if er != nil {
			fmt.Println(er)
		}

	}
	for {
		var tok string
		fmt.Println("Choose event type to filter data")
		fmt.Println("Event types are:-")
		fmt.Printf("1. ViewAll\n2. CreateEvent\n3. DeleteEvent\n4. GollumEvent\n5. PushEvent\n6. WatchEvent\n7. IssuesEvent\n8. MemberEvent\n9. PublicEvent\n10. PullRequestEvent\n11.CommitEvent\n12.Exit -to exit the program\n")
		fmt.Scanf("%s\n", &tok)
		tok = strings.ToLower(tok)
		switch tok {
		case "viewall":
			for r := range msg {
				fmt.Printf(msg[r])
			}
			break
		case "createevent":
			if len(cr_event) > 0 {
				for r := range cr_event {
					fmt.Printf(cr_event[r])
				}
			} else {
				fmt.Println("No create events found")
			}
			break
		case "deleteevent":
			if len(del_event) > 0 {
				for r := range del_event {
					fmt.Printf(del_event[r])
				}
			} else {
				fmt.Println("No delete events found")
			}
			break
		case "commitevent":
			if len(comm_event) > 0 {
				for r := range comm_event {
					fmt.Printf(comm_event[r])
				}
			} else {
				fmt.Println("No commit events found")
			}
			break
		case "gollumevent":
			if len(gol_event) > 0 {
				for r := range gol_event {
					fmt.Printf(gol_event[r])
				}
			} else {
				fmt.Println("No gollum events found")
			}
			break
		case "pushevent":
			if len(push_event) > 0 {
				for r := range push_event {
					fmt.Printf(push_event[r])
				}
			} else {
				fmt.Println("No push events found")
			}
			break
		case "watchevent":
			if len(wtc_event) > 0 {
				for r := range wtc_event {
					fmt.Printf(wtc_event[r])
				}
			} else {
				fmt.Println("No commit comment events found")
			}
			break
		case "issuesevent":
			if len(iss_event) > 0 {
				for r := range iss_event {
					fmt.Printf(iss_event[r])
				}
			} else {
				fmt.Println("No issues events found")
			}
			break
		case "memberevent":
			if len(mem_event) > 0 {
				for r := range mem_event {
					fmt.Printf(mem_event[r])
				}
			} else {
				fmt.Println("No member events found")
			}
			break
		case "publicevent":
			if len(pub_event) > 0 {
				for r := range pub_event {
					fmt.Printf(pub_event[r])
				}
			} else {
				fmt.Println("No commit comment events found")
			}
			break
		case "pullrequestevent":
			if len(pullrq_event) > 0 {
				for r := range comm_event {
					fmt.Printf(comm_event[r])
				}
			} else if len(pullre_rev_event) > 0 {
				for r := range pullre_rev_event {
					fmt.Printf(pullre_rev_event[r])
				}
			} else if len(pullrec_rev_com_event) > 0 {
				for r := range pullrec_rev_com_event {
					fmt.Printf(pullrec_rev_com_event[r])
				}

			} else if len(pullrec_rev_thread_event) > 0 {
				for r := range pullrec_rev_thread_event {
					fmt.Printf(pullrec_rev_thread_event[r])
				}

			} else {
				fmt.Println("No pull events found")
			}
			break
		case "exit":
			os.Exit(0)
			break
		default:
			fmt.Println("Enter valid Event name")
			break
		}

	}
}
