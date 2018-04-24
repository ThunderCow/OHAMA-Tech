package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

// Linecounter for the file given as a argument
func linecounter(filename string) (int64, error) {
	lc := int64(0)
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// NewScanner returns a new scanner to read from f (file givens as argument)
	// The NewScanner does a split function with defaults to ScanLines.
	// The for loop counts given total line the the file.
	s := bufio.NewScanner(f)
	for s.Scan() {
		lc++
	}
	return lc, s.Err()
}


func char_count(chars []string) map[string]int {
	// Iterate over the slice of characters and populate
	// The map with char:frequency pairs
	// (where char is the key and frequency is the value)
	char_freq := make(map[string]int)
	// range string slice gives index, char pairs
	// Index is not needed, so use blank identifier _
	for _, char := range chars {
		// Check if char (the key) is in already in the map
		_, ok := char_freq[char]
		// If true add 1 to frequency (value of map)
		// Else start frequency at 1
		if ok == true {
			char_freq[char] += 1
		} else {
			char_freq[char] = 1
		}
	}
	return char_freq
}

// Define what data type we will use
type char_struct struct {
	freq int
	char string
}

// by_freq implements sort.Interface for []char_struct based on the freq field
type by_freq []char_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq > a[j].freq }


func main() {
	// Get a file as argument from terminal, reads the content
	filen := os.Args[1]
	read, err := ioutil.ReadFile(filen)
	lc, err := linecounter(filen)
	if err != nil {
		fmt.Print(err)
	}

	// Print the filename and linecount
	fmt.Println("Information about", filen+":\n",)
	fmt.Println("Number of lines in file:", lc, "\n")

	// Defines the text for char_count function, read from the file given in terminal as argument
	text := string(read)

	// change text to all lower characters
	text = strings.ToLower(text)

	// Match chars, removes any punctuation/whitespace
	re := regexp.MustCompile(`[^a\s-]`)
	chars := re.FindAllString(text, -1)

	// char_map() returns a map of char:frequency pairs
	char_map := char_count(chars)

	// Convert the map to a slice of structures for sorting
	// Create pointer to an instance of char_struct
	pws := new(char_struct)
	struct_slice := make([]char_struct, len(char_map))
	ix := 0
	for key, val := range char_map {
		pws.freq = val
		pws.char = key
		struct_slice[ix] = *pws
		ix++
	}
	fmt.Println("Most common runes:\n")

	// Sorting slice of structers by field freq in place
	sort.Sort(by_freq(struct_slice))
	for i, ix := 1, 0; i < 6; i, ix = i+1, ix+1 {
		randomString := strconv.Itoa(i) + "."
		fmt.Println(strings.Replace(randomString, " ", "", -1),  "Rune:", struct_slice[ix].char, ", Count:", struct_slice[ix].freq)
		}
}
