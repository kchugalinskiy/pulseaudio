package simple

import (
	"fmt"

	"github.com/kchugalinskiy/pulseaudio/sample"
)

// Example shows basic echo usage of pulseaudio framework. Don't turn your volume too loud before launching it.
// There is one detail you should know: this sample won't work correctly unless you create a window. That's
// because of pulseaudio :-)
func Example() {
	f := sample.Spec{
		Format:   sample.SampleFormatS32LE,
		Rate:     44100,
		Channels: 1,
	}
	play, err := New("", "echo example play", StreamDirectionPlayback, "", "Music", &f, nil, nil)
	if err != nil {
		fmt.Printf("creating sample play: %v\n", err)
		return
	}
	defer play.Close()

	rec, err := New("", "echo example rec", StreamDirectionRecord, "", "Music", &f, nil, nil)
	if err != nil {
		fmt.Printf("creating sample rec: %v\n", err)
		return
	}
	defer rec.Close()

	for {
		b, err := rec.Read8(1024)
		if err != nil {
			fmt.Printf("reading audio: %v\n", err)
			continue
		}

		if err = play.Write8(b); err != nil {
			fmt.Printf("writing audio: %v\n", err)
			continue
		}
	}

	// Output
}
