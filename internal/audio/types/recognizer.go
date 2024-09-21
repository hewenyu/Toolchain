package types

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

func NewOnlineRecognizer(config *sherpa.OnlineRecognizerConfig) *sherpa.OnlineRecognizer {
	return sherpa.NewOnlineRecognizer(config)
}
