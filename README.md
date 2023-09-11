# Go FFmpeg Wrapper Library

This is a simple Go wrapper library for FFmpeg that provides an easy way to perform common video and audio transcoding tasks using FFmpeg's features.

## Features

- Transcode video from one format to another.
- Remove audio from a input video file.
- Extract audio from a  input video file.
- Extract video clip from a input video file.
- Extract audio clip from a input file.
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
	"log"

	"github.com/LinuxSploit/GOFFMPEG"
)

func main() {
	// Initialize the FFmpeg wrapper with the path to the FFmpeg executable.
	ffmpeg := GOFFMPEG.NewFFmpeg("/path/to/ffmpeg")

	// +
	// Example: Transcode a video
	outputPath, err := ffmpeg.TranscodeVideo("/home/linuxsploit/demo1.mp4", ".mkv")
	if err != nil {
		fmt.Println("Error transcoding video:", err)
	}
	log.Println(outputPath)

	// +
	// Example: Remove Audio from video, Extract only video
	outputPath, err = ffmpeg.ExtractVideo("/home/linuxsploit/demo.mp4", ".mp4", false)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
	}
	log.Println(outputPath)

	// +
	// Example: Extract audio from a video
	outputPath, err = ffmpeg.ExtractAudio("/home/linuxsploit/demo.mp4", ".mp3", false)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
	}
	log.Println(outputPath)

	// +
	// Example: Extract Video Clip of 20 seconds to 70 seconds timestamp from a input video
	outputPath, err = ffmpeg.ExtractVideoClip("/home/linuxsploit/demo.mp4", ".mp4", 20, 70, false, false)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
	}
	log.Println(outputPath)

	// +
	// Example: Extract Audio Clip of 20 seconds to 70 seconds timestamp from a input video
	outputPath, err = ffmpeg.ExtractAudioClip("/home/linuxsploit/demo.mp4", ".mp3", 20, 70, false)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
	}
	log.Println(outputPath)

	// +
	// Example: Concatenate multiple videos, input videos should have same resolution
	outputPath, err = ffmpeg.ConcatVideos(
		[]string{
			"/home/linuxsploit/demo1.mp4",
			"/home/linuxsploit/demo2.mp4",
			"/home/linuxsploit/demo3.mp4",
		},
		".mkv",
	)
	if err != nil {
		fmt.Println("Error Concatenating videos:", err)
	}
	log.Println(outputPath)
}
```

## License:
This FFmpeg wrapper library is open-source and distributed under the GPL-3.0 License. See the [LICENSE file](https://github.com/LinuxSploit/GOFFMPEG/blob/main/LICENSE) for details.

## Contribution:
Contributions are welcome! Please feel free to open issues or submit pull requests to enhance this library.

## Disclaimer:
This library is a basic wrapper for FFmpeg and may not cover all FFmpeg features and options. It is intended for common use cases and can be extended to support additional functionality.




