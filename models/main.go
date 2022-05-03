package models

import "net/url"

type Link struct {
	Url      url.URL `json:"url"`
	Appeared int     `json:"appeared"`
	Parent   url.URL `json:"parent"`
}
