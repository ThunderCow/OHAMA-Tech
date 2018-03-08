package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"io/ioutil"
	"strings"
	"regexp"
)

//Linecounter for the file given as a argument
func linjeteller(filename string) (int64, error) {
	lc := int64(0)
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		lc++
	}
	return lc, s.Err()
}

/*func bokstavteller() {
	filen := os.Args[1]
	b, err := ioutil.ReadFile(filen)
	if err != nil {
		fmt.Print(err)
	}

	asciiStr := b
	asciiBytes := []byte(asciiStr)

	fmt.Printf("test string=%v, bytes=%v\n", asciiStr, asciiBytes)
}*/

func char_count(chars []string) map[string]int {
	// iterate over the slice of charcters and populate
	// the map with char:frequency pairs
	// (where char is the key and frequency is the value)
	char_freq := make(map[string]int)
	// range string slice gives index, char pairs
	// index is not needed, so use blank identifier _
	for _, char := range chars {
		// check if char (the key) is in already in the map
		_, ok := char_freq[char]
		// if true add 1 to frequency (value of map)
		// else start frequency at 1
		if ok == true {
			char_freq[char] += 1
		} else {
			char_freq[char] = 1
		}
	}
	return char_freq
}

type char_struct struct {
	freq int
	char string
}

// by_freq implements sort.Interface for []char_struct based on the freq field
type by_freq []char_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq > a[j].freq }

// by_char implements sort.Interface for []char_struct based on the word field
type by_char []char_struct

func (a by_char) Len() int           { return len(a) }
func (a by_char) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_char) Less(i, j int) bool { return a[i].char > a[j].char }

func main() {
	filen := os.Args[1]
	read, err := ioutil.ReadFile(filen)
	lc, err := linjeteller(filen)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Information about", filen+":\n",)
	fmt.Println("Number of lines in file:", lc)

	text := string(read)

	// change text to all lower characters
	text = strings.ToLower(text)

	// match chars, removes any punctuation/whitespace
	re := regexp.MustCompile(`[^a\s-]`)
	chars := re.FindAllString(text, -1)

	// char_map() returns a map of char:frequency pairs
	char_map := char_count(chars)

	// convert the map to a slice of structures for sorting
	// create pointer to an instance of char_struct
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
	// sorting slice of structers by field freq in place
	sort.Sort(by_freq(struct_slice))

	for i, ix := 1, 0; i < 6; i, ix = i+1, ix+1 {
		fmt.Println(i, "Rune:", struct_slice[ix].char, " Count:", struct_slice[ix].freq)
	}
}
