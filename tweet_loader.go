package main

import (
	"flag"
)

func getArgs() (string, string) {
	filename := flag.String("f", "tweets.csv", "tweets.csv File (with header)")
	esURLFlag := flag.String("esURL", "http://frodux.in:9200", "URL for elasticsearch instance")

	flag.Parse()

	return *filename, *esURLFlag
}

func main() {
	filename, esURL := getArgs()
	tweetRows := ReadCSV(filename)
	records := MarshallRecords(tweetRows)
	DoESWork(esURL, records)
}
