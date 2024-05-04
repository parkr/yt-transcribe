package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s [flags] <video-url>\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	videoURL := flag.Arg(0)
	if videoURL == "" {
		flag.Usage()
		return
	}

	if err := runCmd("youtube-dl", []string{"-f", "bestaudio", "-o", "audio.%(ext)s", videoURL}); err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var audioFile string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		basename := filepath.Base(file.Name())
		if strings.TrimSuffix(basename, filepath.Ext(basename)) == "audio" {
			audioFile = file.Name()
			break
		}
	}

	if audioFile == "" {
		log.Fatal("No audio file found")
	}

	if err := runCmd("ffmpeg", []string{"-i", audioFile, "-ar", "16000", "-vn", "audio.wav"}); err != nil {
		log.Fatal(err)
	}

	var whisperModelFile string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		basename := filepath.Base(file.Name())
		if filepath.Ext(basename) == ".bin" {
			whisperModelFile = file.Name()
			break
		}
	}

	if err := runCmd("whisper-cpp", []string{"-m", whisperModelFile, "--output-txt", "audio.wav"}); err != nil {
		log.Fatal(err)
	}
	log.Println("Completed!")
}

func runCmd(cmd string, args []string) error {
	log.Printf("Running %s %v", cmd, args)
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stderr
	c.Stderr = os.Stderr
	return c.Run()
}
