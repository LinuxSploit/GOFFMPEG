package GOFFMPEG

import (
	"fmt"
	"os"
)

var (
	FFMPEG_TempDir string
)

func init() {
	FFMPEG_TempDir = os.TempDir() + "/_ffmpeg_tmpfiles101"
	err := createDirectoryIfNotExists(FFMPEG_TempDir)
	if err != nil {
		FFMPEG_TempDir = os.TempDir()
	}
}

// FFmpeg represents the FFmpeg wrapper.
type FFmpeg struct {
	ffmpegPath string // Path to the FFmpeg executable
}

// NewFFmpeg creates a new FFmpeg instance with the specified path.
func NewFFmpeg(ffmpegPath string) *FFmpeg {
	return &FFmpeg{ffmpegPath: ffmpegPath}
}

// TranscodeVideo transcodes a video file to a different format.
// example: outputVideoPath,err := TranscodeVideo("/home/linuxsploit/demo.mp4", ".mkv)
func (f FFmpeg) TranscodeVideo(inputPath, OutPutFormatExtension string) (string, error) {

	// Define the output file path
	outputPath, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}

	args := []string{"-i", inputPath, outputPath, "-y"}

	return outputPath, f.Run(args)
}

// ExtractVideo extracts video stream from a video file with the specified output format extension and video codec options.
// It takes the input video file path, output file path with the desired extension (e.g., ".mp4", ".avi"), and a boolean flag to control video codec copying.
// It returns the path to the extracted video file and any encountered errors.
// Example: outputVideoPath, err := ExtractVideo("/home/linuxsploit/demo.mp4", ".mp4", false)
func (f FFmpeg) ExtractVideo(inputPath, OutPutFormatExtension string, copyVideoCodec bool) (string, error) {
	// Define the output file path for the extracted video
	outputVideo, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}

	// Generate FFmpeg command to extract the video
	args := []string{"-i", inputPath, "-an"} // "-an" to disable audio

	// Add video codec options based on the boolean flag
	if copyVideoCodec {
		args = append(args, "-vcodec", "copy") // Copy video codec
	}

	args = append(args, outputVideo, "-y") // Output file path

	// Run the FFmpeg command
	return outputVideo, f.Run(args)
}

// ExtractAudio extracts audio stream from a video file with the specified output format extension and audio codec options.
// It takes the input video file path, output file path with the desired extension (e.g., ".mp3", ".ogg"), and a boolean flag to control audio codec copying.
// It returns the path to the extracted audio file and any encountered errors.
// example: outputVideoPath,err := ExtractAudio("/home/linuxsploit/demo.mp4",".mp3",false)
func (f FFmpeg) ExtractAudio(inputPath, OutPutFormatExtension string, copyAudioCodec bool) (string, error) {
	// Define the output file path for the extracted audio
	outputAudio, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}
	// Generate FFmpeg command to extract the audio
	args := []string{"-i", inputPath, "-vn"}

	// Add audio codec options based on the boolean flag
	if copyAudioCodec {
		args = append(args, "-acodec", "copy") // Copy audio codec
	}

	args = append(args, outputAudio, "-y") // Output file path

	// Run the FFmpeg command
	return outputAudio, f.Run(args)
}

// ExtractVideoClip extracts a video clip from a specified time range in a video file.
// It takes the input video file path, start time in seconds, end time in seconds,
// and boolean flags to control video and audio options.
// It returns the path to the extracted video clip file and any encountered errors.
// example: ExtractVideoClip("/home/linuxsploit/demo.mp4", ".mp4", 25, 50, false, false)
func (f FFmpeg) ExtractVideoClip(inputPath, OutPutFormatExtension string, startSec, endSec float64, copyVideoCodec, disableAudio bool) (string, error) {
	// Ensure startSec and endSec are valid time ranges
	if startSec < 0 || endSec <= startSec {
		return "", fmt.Errorf("invalid time range for video extraction")
	}

	// Define the output file path for the extracted video clip
	outputVideo, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}

	// Generate FFmpeg command to extract the video clip
	args := []string{
		"-i", inputPath,
		"-ss", fmt.Sprintf("%f", startSec), // Start time
		"-to", fmt.Sprintf("%f", endSec), // End time
	}

	// Add video and audio options based on the boolean flags
	if copyVideoCodec {
		args = append(args, "-c:v", "copy") // Copy video codec
	}
	if disableAudio {
		args = append(args, "-an") // Disable audio
	}

	args = append(args, outputVideo, "-y") // Output file path

	// Run the FFmpeg command
	return outputVideo, f.Run(args)
}

// ExtractAudioClip extracts an audio clip from a specified time range in a video file.
// It takes the input video file path, start time in seconds, end time in seconds,
// and boolean flags to control video and audio options.
// It returns the path to the extracted audio clip file and any encountered errors.
// example: ExtractAudioClip("/home/linuxsploit/demo.mp4", ".mp3", 25, 50, false)
func (f FFmpeg) ExtractAudioClip(inputPath, OutPutFormatExtension string, startSec, endSec float64, copyAudioCodec bool) (string, error) {
	// Ensure startSec and endSec are valid time ranges
	if startSec < 0 || endSec <= startSec {
		return "", fmt.Errorf("invalid time range for audio extraction")
	}

	// Define the output file path for the extracted audio clip
	outputAudio, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}

	// Generate FFmpeg command to extract the audio clip
	args := []string{
		"-i", inputPath,
		"-ss", fmt.Sprintf("%f", startSec), // Start time
		"-to", fmt.Sprintf("%f", endSec), // End time
		"-vn", // Disable video
	}

	if copyAudioCodec {
		args = append(args, "-c:a", "copy") // Copy audio codec
	}

	args = append(args, outputAudio, "-y") // Output file path

	// Run the FFmpeg command
	return outputAudio, f.Run(args)
}

// ConcatVideos concatenates multiple video files into a single video.
// []string videos should have same resolution.
// Example:
// ConcatVideos(
//
//		[]string{
//			"/home/linuxsploit/demo1.mp4",
//			"/home/linuxsploit/demo2.mp4",
//			"/home/linuxsploit/demo3.mp4",
//		},
//		".mkv",
//	)
func (f FFmpeg) ConcatVideos(inputPaths []string, OutPutFormatExtension string) (string, error) {
	// Define the output file path for the concatenated video
	outputPath, err := createTempOutputFile(OutPutFormatExtension)
	if err != nil {
		return "", err
	}

	// Create a list of arguments for FFmpeg
	args := []string{}

	// Add the input files
	for _, path := range inputPaths {
		args = append(args, "-i", path)
	}

	// Add the filter_complex and output file options
	args = append(args, "-filter_complex", "concat=n="+fmt.Sprint(len(inputPaths))+":v=1:a=1[outv][outa]", "-map", "[outv]", "-map", "[outa]", outputPath, "-y")

	// Run the FFmpeg command
	err = f.Run(args)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
