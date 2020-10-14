// Copyright 2020 Andrew Merenbach
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/merenbach/gopel/internal/fileutil"
)

func runeInput(prompt string) (rune, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _, err := reader.ReadRune()
	if err != nil {
		return (-1), err
	}
	return input, nil
}

func main() {
	infile := flag.String("infile", "", "input file")
	outfile := flag.String("outfile", "", "output file")
	flag.Parse()

	if *infile == "" || *outfile == "" {
		log.Fatalln("Please specify input and output files")
	}

	var items []string
	if err := fileutil.ReadJSON(*infile, &items); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("A series of interactive prompts will follow.")
	sort.Slice(items, func(i int, j int) bool {
		for {
			fmt.Println()
			fmt.Println("Which is better?")
			fmt.Printf("1 -> %s\n", items[i])
			fmt.Printf("2 -> %s\n", items[j])
			r, err := runeInput("? ")
			if err != nil {
				log.Println("Unrecognized input")
				continue
			}
			switch r {
			case '1':
				return true
			case '2':
				return false
			default:
				fmt.Println("Please enter 1 or 2")
			}
		}
	})
	fmt.Println("All prompts have been answered.")

	if err := fileutil.WriteJSON(*outfile, items); err != nil {
		log.Fatalln(err)
	}
}
