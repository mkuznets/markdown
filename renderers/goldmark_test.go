package renderers_test

import (
	"testing"

	"mkuznets.com/go/markdown/renderers"
)

func TestGoldMark_Render(t *testing.T) {
	r := &renderers.GoldMark{}
	testRenderer(t, r)
}
