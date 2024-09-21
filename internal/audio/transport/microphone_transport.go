package audio

import "github.com/gordonklaus/portaudio"

// listen microphone data
type MicrophoneTransport struct {
	portaudio.StreamParameters
}
