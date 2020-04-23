package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"ezmix/bind"
	"ezmix/bind/spec"
	"ezmix/mix"
)

var (
	out         string
	profileMode string
	sampleHz    = float64(48000)
	specs       = spec.AudioSpec{
		Freq:     sampleHz,
		Format:   spec.AudioF32,
		Channels: 2,
	}
	bpm     = 20
	step    = time.Minute / time.Duration(bpm*4)
	loops   = 8
	prefix  = "sound/"
	marac   = "maracas.wav"
	clhat   = "cl_hihat.wav"
	pattern = []string{
		clhat,
		marac,
	}
)

func main() {
	bind.UseOutputString("portaudio")
	specs.Validate()
	bind.SetOutputCallback(mix.NextSample)
	bind.Configure(specs)
	mix.Configure(specs)
	mix.SetSoundsPath(prefix)

	t := 1 * time.Second
	for n := 0; n < loops; n++ {
		for s := 0; s < len(pattern); s++ {
			mix.SetFire(
				pattern[s], t+time.Duration(s)*step, 0, 1.0, rand.Float64()*2-1)
		}
		t += time.Duration(len(pattern)) * step
	}
	t += 5 * time.Second

	mix.StartAt(time.Now().Add(1 * time.Second))
	fmt.Printf("Mix: Example - pid:%v playback:%v spec:%v\n", os.Getpid(), out, specs)
	for mix.FireCount() > 0 {
		time.Sleep(1 * time.Second)
	}

}
