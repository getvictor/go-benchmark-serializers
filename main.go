package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

type Sample struct {
	String  string    `json:"string"`
	Date    time.Time `json:"date"`
	Number  int       `json:"number"`
	Float   float64   `json:"float"`
	Valid   bool      `json:"valid"`
	Numbers []int     `json:"numbers"`
	Strings []string  `json:"strings"`
}

const files = 1000
const itemsPerFile = 1000

func main() {

	// Create testdata directory if it doesn't exist
	_ = os.Mkdir("testdata", os.ModeDir|0755)

	// Create test data
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < files; i++ {
		var samples []Sample
		for j := 0; j < itemsPerFile; j++ {
			sample := Sample{
				String: uuid.NewString(),
				Date:   time.Now(),
				Number: r.Int(),
				Float:  r.Float64(),
				Valid:  r.Intn(2) == 0,
			}
			numbersLen := r.Intn(10)
			for k := 0; k < numbersLen; k++ {
				sample.Numbers = append(sample.Numbers, r.Int())
			}
			stringsLen := r.Intn(10)
			for k := 0; k < stringsLen; k++ {
				sample.Strings = append(sample.Strings, uuid.NewString())
			}
			samples = append(samples, sample)
		}

		// Serialize to JSON
		jsonBytes, err := json.Marshal(samples)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(fmt.Sprintf("testdata/sample_%d.json", i), jsonBytes, 0644)
		if err != nil {
			panic(err)
		}

		// Serialize to gob
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		err = enc.Encode(samples)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(fmt.Sprintf("testdata/sample_%d.bin", i), buf.Bytes(), 0644)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Created sample_%d.json and sample_%d.bin\n", i, i)
	}
}
