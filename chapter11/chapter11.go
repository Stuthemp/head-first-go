package main

import (
	"awesomeProject/chapter11/gadget"
	"fmt"
)

func genericsTest(...interface{}) {

}

func innerGenericTest(...any) {

}

func playList(device gadget.AudioDevice, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	recorder, ok := device.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Device is not recorder, it is", device)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	recorder := gadget.TapeRecorder{}
	mixtape := []string{"Shake it off", "I kissed the girl", "friday"}
	playList(player, mixtape)
	playList(recorder, mixtape)
	genericsTest(player, mixtape)
	innerGenericTest(player, recorder, mixtape)
}
