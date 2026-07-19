package parser

import (
	
	"golang.org/x/net/html"
	"strings"
)

func ParseUrl(rawHtml string) []string {
	tokenizer := html.NewTokenizer(strings.NewReader(rawHtml))
	var links []string

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return links

		case html.StartTagToken, html.SelfClosingTagToken:

			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
						break
					}
				}

			}

		}

	}
}
