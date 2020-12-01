package renderers

import (
	"context"
	"io"
	"os/exec"
)

type CMark struct {
}

func NewCMark() *CMark {
	return &CMark{}
}

func (cm *CMark) ID() string {
	return "cmark"
}

func (cm *CMark) Render(ctx context.Context, r io.Reader, w io.Writer) error {
	c := exec.Command("cmark-gfm",
		"--to", "html",
		"-e", "table",
		"-e", "strikethrough",
		"-e", "autolink",
		"-e", "tagfilter",
		"-e", "tasklist",
		"--table-prefer-style-attributes",
	)
	c.Stdout = w
	c.Stdin = r
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
