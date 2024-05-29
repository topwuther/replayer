package record

import (
	"errors"
	"time"

	"github.com/jfreymuth/pulse"
)

type Record struct {
	Err      error
	filename string
	file     *File
	record   *pulse.Client
	rstream  *pulse.RecordStream
	play     *pulse.Client
	pstream  *pulse.PlaybackStream
}

func (r *Record) SetSourceByName(name string) {
	var err error
	r.record, err = pulse.NewClient()
	if err != nil {
		r.Err = err
		return
	}
	r.play, err = pulse.NewClient()
	if err != nil {
		r.Err = err
		return
	}

	// defer r.play.Close()
	// defer r.record.Close()
	devices, err := r.record.ListSources()
	if err != nil {
		r.Err = err
		return
	}
	for i := 0; i < len(devices); i++ {
		if devices[i].Name() == name {
			pulse.RecordSource(devices[i])
			return
		}
	}
	r.Err = errors.New("device not found")
}

func (r *Record) Start() {
	var err error
	r.record, err = pulse.NewClient()
	if err != nil {
		r.Err = err
		return
	}

	r.file = CreateFile(r.filename, 44100, 1)
	r.rstream, err = r.record.NewRecord(pulse.Float32Writer(r.file.Write), pulse.RecordBufferFragmentSize(1024), pulse.RecordSampleRate(44100))
	if err != nil {
		r.Err = err
		return
	}
	r.pstream, err = r.play.NewPlayback(pulse.Float32Reader(r.file.Read), pulse.PlaybackBufferSize(1024), pulse.PlaybackSampleRate(44100))
	if err != nil {
		r.Err = err
		return
	}
	r.rstream.Start()
	go func() {
		time.Sleep(time.Second)
		r.pstream.Start()
	}()
}

func (r *Record) Stop() {
	r.rstream.Stop()
	r.file.Close()
	r.record.Close()
}

func NewRecord(f string) *Record {
	return &Record{
		filename: f,
	}
}
