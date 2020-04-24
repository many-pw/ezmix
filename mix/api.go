package mix

import (
	"time"

	"ezmix/wav"
)

// Configure begins streaming to the bound out audio interface, via a callback function
func ApiConfigure(s AudioSpec) {
	SampleConfigureOutput(s)
	switch useOutput {
	case OptOutputPortAudio:
		PortAudioConfigureOutput(s)
	case OptOutputWAV:
		wav.ConfigureOutput(s)
	case OptOutputNull:
		NullConfigureOutput(s)
	}
}

func ApiIsDirectOutput() bool {
	return useOutput == OptOutputWAV
}

// SetMixNextOutFunc to stream mix out from mix
func ApiSetOutputCallback(fn SampleOutNextCallbackFunc) {
	SampleSetOutputCallback(fn)
}

// OutputStart requires a known length
func ApiOutputStart(length time.Duration) {
	switch useOutput {
	case OptOutputWAV:
		wav.OutputStart(length)
	case OptOutputNull:
		// do nothing
	}
}

// OutputNext using the configured writer.
func ApiOutputNext(numSamples Tz) {
	switch useOutput {
	case OptOutputWAV:
		wav.OutputNext(numSamples)
	case OptOutputNull:
		// do nothing
	}
}

// LoadWAV into a buffer
func ApiLoadWAV(file string) ([]Sample, *AudioSpec) {
	switch useLoader {
	case OptInputWAV:
		return wav.Load(file)
	default:
		return make([]Sample, 0), &AudioSpec{}
	}
}

func ApiTeardown() {
	switch useOutput {
	case OptOutputPortAudio:
		PortAudioTeardownOutput()
	case OptOutputWAV:
		wav.TeardownOutput()
	case OptOutputNull:
		// do nothing
	}
}

// UseLoader to select the file loading interface
func ApiUseLoader(opt OptInput) {
	useLoader = opt
}

// UseLoaderString to select the file loading interface by string
func ApiUseLoaderString(loader string) {
	switch loader {
	case string(OptInputWAV):
		useLoader = OptInputWAV
	default:
		panic("No such Loader: " + loader)
	}
}

// UseOutput to select the outback interface
func ApiUseOutput(opt OptOutput) {
	useOutput = opt
}

// UseOutputString to select the outback interface by string
func ApiUseOutputString(output string) {
	switch output {
	case string(OptOutputPortAudio):
		useOutput = OptOutputPortAudio
	case string(OptOutputWAV):
		useOutput = OptOutputWAV
	case string(OptOutputNull):
		useOutput = OptOutputNull
	default:
		panic("No such Output: " + output)
	}
}

var (
	useLoader = OptInputWAV
	useOutput = OptOutputPortAudio
)
