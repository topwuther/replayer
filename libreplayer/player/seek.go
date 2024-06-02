package player

func (p *Player) Seek(offset float32, whence int) int64 {
	var now int64
	if p.decoder == nil {
		return 0
	}
	length := p.decoder.Length()
	current := int64(float32(length) * offset)
	current = current - (current % 4)
	now, p.Err = p.decoder.Seek(current, whence)
	if p.Err != nil {
		return 0
	}
	if !p.IsPlaying() {
		p.Play(p.music)
	}
	return now
}
