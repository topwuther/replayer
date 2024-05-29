package player

func (p *Player) Pause() {
	p.Err = nil
	if p.music == nil {
		return
	}
	p.stream.Pause()
}
