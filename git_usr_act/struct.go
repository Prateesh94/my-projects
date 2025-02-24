package main

type Event struct {
	Id         string  `json:"id"`
	Type       string  `json:"type"`
	Actor      Actor   `json:"actor"`
	Repo       Repo    `json:"repo"`
	Payload    Payload `json:"payload"`
	Public     bool    `json:"public"`
	Created_at string  `json:"created_at"`
	Org        Org     `json:"org"`
}

type Actor struct {
	Id            int    `json:"id"`
	Login         string `json:"login"`
	Display_login string `json:"display_login"`
	Gravatar_id   string `json:"gravatar_id"`
	Url           string `json:"url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Org struct {
	Id          int    `json:"id"`
	Login       string `json:"login"`
	Gravatar_id string `json:"gravatar_id"`
	Url         string `json:"url"`
	Avatar_url  string `json:"avatar_url"`
}
type Comment struct {
	Verifed     bool   `json:"verified"`
	Reason      string `json:"reason"`
	Signature   string `json:"signature"`
	Payload     string `json:"payload"`
	Verified_at string `json:"verified_at"`
}
type Payload struct {
	Comment Comment `json:"comment"`
	//forkee
	//changes, changes[body][from]
	//issue
	//assignee
	//label
	//member
	//pull request
	//review
	//thread
	//release
	Action         string    `json:"action"`
	Ref            string    `json:"ref"`
	Ref_type       string    `json:"ref_type"`
	Master_branch  string    `json:"master_branch"`
	Description    string    `json:"description"`
	Pusher_type    string    `json:"pusher_type"`
	Pages          []Pages   `json:"pages"`
	Number         int       `json:"number"`
	Reason         string    `json:"reason"`
	Push_id        int       `json:"push_id"`
	Size           int       `json:"size"`
	Distinct_size  int       `json:"distinct_size"`
	Head           string    `json:"head"`
	Before         string    `json:"before"`
	Commits        []Commits `json:"commits"`
	Effective_date string    `json:"effective_date"`
}
type Pages struct {
	Page_name string `json:"page_name"`
	Title     string `json:"title"`
	Action    string `json:"action"`
	Sha       string `json:"sha"`
	Html_url  string `json:"html_url"`
}
type Commits struct {
	Sha      string `json:"sha"`
	Message  string `json:"message"`
	Author   Author `json:"author"`
	Url      string `json:"url"`
	Distinct bool   `json:"distinct"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
