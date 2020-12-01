package html

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
)

func Normalise(r io.Reader, w io.Writer) error {
	colonRE := regexp.MustCompile(`:\s+`)

	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if len(n.Attr) > 0 {
			attrs := make([]html.Attribute, 0)
			for _, a := range n.Attr {
				if a.Key == "class" {
					continue
				}
				if a.Key == "style" {
					a.Val = strings.TrimSuffix(a.Val, ";")
					a.Val = colonRE.ReplaceAllString(a.Val, ":")
				}
				attrs = append(attrs, a)
			}
			n.Attr = attrs
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	processedHTML := bytes.NewBuffer(nil)
	if err := html.Render(processedHTML, doc); err != nil {
		return err
	}

	output := processedHTML.Bytes()
	output = bytes.TrimSpace(output)
	output = gohtml.FormatBytes(output)

	if _, err := w.Write(output); err != nil {
		return err
	}

	return nil
}
