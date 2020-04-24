// Package fire model an audio source playing at a specific time
package mix

// New Fire to represent a single audio source playing at a specific time in the future.
func FireNew(source string, beginTz Tz, endTz Tz, volume float64, pan float64) *Fire {
	// debug.Printf("NewFire(%v, %v, %v, %v, %v)\n", source, beginTz, endTz, volume, pan)
	s := &Fire{
		/* setup */
		Source:  source,
		Volume:  volume,
		Pan:     pan,
		BeginTz: beginTz,
		EndTz:   endTz,
		/* playback */
		state: fireStateReady,
	}
	return s
}

// Fire represents a single audio source playing at a specific time in the future.
type Fire struct {
	/* setup */
	BeginTz Tz
	EndTz   Tz
	Source  string
	Volume  float64 // 0 to 1
	Pan     float64 // -1 to +1
	/* playback */
	nowTz Tz
	state fireStateEnum
}

// At the series of Tz it's playing for, return the series of Tz corresponding to source audio.
func (f *Fire) FireAt(at Tz) (t Tz) {
	//	debug.Printf("*Fire[%s].At(%v vs %v)\n", f.Source, at, f.BeginTz)
	switch f.state {
	case fireStateReady:
		if at >= f.BeginTz {
			f.state = fireStatePlay
			f.nowTz++
		}
	case fireStatePlay:
		t = f.nowTz
		f.nowTz++
		if f.EndTz != 0 {
			if at >= f.EndTz {
				f.state = fireStateDone
			}
		} else {
			f.EndTz = f.BeginTz + f.sourceLength()
		}
	case fireStateDone:
		// garbage collection
	}
	return
}

// IsAlive the Fire?
func (f *Fire) IsAlive() bool {
	return f.state < fireStateDone
}

// IsPlaying the Fire?
func (f *Fire) IsPlaying() bool {
	return f.state == fireStatePlay
}

// Teardown the Fire and release its memory
func (f *Fire) Teardown() {
	// TODO: confirm that all memory of this object is released when its pointer is deleted from the *Mixer.fires slice, else make sure it does get released somehow
}

/*
 *
 private */

type fireStateEnum uint

const (
	fireStateReady fireStateEnum = 1
	fireStatePlay  fireStateEnum = 2
	// it is assumed that all alive states are < SOURCE_FINISHED
	fireStateDone fireStateEnum = 6
)

func (f *Fire) sourceLength() Tz {
	return GetLength(f.Source)
}
