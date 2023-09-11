# Go FFmpeg Wrapper Library

This is a simple Go wrapper library for FFmpeg that provides an easy way to perform common video and audio transcoding tasks using FFmpeg's features.

## Features

- Transcode video from one format to another.
- Extract audio from a video file.
- Concatenate multiple video files into a single video.

## Installation

Before using this library, ensure that you have FFmpeg installed on your system.

```sh
# Install FFmpeg on Ubuntu
sudo apt-get install ffmpeg

# Install FFmpeg on macOS using Homebrew
brew install ffmpeg

# To use this library in your Go project, you can install it via go get:
go get -u github.com/LinuxSploit/GOFFMPEG/
```
# Usage
```go
package main

import (
	"fmt"
	"github.com/LinuxSploit/GOFFMPEG/"
)

func main() {
	// Initialize the FFmpeg wrapper with the path to the FFmpeg executable.
	ffmpeg := ffmpegwrapper.NewFFmpeg("/path/to/ffmpeg")

	// Example: Transcode a video
	err := ffmpeg.TranscodeVideo("input.mp4", "output.mp4")
	if err != nil {
		fmt.Println("Error transcoding video:", err)
	}

	// Example: Extract audio from a video
	err = ffmpeg.ExtractAudio("input.mp4", "output.mp3")
	if err != nil {
		fmt.Println("Error extracting audio:", err)
	}

	// Example: Concatenate multiple videos
	inputPaths := []string{"video1.mp4", "video2.mp4"}
	err = ffmpeg.ConcatVideos(inputPaths, "output.mp4")
	if err != nil {
		fmt.Println("Error concatenating videos:", err)
	}
}
```

## License:
This FFmpeg wrapper library is open-source and distributed under the GPL-3.0 License. See the [LICENSE file](https://github.com/LinuxSploit/GOFFMPEG/blob/main/LICENSE) for details.

## Contribution:
Contributions are welcome! Please feel free to open issues or submit pull requests to enhance this library.

## Disclaimer:
This library is a basic wrapper for FFmpeg and may not cover all FFmpeg features and options. It is intended for common use cases and can be extended to support additional functionality.




