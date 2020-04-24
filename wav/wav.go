// Package wav is direct WAV filo I/O
package wav

import (
	"io"
	"os"

	"ezmix/mix"
)

// Load a WAV file into memory
func Load(path string) (out []mix.Sample, specs *mix.AudioSpec) {
	//	data, sdlSpec := sdl.LoadWAV(file, sdl2Spec(spec))
	// return data, sdl2Unspec(sdlSpec)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File not found: " + path)
	}
	file, _ := os.Open(path)
	reader, err := NewReader(file)
	if err != nil {
		panic(err)
	}
	specs = &mix.AudioSpec{
		Freq:     float64(reader.Format.SampleRate),
		Format:   reader.AudioFormat,
		Channels: int(reader.Format.NumChannels),
	}
	for {
		samples, err := reader.ReadSamples()
		if err == io.EOF {
			break
		}
		out = append(out, samples...)
	}
	return
}

/*
 *
 private */
