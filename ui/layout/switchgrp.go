package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

type SwitchGrp struct {
	Title    string
	Switches []Switch
	Axis     layout.Axis
	Spacer   int
}

func (switchGrp *SwitchGrp) Layout(gtx C) D {
	var childs []layout.FlexChild
	t := &Label{Text: switchGrp.Title}
	childs = append(childs, layout.Rigid(t.Layout))
	for _, sw := range switchGrp.Switches {
		childs = append(childs, layout.Rigid(layout.Spacer{Width: unit.Dp(switchGrp.Spacer)}.Layout),layout.Rigid(sw.Layout))
	}
	return layout.Flex{Axis: switchGrp.Axis}.Layout(gtx, childs...)
}
