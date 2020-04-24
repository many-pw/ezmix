package mix

import (
	"io"
	"os"
)

func WavLoad(path string) (out []Sample, specs *AudioSpec) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("File not found: " + path)
	}
	file, _ := os.Open(path)
	reader, err := NewReader(file)
	if err != nil {
		panic(err)
	}
	specs = &AudioSpec{
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
