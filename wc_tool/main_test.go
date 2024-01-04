package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

var testData = `The sentence quoted above from VI. § 21 hardly strikes me as one that
could have been written in the full flush of victory.512	Ho Lu attacks Ch’u, but is dissuaded from entering Ying,
the capital. Shih Chi mentions Sun Wu as general.`

func TestWcTestData(t *testing.T) {
	rd := strings.NewReader(testData)
	nLines, nWords, nChars, nBytes, err := wc(rd)
	checkEqual(t, nil, err)
	checkEqual(t, uint64(2), nLines)
	checkEqual(t, uint64(43), nWords)
	checkEqual(t, uint64(233), nChars)
	checkEqual(t, uint64(236), nBytes)
}

func TestWcFile(t *testing.T) {
	f, err := os.Open("test.txt")
	checkEqual(t, nil, err)
	defer f.Close()
	nLines, nWords, nChars, nBytes, err := wc(f)
	checkEqual(t, nil, err)
	checkEqual(t, uint64(7143), nLines)
	checkEqual(t, uint64(58164), nWords)
	checkEqual(t, uint64(332144), nChars)
	checkEqual(t, uint64(335040), nBytes)
}

func checkEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	t.Errorf("Not equal: \n" + 
		"expected: %T(%v)\n" +
		"actual  : %T(%v)", expected, expected, actual, actual)
}

func BenchmarkWc(b *testing.B) {
	rd := strings.NewReader(testData)
	b.ResetTimer() // If a benchmark needs some expensive setup before running, the timer may be reset
	for i := 0; i < b.N; i++ {
        wc(rd)
    }
}