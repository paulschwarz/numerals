package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sequence = "IVXLCDM"

func main() {
	args := os.Args[1:]
	result := ""
	val := ""

	if len(args) < 1 {
		fmt.Println("Usage: numerals [romanize|arabize] value")
		os.Exit(1)
	} else if len(args) < 2 {
		info, _ := os.Stdin.Stat()

		if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage: cat yourfile.txt | numerals [romanize|arabize]")
		} else if info.Size() > 0 {
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Failed to read input")
				os.Exit(1)
			}

			val = strings.TrimRight(input, "\n")
		}
	} else {
		val = args[1]
	}

	switch args[0] {
	case "romanize":
		val, err := strconv.Atoi(val)

		if err != nil {
			fmt.Println("Usage: numerals romanize integers")
			os.Exit(1)
		}

		result = Roman(val)
		break
	case "arabize":
		for _, s := range strings.Split(val, "") {
			if !strings.Contains(sequence, s) {
				fmt.Println("Usage: numerals arabize [" + sequence + "]")
				os.Exit(1)
			}
		}

		result = strconv.Itoa(Arab(val))
		break
	}

	fmt.Println(result)
}

// Roman converts an arabic numeral to a roman numeral
func Roman(val int) string {

	type step struct {
		inc  int
		full string
		half string
		unit string
	}

	steps := []step{
		{1000, "M", "D", "C"},
		{100, "C", "L", "X"},
		{10, "X", "V", "I"},
	}

	result := []string{}

	for _, step := range steps {
		for val >= int(float64(step.inc)*.9) {
			val -= step.inc
			result = append(result, step.full)
		}
		for val >= int(float64(step.inc)*.4) {
			val -= int(float64(step.inc) * .5)
			result = append(result, step.half)
		}
		if val < 0 {
			start := result[0 : len(result)-1]
			end := result[len(result)-1]
			for val < 0 {
				val += int(float64(step.inc) * .1)
				result = append(start, step.unit)
			}
			result = append(result, end)
		}
	}

	for val >= 1 {
		val--
		result = append(result, "I")
	}

	return strings.Join(result, "")
}

// Arab converts a roman numeral to an arabic numeral
func Arab(val string) int {

	numerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	seq := strings.Split(val, "")
	len := len(seq)
	i := 0
	result := 0

	for i < len {
		if i+1 < len && greater(seq[i+1], seq[i]) {
			result += numerals[seq[i+1]] - numerals[seq[i]]
			i += 2
		} else {
			result += numerals[seq[i]]
			i++
		}
	}

	return result
}

func greater(val1, val2 string) bool {
	return strings.Index(sequence, val1) > strings.Index(sequence, val2)
}
