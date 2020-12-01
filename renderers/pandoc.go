package renderers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Pandoc struct {
	version string
}

func NewPandoc() (*Pandoc, error) {

	c := exec.Command("pandoc", "--version")
	out, err := c.Output()
	if err != nil {
		return nil, err
	}

	lines := bytes.SplitN(out, []byte("\n"), 2)
	versionLine := strings.Replace(string(lines[0]), " ", "-", 1)

	return &Pandoc{version: versionLine}, nil
}

func (p *Pandoc) ID() string {
	return fmt.Sprintf("%s-gfm", p.version)
}

func (p *Pandoc) Render(ctx context.Context, r io.Reader, w io.Writer) error {
	c := exec.Command("pandoc", "--from=gfm-gfm_auto_identifiers", "--to=html")
	c.Stdout = w
	c.Stdin = r
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}
