package goose

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/fatih/set.v0"
)

type Article struct {
	Title           string             `json:"title,omitempty"`
	CleanedText     string             `json:"content,omitempty"`
	MetaDescription string             `json:"description,omitempty"`
	MetaLang        string             `json:"lang,omitempty"`
	MetaFavicon     string             `json:"favicon,omitempty"`
	MetaKeywords    string             `json:"keywords,omitempty"`
	CanonicalLink   string             `json:"canonicalurl,omitempty"`
	Domain          string             `json:"domain,omitempty"`
	TopNode         *goquery.Selection `json:"-"`
	TopImage        string             `json:"image,omitempty"`
	Tags            *set.Set           `json:"tags,omitempty"`
	Movies          *set.Set           `json:"movies,omitempty"`
	FinalUrl        string             `json:"url,omitempty"`
	LinkHash        string             `json:"linkhash,omitempty"`
	RawHtml         string             `json:"rawhtml,omitempty"`
	Doc             *goquery.Document  `json:"-"`
	//raw_doc
	PublishDate    string            `json:"publishdate,omitempty"`
	AdditionalData map[string]string `json:"additionaldata,omitempty"`
	Delta          int64             `json:"delta,omitempty"`
}

//Simple ToString: it shows just the title
//TODO: add more fields and pretty print
func (article *Article) ToString() string {
	out := article.Title
	return out
}
