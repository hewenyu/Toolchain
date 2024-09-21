# non-streaming-decode-files

copy from [sherpa-onnx](https://github.com/k2-fsa/sherpa-onnx/)


## run-nemo-ctc

```bash
#!/usr/bin/env bash

set -ex

if [ ! -d sherpa-onnx-nemo-ctc-en-conformer-medium ]; then
  curl -SL -O https://github.com/k2-fsa/sherpa-onnx/releases/download/asr-models/sherpa-onnx-nemo-ctc-en-conformer-medium.tar.bz2
  tar xvf sherpa-onnx-nemo-ctc-en-conformer-medium.tar.bz2
  rm sherpa-onnx-nemo-ctc-en-conformer-medium.tar.bz2
fi

./build/sherpa_onnx/no_streaming_decode_files \
  --nemo-ctc ./sherpa-onnx-nemo-ctc-en-conformer-medium/model.onnx \
  --tokens ./sherpa-onnx-nemo-ctc-en-conformer-medium/tokens.txt \
  --model-type nemo_ctc \
  --debug 0 \
  ./sherpa-onnx-nemo-ctc-en-conformer-medium/test_wavs/0.wav
```

## run-paraformer-itn

```bash
#!/usr/bin/env bash

set -ex

if [ ! -d sherpa-onnx-paraformer-zh-2023-09-14 ]; then
  curl -SL -O https://github.com/k2-fsa/sherpa-onnx/releases/download/asr-models/sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
  tar xvf sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
  rm sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
fi

if [ ! -f ./itn-zh-number.wav ]; then
  curl -SL -O https://github.com/k2-fsa/sherpa-onnx/releases/download/asr-models/itn-zh-number.wav
fi

if [ ! -f ./itn_zh_number.fst ]; then
  curl -SL -O https://github.com/k2-fsa/sherpa-onnx/releases/download/asr-models/itn_zh_number.fst
fi

go mod tidy
go build

./non-streaming-decode-files \
  --paraformer ./sherpa-onnx-paraformer-zh-2023-09-14/model.int8.onnx \
  --tokens ./sherpa-onnx-paraformer-zh-2023-09-14/tokens.txt \
  --model-type paraformer \
  --rule-fsts ./itn_zh_number.fst \
  --debug 0 \
```

## run-paraformer

```bash
#!/usr/bin/env bash

set -ex

if [ ! -d sherpa-onnx-paraformer-zh-2023-09-14 ]; then
  curl -SL -O https://github.com/k2-fsa/sherpa-onnx/releases/download/asr-models/sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
  tar xvf sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
  rm sherpa-onnx-paraformer-zh-2023-09-14.tar.bz2
fi

go mod tidy
go build

./non-streaming-decode-files \
  --paraformer ./sherpa-onnx-paraformer-zh-2023-09-14/model.int8.onnx \
  --tokens ./sherpa-onnx-paraformer-zh-2023-09-14/tokens.txt \
  --model-type paraformer \
  --debug 0 \
  ./sherpa-onnx-paraformer-zh-2023-09-14/test_wavs/0.wav
```

## 

```bash

```

## 

```bash

```

## 

```bash

```

## 

```bash

```

## 

```bash

```