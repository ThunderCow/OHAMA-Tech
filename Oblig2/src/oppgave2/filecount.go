/* The code lines 114-119 & 131-136 are not completed yet(*to take out most common runes from a file & look values
from highest to lowest )but major part has been done.*/
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"os"
	"io/ioutil"
	"bufio"
)
//Linecounter for the file given as a argument
func linjeteller(filename string) (int64, error) {
	lc := int64(0)
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	//to close a file
	defer f.Close()
	//scaning for line counter
	s := bufio.NewScanner(f)
	for s.Scan() {
		lc++
	}
	return lc, s.Err()
}
//to find number of words in a file
func word_count(words []string) map[string]int {
	// iterate over the slice of words and populate
	// the map with word:frequency pairs
	// (where word is the key and frequency is the value)
	word_freq := make(map[string]int)
	// range string slice gives index, word pairs
	// index is not needed, so use blank identifier _
	for _, word := range words {
		// check if word (the key) is in already in the map
		_, ok := word_freq[word]
		// if true add 1 to frequency (value of map)
		// else start frequency at 1
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}
//declaring type structures
type word_struct struct {
	freq int
	word string
}

// word_struct will be displayed in this format
func (p word_struct) String() string {
	return fmt.Sprintf("%3d   %s", p.freq, p.word)
}

// by_freq implements sort.Interface for []word_struct based on the freq field
type by_freq []word_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq < a[j].freq }

// by_word implements sort.Interface for []word_struct based on the word field
type by_word []word_struct

func (a by_word) Len() int           { return len(a) }
func (a by_word) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_word) Less(i, j int) bool { return a[i].word < a[j].word }

func main() {
	//to take argument
	filename := os.Args[1]
	lc, err := linjeteller(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	//print out file info
	fmt.Println("Information about", filename+":\n",)
	fmt.Println("Number of lines in file:", lc)
	
	//code to present data in rune form and find most common runes
	filen := os.Args[1]
	read, err := ioutil.ReadFile(filen)
	if err != nil {
		fmt.Print(err)
	}
	text := string(read)
	// change text to all lower characters
	text = strings.ToLower(text)

	// match whole words, removes any punctuation/whitespace
	re := regexp.MustCompile(`[^\s-]`)
	words := re.FindAllString(text, -1)

	// word_map() returns a map of word:frequency pairs
	word_map := word_count(words)

	// convert the map to a slice of structures for sorting
	// create pointer to an instance of word_struct
	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	ix := 0
	maxK := ""
	maxV := 0
	for key, val := range word_map {
		pws.freq = val
		pws.word = key
		if pws.freq > maxV {
			maxV = pws.freq
			//temp := maxV
			//maxV = 0
			maxK = pws.word
		}
		struct_slice[ix] = *pws
		ix++
	}

	// sorting slice of structers by field freq in place
	sort.Sort(by_freq(struct_slice))
	for r := 0; r < len(struct_slice); r++ {
		fmt.Println("Rune:", struct_slice[r].word, "Count:", struct_slice[r].freq)

	}
	//Printing most common runes in a file
	fmt.Println("Most common runes are:\n")
	for v := 0; v < 5; v++ {
		fmt.Println("Rune:", maxK, "Count:", maxV)
		maxV--
		//strings.Trim(pws.word, maxK)
	}
}
