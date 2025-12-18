// Name: Nithya Santhosh
// StudentID: 241ADB038

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// -----------------------------
// Entry point
// -----------------------------

func main() {
	rFlag := flag.Int("r", -1, "generate N random integers (N >= 10)")
	iFlag := flag.String("i", "", "input file with integers (one per line)")
	dFlag := flag.String("d", "", "directory with .txt files to sort")
	flag.Parse()

	hasR := *rFlag != -1
	hasI := *iFlag != ""
	hasD := *dFlag != ""

	count := 0
	if hasR {
		count++
	}
	if hasI {
		count++
	}
	if hasD {
		count++
	}

	if count != 1 {
		log.Fatal("Usage:\n  gosort -r N\n  gosort -i input.txt\n  gosort -d directory")
	}

	var err error
	switch {
	case hasR:
		err = runRandom(*rFlag)
	case hasI:
		err = runInputFile(*iFlag)
	case hasD:
		err = runDirectory(*dFlag)
	}

	if err != nil {
		log.Fatal(err)
	}
}

// -----------------------------
// -r mode logic
// -----------------------------

func runRandom(n int) error {
	if n < 10 {
		return errors.New("N must be >= 10")
	}

	numbers := generateRandomNumbers(n)

	fmt.Println("Original numbers:")
	fmt.Println(numbers)

	chunks := splitIntoChunks(numbers)

	fmt.Println("\nChunks before sorting:")
	printChunks(chunks)

	sortedChunks := sortChunksConcurrently(chunks)

	fmt.Println("\nChunks after sorting:")
	printChunks(sortedChunks)

	result := mergeSortedChunks(sortedChunks)

	fmt.Println("\nFinal sorted result:")
	fmt.Println(result)

	return nil
}

// -----------------------------
// -i mode logic
// -----------------------------

func runInputFile(filename string) error {
	numbers, err := readNumbersFromFile(filename)
	if err != nil {
		return err
	}

	if len(numbers) < 10 {
		return errors.New("input file must contain at least 10 valid integers")
	}

	fmt.Println("Original numbers:")
	fmt.Println(numbers)

	chunks := splitIntoChunks(numbers)

	fmt.Println("\nChunks before sorting:")
	printChunks(chunks)

	sortedChunks := sortChunksConcurrently(chunks)

	fmt.Println("\nChunks after sorting:")
	printChunks(sortedChunks)

	result := mergeSortedChunks(sortedChunks)

	fmt.Println("\nFinal sorted result:")
	fmt.Println(result)

	return nil
}

// -----------------------------
// -d mode logic
// -----------------------------

func runDirectory(dir string) error {
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		return errors.New("invalid directory")
	}

	outDir := fmt.Sprintf("%s_sorted_nithya_santhosh_241ADB038", dir)
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) != ".txt" {
			continue
		}

		inputPath := filepath.Join(dir, f.Name())
		numbers, err := readNumbersFromFile(inputPath)
		if err != nil {
			continue
		}

		chunks := splitIntoChunks(numbers)
		sortedChunks := sortChunksConcurrently(chunks)
		result := mergeSortedChunks(sortedChunks)

		outFile := filepath.Join(outDir, f.Name())
		_ = writeNumbersToFile(outFile, result)
	}

	return nil
}

// -----------------------------
// Chunking logic
// -----------------------------

func splitIntoChunks(numbers []int) [][]int {
	n := len(numbers)

	numChunks := int(math.Ceil(math.Sqrt(float64(n))))
	if numChunks < 4 {
		numChunks = 4
	}

	base := n / numChunks
	extra := n % numChunks

	var chunks [][]int
	start := 0

	for i := 0; i < numChunks; i++ {
		size := base
		if i < extra {
			size++
		}
		end := start + size
		chunks = append(chunks, numbers[start:end])
		start = end
	}

	return chunks
}

// -----------------------------
// Concurrent sorting
// -----------------------------

func sortChunksConcurrently(chunks [][]int) [][]int {
	var wg sync.WaitGroup
	result := make([][]int, len(chunks))

	for i, chunk := range chunks {
		wg.Add(1)
		go func(idx int, data []int) {
			defer wg.Done()
			tmp := make([]int, len(data))
			copy(tmp, data)
			sort.Ints(tmp)
			result[idx] = tmp
		}(i, chunk)
	}

	wg.Wait()
	return result
}

// -----------------------------
// Merge logic
// -----------------------------

func mergeSortedChunks(chunks [][]int) []int {
	if len(chunks) == 0 {
		return []int{}
	}
	if len(chunks) == 1 {
		return chunks[0]
	}

	for len(chunks) > 1 {
		var merged [][]int
		for i := 0; i < len(chunks); i += 2 {
			if i+1 < len(chunks) {
				merged = append(merged, mergeTwo(chunks[i], chunks[i+1]))
			} else {
				merged = append(merged, chunks[i])
			}
		}
		chunks = merged
	}
	return chunks[0]
}

func mergeTwo(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)
	return result
}

// -----------------------------
// Helpers
// -----------------------------

func generateRandomNumbers(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := range nums {
		nums[i] = rand.Intn(1000) // range: 0–999
	}
	return nums
}

func readNumbersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		val, err := strconv.Atoi(line)
		if err != nil {
			return nil, errors.New("invalid integer in file")
		}
		nums = append(nums, val)
	}

	return nums, nil
}

func writeNumbersToFile(filename string, nums []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, n := range nums {
		fmt.Fprintln(writer, n)
	}
	return writer.Flush()
}

func printChunks(chunks [][]int) {
	for i, c := range chunks {
		fmt.Printf("Chunk %d: %v\n", i, c)
	}
}