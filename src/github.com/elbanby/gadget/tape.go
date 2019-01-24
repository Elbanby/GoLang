package gadget

import "fmt"

type TapePlayer struct {
	Battaries string
}

func (t *TapePlayer) Play(song string) {
	fmt.Println("Playing " + song)
}

func (t *TapePlayer) Stop() {
	fmt.Println("Stopping music")
}

type TapeRecorder struct {
	Microphones int
}

func (t *TapeRecorder) Play(song string) {
	fmt.Println("Playing " + song)
}

func (t *TapeRecorder) Stop() {
	fmt.Println("Stopping")
}

func (t *TapeRecorder) Record() {
	fmt.Println("Recording")
}
