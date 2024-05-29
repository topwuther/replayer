package layout

import (
	"github.com/topwuther/replayer/ui/values"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type CheckBox struct {
	Label   string
	Checked *widget.Bool
}

type CheckBoxGrp struct {
	Title      string
	CheckBoxes []CheckBox
	Axis       layout.Axis
	Spacer     int
}

func (checkBoxGrp *CheckBoxGrp) Layout(gtx C) D {
	var childs []layout.FlexChild
	t := &Label{Text: checkBoxGrp.Title}
	childs = append(childs, layout.Rigid(t.Layout))
	for _, checkbox := range checkBoxGrp.CheckBoxes {
		childs = append(childs, layout.Rigid(material.CheckBox(Theme, checkbox.Checked, values.GetText(checkbox.Label)).Layout), layout.Rigid(layout.Spacer{Width: unit.Dp(checkBoxGrp.Spacer)}.Layout))
	}
	return layout.Flex{Axis: checkBoxGrp.Axis}.Layout(gtx, childs...)
}
