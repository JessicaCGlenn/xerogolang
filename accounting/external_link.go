package accounting

type ExternalLink struct {

	// See External link types
	LinkType string `json:"LinkType,omitempty"`

	// URL for service e.g. http://twitter.com/jessicacglenn
	URL string `json:"Url,omitempty"`
}
