package types

import (
	sherpa "github.com/k2-fsa/sherpa-onnx-go/sherpa_onnx"
)

// type Config sherpa.OnlineRecognizerConfig

func NewConfig() sherpa.OnlineRecognizerConfig {
	return sherpa.OnlineRecognizerConfig{
		FeatConfig: sherpa.FeatureConfig{SampleRate: 16000, FeatureDim: 80},
		ModelConfig: sherpa.OnlineModelConfig{
			Transducer: sherpa.OnlineTransducerModelConfig{
				Encoder: "encoder.onnx",
				Decoder: "decoder.onnx",
				Joiner:  "joiner.onnx",
			},
			Tokens:     "tokens.txt",
			NumThreads: 1,
			Provider:   "cpu",
			Debug:      0,
			ModelType:  "zipformer2",
		},
		DecodingMethod:          "greedy_search",
		EnableEndpoint:          1,
		Rule1MinTrailingSilence: 2.4,
		Rule2MinTrailingSilence: 1.2,
		Rule3MinUtteranceLength: 20,
	}
}
