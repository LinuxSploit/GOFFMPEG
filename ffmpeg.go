package GOFFMPEG

import (
	"fmt"
	"os"
	"os/exec"
)

// FFmpeg represents the FFmpeg wrapper.
type FFmpeg struct {
	FFmpegPath string // Path to the FFmpeg executable
}

// NewFFmpeg creates a new FFmpeg instance with the specified path.
func NewFFmpeg(ffmpegPath string) *FFmpeg {
	return &FFmpeg{FFmpegPath: ffmpegPath}
}

// TranscodeVideo transcodes a video file to a different format.
func (f *FFmpeg) TranscodeVideo(inputPath, outputPath string) error {
	cmd := exec.Command(f.FFmpegPath, "-i", inputPath, outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ExtractAudio extracts audio from a video file.
func (f *FFmpeg) ExtractAudio(inputPath, outputPath string) error {
	cmd := exec.Command(f.FFmpegPath, "-i", inputPath, "-vn", "-acodec", "copy", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ConcatVideos concatenates multiple video files into a single video.
func (f *FFmpeg) ConcatVideos(inputPaths []string, outputPath string) error {
	// Generate a list of input files in the required format for FFmpeg.
	inputFiles := ""
	for _, path := range inputPaths {
		inputFiles += fmt.Sprintf("-i %s ", path)
	}

	cmd := exec.Command(f.FFmpegPath, inputFiles, "-filter_complex", "concat=n="+fmt.Sprint(len(inputPaths))+":v=1:a=1[outv][outa]", "-map", "[outv]", "-map", "[outa]", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
