package main

import (
	"fmt"
	"log"
	"strings"

	portaudio "github.com/gordonklaus/portaudio"
	"github.com/hewenyu/toolchain/internal/audio/types"
	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

func main() {
	err := portaudio.Initialize()
	if err != nil {
		log.Fatalf("Unable to initialize portaudio: %v\n", err)
	}
	defer portaudio.Terminate()

	allDevices, err := portaudio.Devices()
	if err != nil {
		log.Fatalf("Failed to get devices: %v\n", err)
	}
	// Print all devices
	for i, device := range allDevices {
		fmt.Printf("%d: %s\n", i, device.Name)
	}

	// input number of the device
	var deviceNum int
	fmt.Print("Select the input device number: ")
	fmt.Scan(&deviceNum)

	param := portaudio.StreamParameters{}
	param.Input.Device = allDevices[deviceNum]
	param.Input.Channels = 1
	param.Input.Latency = allDevices[deviceNum].DefaultLowInputLatency
	param.SampleRate = 16000
	param.FramesPerBuffer = 0
	param.Flags = portaudio.ClipOff

	// Create a new config
	config := types.NewConfig()

	log.Println("Initializing recognizer (may take several seconds)")
	recognizer := sherpa.NewOnlineRecognizer(&config)
	log.Println("Recognizer created!")
	defer sherpa.DeleteOnlineRecognizer(recognizer)

	stream := sherpa.NewOnlineStream(recognizer)
	defer sherpa.DeleteOnlineStream(stream)

	// you can choose another value for 0.1 if you want
	samplesPerCall := int32(param.SampleRate * 0.1) // 0.1 second

	samples := make([]float32, samplesPerCall)
	s, err := portaudio.OpenStream(param, samples)
	if err != nil {
		log.Fatalf("Failed to open the stream")
	}
	defer s.Close()
	chk(s.Start())

	var last_text string

	segment_idx := 0

	fmt.Println("Started! Please speak")

	for {
		chk(s.Read())
		stream.AcceptWaveform(int(param.SampleRate), samples)

		for recognizer.IsReady(stream) {
			recognizer.Decode(stream)
		}

		text := recognizer.GetResult(stream).Text
		if len(text) != 0 && last_text != text {
			last_text = strings.ToLower(text)
			fmt.Printf("\r%d: %s", segment_idx, last_text)
		}

		if recognizer.IsEndpoint(stream) {
			if len(text) != 0 {
				segment_idx++
				fmt.Println()
			}
			recognizer.Reset(stream)
		}
	}

	// chk(s.Stop())
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
