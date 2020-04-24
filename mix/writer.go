package mix

import (
	"encoding/binary"
	"io"
	"syscall"

	riff "github.com/youpy/go-riff"

	"os"
	"time"
)

func WavConfigureOutput(s AudioSpec) {
	outputSpec = &s
}

func WavOutputStart(length time.Duration) {
	writer = NewWriter(stdout, FormatFromSpec(outputSpec), length)
}

type Writer struct {
	io.Writer
	Format *Format
}

func NewWriter(w io.Writer, format Format, length time.Duration) (writer *Writer) {
	dataSize := uint32(float64(length/time.Second)*float64(format.SampleRate)) * uint32(format.BlockAlign)
	riffSize := 4 + 8 + 16 + 8 + dataSize
	riffWriter := riff.NewWriter(w, []byte("WAVE"), riffSize)

	writer = &Writer{riffWriter, &format}
	riffWriter.WriteChunk([]byte("fmt "), 16, func(w io.Writer) {
		binary.Write(w, binary.LittleEndian, format)
	})
	riffWriter.WriteChunk([]byte("data"), dataSize, func(w io.Writer) {})

	return writer
}

func OutputNext(numSamples Tz) (err error) {
	for n := Tz(0); n < numSamples; n++ {
		//writer.Write(wavOutNextBytes())
	}
	return
}

var (
	stdout     = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	writer     *Writer
	outputSpec *AudioSpec
)
