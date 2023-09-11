package GOFFMPEG

import (
	"os"
	"os/exec"
)

func (f FFmpeg) Run(args []string) error {

	cmd := exec.Command(f.ffmpegPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
