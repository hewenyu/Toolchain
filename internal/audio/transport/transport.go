package audio

import (
	"github.com/hewenyu/toolchain/internal/audio/types"
	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

// Transport is an interface that represents a transport for audio data.
type Transport interface {
	Consume() <-chan types.TextData                // consume data
	Connect(recognizer *sherpa.OnlineStream) error // need connect
}
