package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	go_json "github.com/goccy/go-json"
)

func BenchmarkJSON(b *testing.B) {
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
		if err := json.Unmarshal(data, &samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
	}
}

// func BenchmarkDecode(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 		// Read in the file (do not time)
// 		b.StopTimer()
// 		fileNumber := i % files
// 		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StartTimer()
//
// 		var samples []Sample
// 		dec := json.NewDecoder(bytes.NewReader(data))
// 		if err := dec.Decode(&samples); err != nil {
// 			b.Fatal(err)
// 		}
// 		if len(samples) != itemsPerFile {
// 			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
// 		}
// 	}
// }

func BenchmarkPartialJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// Read in the file (do not time)
		b.StopTimer()
		fileNumber := i % files
		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()

		var partials []json.RawMessage
		if err := json.Unmarshal(data, &partials); err != nil {
			b.Fatal(err)
		}
		if len(partials) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(partials))
		}
		var samples []Sample
		for _, p := range partials {
			var s Sample
			if err := json.Unmarshal(p, &s); err != nil {
				b.Fatal(err)
			}
			samples = append(samples, s)
		}
	}
}

// func BenchmarkPartialDecode(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 		// Read in the file (do not time)
// 		b.StopTimer()
// 		fileNumber := i % files
// 		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StartTimer()
//
// 		var partials []json.RawMessage
// 		dec := json.NewDecoder(bytes.NewReader(data))
// 		if err := dec.Decode(&partials); err != nil {
// 			b.Fatal(err)
// 		}
// 		if len(partials) != itemsPerFile {
// 			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(partials))
// 		}
// 		var samples []Sample
// 		for _, p := range partials {
// 			var s Sample
// 			itemDec := json.NewDecoder(bytes.NewReader(p))
// 			if err := itemDec.Decode(&p); err != nil {
// 				b.Fatal(err)
// 			}
// 			samples = append(samples, s)
// 		}
// 	}
// }

func BenchmarkGoccyGoJSON(b *testing.B) {
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
		if err := go_json.Unmarshal(data, &samples); err != nil {
			b.Fatal(err)
		}
		if len(samples) != itemsPerFile {
			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
		}
	}
}

// func BenchmarkGoccyGoDecode(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
//
// 		// Read in the file (do not time)
// 		b.StopTimer()
// 		fileNumber := i % files
// 		data, err := os.ReadFile(fmt.Sprintf("testdata/sample_%d.json", fileNumber))
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StartTimer()
//
// 		var samples []Sample
// 		dec := go_json.NewDecoder(bytes.NewReader(data))
// 		if err := dec.Decode(&samples); err != nil {
// 			b.Fatal(err)
// 		}
// 		if len(samples) != itemsPerFile {
// 			b.Fatalf("expected %d samples, got %d", itemsPerFile, len(samples))
// 		}
// 	}
// }
