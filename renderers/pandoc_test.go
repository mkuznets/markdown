package renderers_test

import (
	"testing"

	"mkuznets.com/go/markdown/renderers"
)

func TestPandoc_Render(t *testing.T) {
	r, err := renderers.NewPandoc()
	if err != nil {
		t.Fatalf("could not initialise pandoc: %v", err)
	}
	testRenderer(t, r)
}
