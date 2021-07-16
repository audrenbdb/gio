// SPDX-License-Identifier: Unlicense OR MIT

package headless

import (
	"errors"
	"syscall/js"

	"github.com/audrenbdb/gio/gpu"
	"github.com/audrenbdb/gio/internal/gl"
)

type jsContext struct {
	ctx js.Value
}

func newGLContext() (context, error) {
	doc := js.Global().Get("document")
	cnv := doc.Call("createElement", "canvas")
	ctx := cnv.Call("getContext", "webgl2")
	if ctx.IsNull() {
		ctx = cnv.Call("getContext", "webgl")
	}
	if ctx.IsNull() {
		return nil, errors.New("headless: webgl is not supported")
	}
	c := &jsContext{
		ctx: ctx,
	}
	return c, nil
}

func (c *jsContext) API() gpu.API {
	return gpu.OpenGL{Context: gl.Context(c.ctx)}
}

func (c *jsContext) Release() {
}

func (c *jsContext) ReleaseCurrent() {
}

func (c *jsContext) MakeCurrent() error {
	return nil
}
