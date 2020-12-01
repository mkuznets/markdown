package renderers

import (
	"context"
	"io"
)

type Renderer interface {
	ID() string
	Render(ctx context.Context, r io.Reader, w io.Writer) error
}
