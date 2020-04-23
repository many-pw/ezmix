// Package bind is for modular binding of mix to audio interface
package bind

import (
	"time"

	"ezmix/bind/hardware/null"
	"ezmix/bind/hardware/portaudio"
	"ezmix/bind/opt"
	"ezmix/bind/sample"
	"ezmix/bind/spec"
	"ezmix/bind/wav"
)

// Configure begins streaming to the bound out audio interface, via a callback function
func ApiConfigure(s spec.AudioSpec) {
	sample.ConfigureOutput(s)
	switch useOutput {
	case opt.OutputPortAudio:
		portaudio.ConfigureOutput(s)
	case opt.OutputWAV:
		wav.ConfigureOutput(s)
	case opt.OutputNull:
		null.ConfigureOutput(s)
	}
}

func ApiIsDirectOutput() bool {
	return useOutput == opt.OutputWAV
}

// SetMixNextOutFunc to stream mix out from mix
func ApiSetOutputCallback(fn sample.OutNextCallbackFunc) {
	sample.SetOutputCallback(fn)
}

// OutputStart requires a known length
func ApiOutputStart(length time.Duration) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputStart(length)
	case opt.OutputNull:
		// do nothing
	}
}

// OutputNext using the configured writer.
func ApiOutputNext(numSamples spec.Tz) {
	switch useOutput {
	case opt.OutputWAV:
		wav.OutputNext(numSamples)
	case opt.OutputNull:
		// do nothing
	}
}

// LoadWAV into a buffer
func ApiLoadWAV(file string) ([]sample.Sample, *spec.AudioSpec) {
	switch useLoader {
	case opt.InputWAV:
		return wav.Load(file)
	default:
		return make([]sample.Sample, 0), &spec.AudioSpec{}
	}
}

// Teardown to close all hardware bindings
func ApiTeardown() {
	switch useOutput {
	case opt.OutputPortAudio:
		portaudio.TeardownOutput()
	case opt.OutputWAV:
		wav.TeardownOutput()
	case opt.OutputNull:
		// do nothing
	}
}

// UseLoader to select the file loading interface
func ApiUseLoader(opt opt.Input) {
	useLoader = opt
}

// UseLoaderString to select the file loading interface by string
func ApiUseLoaderString(loader string) {
	switch loader {
	case string(opt.InputWAV):
		useLoader = opt.InputWAV
	default:
		panic("No such Loader: " + loader)
	}
}

// UseOutput to select the outback interface
func ApiUseOutput(opt opt.Output) {
	useOutput = opt
}

// UseOutputString to select the outback interface by string
func ApiUseOutputString(output string) {
	switch output {
	case string(opt.OutputPortAudio):
		useOutput = opt.OutputPortAudio
	case string(opt.OutputWAV):
		useOutput = opt.OutputWAV
	case string(opt.OutputNull):
		useOutput = opt.OutputNull
	default:
		panic("No such Output: " + output)
	}
}

var (
	useLoader = opt.InputWAV
	useOutput = opt.OutputPortAudio
)
