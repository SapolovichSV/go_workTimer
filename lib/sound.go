package lib

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func PlaySound(pathToSound string) error {
	f, err := os.Open(pathToSound)
	if err != nil {
		panic(err)
	}
	s, format, err := mp3.Decode(f)
	if err != nil {
		log.Printf("Error decoding mp3 file %v", err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(s)
	return err
}
