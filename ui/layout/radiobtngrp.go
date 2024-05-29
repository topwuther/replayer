package layout

import (
	"strconv"

	"github.com/topwuther/replayer/ui/values"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type RadioBtn struct {
	Label string
}

type RadioBtnGrp struct {
	Enum      *widget.Enum
	Title     string
	RadioBtns []RadioBtn
	Axis      layout.Axis
	Weight    float32
}

func (radioBtnGrp *RadioBtnGrp) Layout(gtx C) D {
	var ch []layout.FlexChild
	sp := &Spacer{
		Axis:   radioBtnGrp.Axis,
		Weight: radioBtnGrp.Weight,
	}
	if radioBtnGrp.Title != "" {
		title := Label{
			Text: radioBtnGrp.Title,
			AppendText: ": ",
		}
		ch = append(ch, layout.Rigid(func(gtx C) D {
			return title.Layout(gtx)
		}))
	}

	for index, radioBtn := range radioBtnGrp.RadioBtns {
		ch = append(ch,
			layout.Rigid(material.RadioButton(Theme, radioBtnGrp.Enum, strconv.Itoa(index), values.GetText(radioBtn.Label)).Layout),
			layout.Rigid(sp.Layout),
		)
	}
	return layout.Flex{Axis: radioBtnGrp.Axis}.Layout(gtx, ch...)
}
