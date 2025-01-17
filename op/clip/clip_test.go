// SPDX-License-Identifier: Unlicense OR MIT

package clip

import (
	"testing"

	"github.com/audrenbdb/gio/f32"
	"github.com/audrenbdb/gio/op"
)

func TestOpenPathOutlinePanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("Outline of an open path didn't panic")
		}
	}()
	var p Path
	p.Begin(new(op.Ops))
	p.Line(f32.Pt(10, 10))
	Outline{Path: p.End()}.Op()
}
