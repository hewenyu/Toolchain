# non-streaming-decode-files

copy from [sherpa-onnx](https://github.com/k2-fsa/sherpa-onnx/)

# Introduction

This examples shows how to use the golang package of [sherpa-onnx][sherpa-onnx]
for real-time speech recognition from microphone.

It uses <https://github.com/csukuangfj/portaudio-go>
to read the microphone and you have to install `portaudio` first.

One Debain, you can use

```zsh
apt-get install portaudio19-dev
```

On macOS, you can use

```
brew install portaudio
```

and it will install `portaudio` into `/usr/local/Cellar/portaudio/19.7.0`.
You need to set the following environment variable
```
export PKG_CONFIG_PATH=/usr/local/Cellar/portaudio/19.7.0
```

so that `pkg-config --cflags --libs portaudio-2.0` can run successfully.

[sherpa-onnx]: https://github.com/k2-fsa/sherpa-onnx


## 


```bash 
$ dpkg -l | grep portaudio

ii  libportaudio2:amd64                   19.6.0-1.1                         amd64        Portable audio I/O - shared library
ii  libportaudiocpp0:amd64                19.6.0-1.1                         amd64        Portable audio I/O C++ bindings - shared library
ii  portaudio19-dev:amd64                 19.6.0-1.1                         amd64        Portable audio I/O - development files
```