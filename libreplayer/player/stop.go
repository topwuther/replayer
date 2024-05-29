package player

func (p *Player) Stop() {
	p.Err = nil
	if p.file != nil {
		p.file.Close()
		p.file = nil
	}
	if p.stream != nil && (!p.stream.Closed()) {
		p.stream.Close()
	}
}
