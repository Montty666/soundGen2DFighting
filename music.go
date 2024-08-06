package soundGen2DFighting

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"

	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func logg(txt string) {
	fmt.Println(txt)
}

func PlayMusic() {
	logg("start")
	f, err := os.Open("./caravans.mp3")
	if err != nil {
		logg("err1")
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		logg("error")
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
