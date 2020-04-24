package mix

import (
	"github.com/gordonklaus/portaudio"
)

var outPortaudioStream *portaudio.Stream

func PortAudioConfigureOutput(s AudioSpec) {
	var err error
	outSpec = &s
	portaudio.Initialize()
	outPortaudioStream, err = portaudio.OpenDefaultStream(0, s.Channels, s.Freq, 0, outPortaudioStreamCallback)
	noErr(err)
	noErr(outPortaudioStream.Start())
}

func PortAudioTeardownOutput() {
	//	noErr(out.Stop())
	//	noErr(out.Close())
	portaudio.Terminate()
}

var (
	outSpec *AudioSpec
)

func outPortaudioStreamCallback(out [][]float32) {
	var smp []Value
	for s := range out[0] {
		smp = SampleOutNext()
		for c := 0; c < outSpec.Channels; c++ {
			out[c][s] = float32(smp[c])
		}
	}
}

func noErr(err error) {
	if err != nil {
		panic(err)
	}
}
