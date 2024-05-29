package player

import (
	"os"

	"github.com/hajimehoshi/go-mp3"
)

func (p *Player) Seek(offset float32, whence int) int64 {
	var now int64
	p.Stop()
	p.file, p.Err = os.Open(musicdir + p.music.Filename)
	if p.Err != nil {
		return 0
	}
	p.decoder, p.Err = mp3.NewDecoder(p.file)
	now, p.Err = p.decoder.Seek(int64(float32(p.decoder.Length())*offset), whence)
	if p.Err != nil {
		return 0
	}
	p.Pause()
	p.Resume()
	return now
}
