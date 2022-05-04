package models

type Link struct {
	Url    string `json:"url"`
	Parent string `json:"parent"`
	Links  []Link `json:"links"`
}

type Links struct {
	EntryPoint string `json:"entryPoint"`
	LinksFound []Link `json:"linksFound"`
}
