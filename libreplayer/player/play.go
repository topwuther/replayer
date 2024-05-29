package player

import (
	"errors"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/hajimehoshi/go-mp3"
	"github.com/jfreymuth/pulse"
	"github.com/jfreymuth/pulse/proto"
	"github.com/topwuther/replayer/ui/pages/settings"
)

var musicdir = "/tmp/"

type Player struct {
	Err         error
	file        *os.File
	client      *pulse.Client
	stream      *pulse.PlaybackStream
	decoder     *mp3.Decoder
	music       *Music
	Downloading bool
}

type Music struct {
	UUID     string
	Name     string
	Singer   string
	Filename string
	Content  []byte
	Id       int
}

func (p *Player) SetSinkByName(name string) {
	p.Err = nil
	devices, err := p.client.ListSinks()
	if err != nil {
		p.Err = err
		return
	}
	for i := 0; i < len(devices); i++ {
		if devices[i].Name() == name {
			pulse.PlaybackSink(devices[i])
			return
		}
	}
	p.Err = errors.New("device not found")
}

func (p *Player) Play(music *Music) {
	if p.Downloading {
		return
	}
	go func() {
		var err error
		p.Err = nil
		p.music = music
		realfp := musicdir + music.Filename
		_, err = os.Stat(realfp)
		if os.IsNotExist(err) {
			server := settings.SettingConfig.Server
			client := resty.New()
			p.Downloading = true
			_, err = client.R().SetBody(music).SetResult(&music).Post("http://" + server + "/music/get")
			p.Downloading = false
			if err != nil {
				p.Err = err
				return
			}
			if err := os.WriteFile(realfp, music.Content, 0644); err != nil {
				p.Err = err
				return
			}
		}
		if p.file == nil {
			p.file, err = os.Open(realfp)
			if err != nil {
				p.Err = err
				return
			}
			p.decoder, err = mp3.NewDecoder(p.file)
			if err != nil {
				p.Err = err
				return
			}
		}
		if p.stream != nil && p.stream.Running() {
			return
		}
		p.stream, err = p.client.NewPlayback(pulse.Int16Reader(func(i []int16) (int, error) {
			return decoder(p, i)
		}), pulse.PlaybackSampleRate(p.decoder.SampleRate()), pulse.PlaybackChannels(proto.ChannelMap{
			proto.ChannelLeft,
			proto.ChannelRight,
		}), pulse.PlaybackBufferSize(50))
		if err != nil {
			p.Err = err
			return
		}
		p.stream.Start()
	}()

}

func (p *Player) IsPlaying() bool {
	if p.stream == nil {
		return false
	}
	return p.stream.Running()
}

func decoder(p *Player, out []int16) (int, error) {
	sample := make([]byte, 2)
	for i := range out {
		if _, err := p.decoder.Read(sample); err != nil {
			return i, pulse.EndOfData
		}
		v := int(sample[0]) | (int(sample[1]) << 8)
		out[i] = int16(v)
	}
	return len(out), nil
}

func NewPlayer() *Player {
	var err error
	p := &Player{}
	p.client, err = pulse.NewClient()
	if err != nil {
		p.Err = err
	}
	return p
}
