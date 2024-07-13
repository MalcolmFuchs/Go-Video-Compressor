package api

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func CompressVideo(inputFile, outputFile string, wg *sync.WaitGroup) error {

	defer wg.Done()
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-vcodec", "libx264", "-crf", "28", outputFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("fehler bei der videokomprimierung von %s: %w", inputFile, err)
	}

	fmt.Println("Videokomprimierung erfolgreich:", inputFile)

	return nil

}
