package renderers_test

import (
	"testing"

	"mkuznets.com/go/markdown/renderers"
)

func TestCMark_Render(t *testing.T) {
	r := renderers.NewCMark()
	testRenderer(t, r)
}
