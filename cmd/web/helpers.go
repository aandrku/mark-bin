package main

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func render(w io.Writer, c templ.Component) error {
	return c.Render(context.Background(), w)
}
