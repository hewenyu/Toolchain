package audio

import (
	"log"

	"github.com/gordonklaus/portaudio"
	"github.com/hewenyu/toolchain/internal/audio/types"
	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

// listen microphone data
type MicrophoneTransport struct {
	portaudio.StreamParameters
	text chan types.TextData
}

func NewMicrophoneTransport() *MicrophoneTransport {
	return &MicrophoneTransport{}
}

func (m *MicrophoneTransport) Consume() <-chan types.TextData {
	return m.text
}

func (m *MicrophoneTransport) Connect(stream *sherpa.OnlineStream) error {

	// you can choose another value for 0.1 if you want
	samplesPerCall := int32(m.SampleRate * 0.1) // 0.1 second
	samples := make([]float32, samplesPerCall)
	s, err := portaudio.OpenStream(m.StreamParameters, samples)
	if err != nil {
		log.Fatalf("Failed to open the stream")
		return err
	}
	defer s.Close()
	return nil
}
