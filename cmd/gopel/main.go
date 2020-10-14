package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// func lineInput(prompt string) string {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print(prompt)
// 	// fmt.Printf("%s ", prompt)
// 	input, err := reader.ReadString('\n')
// 	if err != nil {
// 		// TODO: return err instead
// 		log.Fatal(err)
// 	}

// 	// err = f(strings.TrimSpace(input))
// 	// if err != nil {
// 	//         fmt.Println(err)
// 	// }

// 	return strings.TrimSpace(input)
// }

func lineInput(prompt string) rune {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	// fmt.Printf("%s ", prompt)
	input, _, err := reader.ReadRune()
	if err != nil {
		// TODO: return err instead
		log.Fatal(err)
	}

	// err = f(strings.TrimSpace(input))
	// if err != nil {
	//         fmt.Println(err)
	// }

	return input
}

func main() {
	infile := flag.String("infile", "", "input file")
	outfile := flag.String("outfile", "", "output file")
	flag.Parse()

	if *infile == "" || *outfile == "" {
		log.Fatalln("Please specify input and output files")
	}

	bb, err := ioutil.ReadFile(*infile)
	if err != nil {
		log.Fatalln(err)
	}

	var items []string
	if err := json.Unmarshal(bb, &items); err != nil {
		log.Fatalln(err)
	}

	sort.Slice(items, func(i int, j int) bool {
		for {
			fmt.Println()
			fmt.Println("Which is better?")
			fmt.Printf("1 -> %s\n", items[i])
			fmt.Printf("2 -> %s\n", items[j])
			r := lineInput("? ")
			switch r {
			case '1':
				return true
			case '2':
				return false
			default:
				fmt.Println("Please enter a number.")
			}
		}
	})

	bb2, err := json.Marshal(items)
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile(*outfile, bb2, 0644); err != nil {
		log.Fatalln(err)
	}
}
