package main

import (
	"flag"
)

func get_args() (string, string) {
	filename := flag.String("f", "tweets.csv", "tweets.csv File (with header)")
	es_url_flag := flag.String("es_url", "http://frodux.in:9200", "URL for elasticsearch instance")

	flag.Parse()

	return *filename, *es_url_flag
}

func main() {
	filename, es_url := get_args()
	tweet_rows := ReadCSV(filename)
	records := MarshallRecords(tweet_rows)
	DoESWork(es_url, records)
}
