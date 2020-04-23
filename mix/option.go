// Package opt specifies valid options
package mix

// OptLoader represents an audio input option
type OptInput string

// OptLoadWav to use Go-Native WAV file I/O
const OptInputWAV OptInput = "wav"

// OptOutput represents an audio output option
type OptOutput string

// OptOutputNull for benchmarking/profiling, because those tools are unable to sample to C-go callback tree
const OptOutputNull OptOutput = "null"

// OptOutputPortAudio to use Portaudio for audio output
const OptOutputPortAudio OptOutput = "portaudio"

// OptOutputWAV to use WAV directly for []byte to stdout
const OptOutputWAV OptOutput = "wav"
