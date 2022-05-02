package models

import "net/url"

type link struct {
	url      url.URL
	appeared int
	parent   url.URL
}
