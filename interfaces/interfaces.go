package main

import (
  "github.com/elbanby/gadget"
)

type Player interface {
  Play(string)
  Stop()
}

func PlayList(device Player, songs []string) {
  for _, song := range songs {
    device.Play(song)
  }
  device.Stop()
}

func main() {
  /*
  When Go decides whether a value satisfies an interface,
  pointer methods aren't included for direct values.
  But they are included for pointers.
  So the solution is to assign a pointer to a Switchto the Toggleablevariable,
  instead of a direct Switchvalue:
  */
  var device Player = &gadget.TapePlayer{}
  mixTape := []string{"Oh my love", "Sweet Child O' mine"}
  PlayList(device, mixTape)
  device = &gadget.TapeRecorder{}
  PlayList(device, mixTape)
}
