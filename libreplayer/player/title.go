package player

import "github.com/topwuther/replayer/ui/values"

func (p *Player) Title() string {
	if p.Downloading {
		return values.GetText("downloading")
	}
	now := TimeString(p.Now())
	total := TimeString(p.TotalSec())
	if p.music == nil || p.music.Name == "" || p.music.Singer == "" {
		return values.GetText("no playing")
	}
	subtitle := p.music.Name + " - " + p.music.Singer
	subtitle += "   " + now + "/" + total
	return subtitle
}
