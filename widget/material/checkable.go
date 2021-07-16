// SPDX-License-Identifier: Unlicense OR MIT

package material

import (
	"image"
	"image/color"

	"github.com/audrenbdb/gio/f32"
	"github.com/audrenbdb/gio/internal/f32color"
	"github.com/audrenbdb/gio/io/pointer"
	"github.com/audrenbdb/gio/layout"
	"github.com/audrenbdb/gio/op/clip"
	"github.com/audrenbdb/gio/op/paint"
	"github.com/audrenbdb/gio/text"
	"github.com/audrenbdb/gio/unit"
	"github.com/audrenbdb/gio/widget"
)

type checkable struct {
	Label              string
	Color              color.NRGBA
	Font               text.Font
	TextSize           unit.Value
	IconColor          color.NRGBA
	Size               unit.Value
	shaper             text.Shaper
	checkedStateIcon   *widget.Icon
	uncheckedStateIcon *widget.Icon
}

func (c *checkable) layout(gtx layout.Context, checked, hovered bool) layout.Dimensions {
	var icon *widget.Icon
	if checked {
		icon = c.checkedStateIcon
	} else {
		icon = c.uncheckedStateIcon
	}

	dims := layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{Alignment: layout.Center}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					size := gtx.Px(c.Size) * 4 / 3
					dims := layout.Dimensions{
						Size: image.Point{X: size, Y: size},
					}
					if !hovered {
						return dims
					}

					background := f32color.MulAlpha(c.IconColor, 70)

					radius := float32(size) / 2
					paint.FillShape(gtx.Ops, background,
						clip.Circle{
							Center: f32.Point{X: radius, Y: radius},
							Radius: radius,
						}.Op(gtx.Ops))

					return dims
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						size := gtx.Px(c.Size)
						icon.Color = c.IconColor
						if gtx.Queue == nil {
							icon.Color = f32color.Disabled(icon.Color)
						}
						gtx.Constraints.Min = image.Point{X: size}
						icon.Layout(gtx)
						return layout.Dimensions{
							Size: image.Point{X: size, Y: size},
						}
					})
				}),
			)
		}),

		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(2)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				paint.ColorOp{Color: c.Color}.Add(gtx.Ops)
				return widget.Label{}.Layout(gtx, c.shaper, c.Font, c.TextSize, c.Label)
			})
		}),
	)
	pointer.Rect(image.Rectangle{Max: dims.Size}).Add(gtx.Ops)
	return dims
}
