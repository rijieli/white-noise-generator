package main

import (
	"crypto/rand"
	"flag"
	"os"

	"github.com/cryptix/wav"
)

func main() {

	bits := flag.Int("bits", 32, "Significant Bits")
	rate := flag.Int("rate", 44100, "Sample Rate")
	filename := flag.String("o", "noise.wav", "Output file name")
	duration := flag.Int("d", 60, "Duration in seconds")
	flag.Parse()

	outputFile, err := os.Create(*filename)
	checkErr(err)
	defer outputFile.Close()

	meta := wav.File{
		Channels:        1,
		SampleRate:      uint32(*rate),
		SignificantBits: uint16(*bits),
	}

	writer, err := meta.NewWriter(outputFile)
	checkErr(err)
	defer writer.Close()

	for i := 0; i < (*duration)*(*rate); i += 1 {
		var sample []byte = make([]byte, *bits/8)

		// Read will fill random number in byte array, parameter sample means sample rate 44100
		_, err = rand.Read(sample)
		checkErr(err)

		err := writer.WriteSample(sample)
		checkErr(err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
