package layout

import (
	"gioui.org/layout"
)

type BtnGrp struct {
	Btns   []Btn
	Axis   layout.Axis
	Weight float32
}

func (btngrp *BtnGrp) Layout(gtx layout.Context) D {
	var ch []layout.FlexChild
	sp := &Spacer{
		Axis:   btngrp.Axis,
		Weight: btngrp.Weight,
	}
	for _, btn := range btngrp.Btns {

		ch = append(ch, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return btn.Layout(gtx)
			// return material.Button(th, c, values.GetText(t)).Layout(gtx)
		}), layout.Rigid(sp.Layout))
	}
	return layout.Flex{Axis: btngrp.Axis}.Layout(gtx, ch...)
}
