package layout

import (
	"gioui.org/layout"
)

type Form struct {
	Editor []*Editor
	RadioBtnGrps []*RadioBtnGrp
	SwitchGrp []*Switch
	BtnGrp *BtnGrp
	Axis   layout.Axis
	Weight float32
}

func (form *Form) Layout(gtx C) D {
	var ch []layout.FlexChild
	sp := &Spacer{
		Axis:   form.Axis,
		Weight: form.Weight,
	}
	for _, ed := range form.Editor {
		ch = append(ch, layout.Rigid(func(gtx C) D {
			return ed.Layout(gtx)
		}), layout.Rigid(sp.Layout))
	}
	for _,radioBtnGrp := range form.RadioBtnGrps {
		ch = append(ch, layout.Rigid(func(gtx C) D {
			return radioBtnGrp.Layout(gtx)
		}),layout.Rigid(sp.Layout))
	}
	for _,sw := range form.SwitchGrp {
		ch = append(ch, layout.Rigid(func(gtx C) D {
			return sw.Layout(gtx)
		}),layout.Rigid(sp.Layout))
	}
	if form.BtnGrp != nil {
		form.BtnGrp.Axis = layout.Horizontal
		ch = append(ch, layout.Rigid(func(gtx C) D {
			return form.BtnGrp.Layout(gtx)
		}))
	}
	return layout.Flex{Axis: form.Axis}.Layout(gtx, ch...)
}
