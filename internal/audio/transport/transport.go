package audio

import "github.com/hewenyu/toolchain/internal/audio/types"

// Transport is an interface that represents a transport for audio data.
type Transport interface {
	Consume() <-chan types.TextData // consume data
	Connect(Transport) error        // need connect
	SendMessage([]byte) error       // send message
}
