package renderers

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

type GoldMark struct {
}

func NewGoldMark() *GoldMark {
	return &GoldMark{}
}

func (g *GoldMark) ID() string {
	return "goldmark-v1.2.1-gfm"
}

func (g *GoldMark) Render(_ context.Context, r io.Reader, w io.Writer) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	if err := md.Convert(src, w); err != nil {
		return err
	}

	return nil
}
