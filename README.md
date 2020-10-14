# Gopel

_Gopel_ is a transliteration of the Yiddish word for "fork." This program presents a series of decision forks to the user to help prioritize a list of strings.

This tool assumes that the items under evaluation hold transitive relative value. In practice this means that if you like pizza more than pasta, pasta more than pancakes, and pancakes more than pizza, the output may not reflect a satisfactory or consistent prioritization.

## Instructions

Run as follows:

    go run cmd/gopel/main.go -infile sample.json -outfile sample-out.json
