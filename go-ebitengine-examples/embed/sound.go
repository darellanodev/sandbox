package main

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type Sound struct {
	f      fs.File
	player *audio.Player
}


const (
	sampleRate   = 48000
)

func NewSoundFromFile(audioContext *audio.Context, path string) (*Sound, error) {
	s := &Sound{}
	
	f, _ := assets.Open(path)
	s.f = f
	d, _ := wav.DecodeWithoutResampling(f)
	player, _ := audioContext.NewPlayer(d)
	s.player = player

	return s, nil
}

func (s *Sound) Play() error {

	if err := s.player.Rewind(); err != nil {
		return err
	}

	s.player.Play()

	return nil
}

func loadSounds() {

	audioContext := audio.NewContext(sampleRate)

	soundEmbeddedFile, _ = NewSoundFromFile(audioContext, "assets/sounds/fuel_pick.wav")

}