package player

func (p *Player) Resume() {
	p.Err = nil
	p.Play(p.music)
}
