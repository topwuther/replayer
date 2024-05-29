package player

import (
	"io"
	"strconv"
)

func (p *Player) Progress() float32 {
	if p.Downloading {
		return 0
	}
	if p.decoder == nil {
		return 0
	}
	pos, err := p.decoder.Seek(0, io.SeekCurrent)
	if err != nil {
		p.Err = err
		return 0
	}
	res := float32(float32(pos) / float32(p.decoder.Length()))
	return res
}

func (p *Player) TotalSec() int {
	if p.decoder == nil {
		return 0
	}
	return int(float32(p.decoder.Length()) / float32(p.decoder.SampleRate()) * 0.25)
}

func (p *Player) Now() int {
	return int(float32(p.TotalSec()) * p.Progress())
}

func fzero(text string, nlength int) string {
	ntext := text
	for i := len(text); i < nlength; i++ {
		ntext = "0" + ntext
	}
	return ntext
}

func TimeString(t int) string {
	currentSec := strconv.Itoa(t % 60)
	currentSec = fzero(currentSec, 2)
	currentMin := strconv.Itoa(t / 60)
	currentMin = fzero(currentMin, 2)
	now := currentMin + ":" + currentSec
	return now
}
