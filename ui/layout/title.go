package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/widget/material"
)

const(
	H1Title = iota
	H2Title
	H3Title
	H4Title
	H5Title
	H6Title
)

type Title struct {
	Type  int
	Text string
}

func (title *Title) Layout(gtx C) D {
	var res material.LabelStyle
	text := values.GetText(title.Text)
	switch title.Type {
	case H1Title:
		res = material.H1(Theme,text)
	case H2Title:
		res = material.H2(Theme,text)
	case H3Title:
		res = material.H3(Theme,text)
	case H4Title:
		res = material.H4(Theme,text)
	case H5Title:
		res = material.H5(Theme,text)
	case H6Title:
		res = material.H6(Theme,text)
	}
	return res.Layout(gtx)
	

}
