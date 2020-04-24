// Package source models a single audio source
package mix

import (
	"math"
)

func SourceConfigure(s AudioSpec) {
	masterChannelsFloat = float64(s.Channels)
	masterSpec = &s
}

// New Source from a "URL" (which is actually only a file path for now)
func SourceNew(URL string) *Source {
	// TODO: implement true URL (for now, it's being used as a path)
	s := &Source{
		state: STAGED,
		URL:   URL,
	}
	s.load()
	return s
}

// Source stores a series of Samples in Channels across Time, for audio playback.
type Source struct {
	URL string
	// private
	sample    []Sample
	maxTz     Tz
	audioSpec *AudioSpec
	state     stateEnum
}

// SampleAt at a specific Tz, volume (0 to 1), and pan (-1 to +1)
func (s *Source) SampleAt(at Tz, vol float64, pan float64) (out []Value) {
	out = make([]Value, masterSpec.Channels)
	if at < s.maxTz {
		// if s.sample[at] != 0 {
		// 	debug.Printf("*Source[%v].SampleAt(%v): %v\n", s.URL, at, s.sample[at])
		// }
		if masterSpec.Channels == s.audioSpec.Channels { // same # channels; easier maths
			for c := int(0); c < masterSpec.Channels; c++ {
				out[c] = volume(float64(c), vol, pan) * s.sample[at].Values[c]
			}
		} else { // need to map # source channels to # destination channels
			tc := float64(s.audioSpec.Channels)
			for c := int(0); c < masterSpec.Channels; c++ {
				out[c] = volume(float64(c), vol, pan) * s.sample[at].Values[int(math.Floor(tc*float64(c)/masterChannelsFloat))]
			}
		}
	}
	return
}

// Length of the source audio in Tz
func (s *Source) Length() Tz {
	return s.maxTz
}

// Spec of the source audio
func (s *Source) Spec() *AudioSpec {
	return s.audioSpec
}

// Teardown the source audio and release its memory.
func (s *Source) Teardown() {
	s.sample = nil
}

/*
 *
 private */

var (
	masterChannelsFloat float64
	sourceMasterSpec    *AudioSpec
)

type stateEnum uint

const (
	STAGED stateEnum = iota
	LOADING
	READY
	// it is assumed that all alive states are < FINISHED
	// FINISHED
	// FAILED
)

func (s *Source) load() {
	s.state = LOADING
	s.sample, s.audioSpec = ApiLoadWAV(s.URL)
	if s.audioSpec == nil {
		// TODO: handle errors loading file
		//debug.Printf("could not load WAV %s\n", s.URL)
	}
	s.maxTz = Tz(len(s.sample))
	s.state = READY
}

// volume (0 to 1), and pan (-1 to +1)
// TODO: ensure implicit panning of source channels! e.g. 2 channels is full left, full right.
func volume(channel float64, volume float64, pan float64) Value {
	if pan == 0 {
		return Value(volume)
	} else if pan < 0 {
		return Value(math.Max(0, 1+pan*channel/masterChannelsFloat))
	} else { // pan > 0
		return Value(math.Max(0, 1-pan*channel/masterChannelsFloat))
	}
}
