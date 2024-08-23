package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func BenchmarkJSONImport(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// Read in the file (do not time)
		b.StopTimer()
		fileNumber := i % files
		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var samples []Sample
		dec := json.NewDecoder(bytes.NewReader(data))
		if err := dec.Decode(&samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
	}
}

func BenchmarkGobImport(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// Read in the file (do not time)
		b.StopTimer()
		fileNumber := i % files
		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.bin", fileNumber))
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		// decode gob
		var samples []Sample
		dec := gob.NewDecoder(bytes.NewReader(data))
		if err := dec.Decode(&samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
	}
}

func BenchmarkJSONImportFile(b *testing.B) {
	for i := 0; i < b.N; i++ {

		fileNumber := i % files
		file, err := os.Open(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
		if err != nil {
			b.Fatal(err)
		}

		var samples []Sample
		dec := json.NewDecoder(file)
		if err := dec.Decode(&samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
		_ = file.Close()
	}
}

func BenchmarkGobImportFile(b *testing.B) {
	for i := 0; i < b.N; i++ {

		fileNumber := i % files
		file, err := os.Open(fmt.Sprintf("testdata/sample_%d.bin", fileNumber))
		if err != nil {
			b.Fatal(err)
		}

		// decode gob
		var samples []Sample
		dec := gob.NewDecoder(file)
		if err := dec.Decode(&samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
	}
}
