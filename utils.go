package GOFFMPEG

import (
	"os"
)

// Create the file in projectTmpPath with extensions like (e.g., /tmp/myproject.mp3, /tmp/myproject.mp4)
func createTempOutputFile(fileExtension string) (string, error) {

	// Generate a unique temporary file name
	tmpfile, err := os.CreateTemp(FFMPEG_TempDir, "tmpmedia_")
	if err != nil {
		return "", err
	}
	defer tmpfile.Close()

	// Rename the temporary file with the desired media file extension
	newFilePath := tmpfile.Name() + fileExtension
	err = os.Rename(tmpfile.Name(), newFilePath)
	if err != nil {
		return "", err
	}

	return newFilePath, nil
}

func createDirectoryIfNotExists(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// Directory does not exist, so create it
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err // Some other error occurred while checking directory existence
	}

	return nil
}
