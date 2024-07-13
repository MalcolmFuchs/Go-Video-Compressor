package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/MalcolmFuchs/Go-Video-Compressor/api"
)

type video struct {
	input  string
	output string
}

func main() {
	var wg sync.WaitGroup

	files := api.GetFiles("./files/input")

	var videos []video
	for _, file := range files {
		if strings.HasSuffix(file, ".mp4") || strings.HasSuffix(file, ".avi") || strings.HasSuffix(file, ".mov") {
			videos = append(videos, video{input: "./files/input/" + file, output: "./files/output/" + "compressed_" + file})
		}
	}

	for _, v := range videos {
		wg.Add(1)
		go func(v video) {
			defer wg.Done()
			start := time.Now()
			err := api.CompressVideo(v.input, v.output, &wg)
			elapsed := time.Since(start)
			if err != nil {
				fmt.Printf("Fehler bei der Komprimierung von %s: %v\n", v.input, err)
			} else {
				fmt.Printf("Die Komprimierung von %s hat %v gedauert.\n", v.input, elapsed)
			}
		}(v)
	}

	wg.Wait()
	fmt.Println("Alle Videos wurden komprimiert.")
}
