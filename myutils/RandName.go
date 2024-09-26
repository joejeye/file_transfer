package myutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func RandName() string {
	pathToNames := GetRootDir() + "/assets/list_of_names.txt"
	file, err := os.Open(pathToNames)
	if err != nil {
		panic(fmt.Errorf("error opening file: %w", err))
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Errorf("error closing file: %w", err))
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(fmt.Errorf("error scanning file: %w", err))
	}

	if len(lines) == 0 {
		panic("no names in file")
	}

	// Create a new random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly select a name from the list
	return lines[rng.Intn(len(lines))]
}
