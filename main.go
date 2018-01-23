package main

import (
	"fmt"
	"os"
	"strings"

	"net/http"

	"github.com/igorpo/financier/bloomberg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
}

func main() {
	args := os.Args[1:]
	ticker := args[0]
	log.WithFields(log.Fields{
		"ticker": ticker,
	}).Info("command line args")

	resp, err := http.Get(fmt.Sprintf("https://www.bloomberg.com/quote/%s:US", ticker))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to parse GET request")
	}

	fields := make([]string, 0)

	doc := html.NewTokenizer(resp.Body)
	for tokenType := doc.Next(); tokenType != html.ErrorToken; {
		token := doc.Token()
		switch tokenType {
		case html.StartTagToken:
			if token.DataAtom != atom.Span && token.DataAtom != atom.Div {
				tokenType = doc.Next()
				continue
			}

			// TODO: add workers and refactor later
			for _, attr := range token.Attr {
				if attr.Key == "class" {
					tagClassName := attr.Val
					if strings.HasPrefix(tagClassName, "companyId") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "exchange") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "priceText") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "changeAbsolute") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "changePercent") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "value") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					} else if strings.HasPrefix(tagClassName, "fieldValue") {
						innerTokenType := doc.Next()
						if innerTokenType != html.ErrorToken && innerTokenType == html.TextToken {
							token = doc.Token()
							// fmt.Println(tagClassName, ": ", token.String())
							fields = append(fields, token.String())
						}
					}
				}
			}
		}
		tokenType = doc.Next()
	}
	resp.Body.Close()

	stock := bloomberg.NewImportedStock(fields)

	log.WithFields(log.Fields{
		"fields": fields,
	}).Info("future json fields of bloomberg stock")
}
