package renderers_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"mkuznets.com/go/markdown/html"
	"mkuznets.com/go/markdown/renderers"
)

func testRenderer(t *testing.T, r renderers.Renderer) {
	ctx := context.Background()
	rb := bytes.NewBufferString(testMarkdown)
	wb := bytes.NewBuffer(nil)

	if err := r.Render(ctx, rb, wb); err != nil {
		t.Fatalf("render error: %v", err)
	}

	outBuf := bytes.NewBufferString("")
	if err := html.Normalise(wb, outBuf); err != nil {
		t.Fatalf("normaliser error: %v", err)
	}
	output := outBuf.String()

	if output != testHTML {
		fmt.Println(output)
		t.Fatal("html output does not match")
	}
}
