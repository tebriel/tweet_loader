package main

import (
	"encoding/csv"
	"github.com/golang/glog"
	"os"
	"strings"
)

// CSVTweet is the structure of a tweet stored in the CSV File
type CSVTweet struct {
	TweetID                  string   `json:"tweet_id"`
	InReplyToStatusID        string   `json:"in_reply_to_status_id"`
	InReplyToUserID          string   `json:"in_reply_to_user_id"`
	Timestamp                string   `json:"timestamp"`
	Source                   string   `json:"source"`
	Text                     string   `json:"text"`
	RetweetedStatusID        string   `json:"retweeted_status_id"`
	RetweetedStatusUserID    string   `json:"retweeted_status_user_id"`
	RetweetedStatusTimestamp string   `json:"retweeted_status_timestamp"`
	ExpandedUrls             []string `json:"expanded_urls"`
}

// ReadCSV reads a csv from a filename and returns a 2D array of strings
func ReadCSV(filename string) [][]string {
	glog.V(2).Infof("Opening %s", filename)
	csvFile, err := os.Open(filename)
	if err != nil {
		glog.Fatalf("Couldn't read file %s because of %s", filename, err)
	}
	csvReader := csv.NewReader(csvFile)
	glog.V(2).Infof("Reading %s", filename)
	records, err := csvReader.ReadAll()
	if err != nil {
		glog.Fatalf("Couldn't read records: %s", err)
	}
	glog.V(2).Infof("Found %d lines in the file", len(records))
	return records
}

// MarshallRecords Converts the 2D array of CSV strings to CSVTweet structs
func MarshallRecords(records [][]string) []CSVTweet {
	result := make([]CSVTweet, len(records)-1)
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		result[idx-1] = ConvertRecordToTweet(record)
	}
	return result
}

// ConvertRecordToTweet takes an individual CSV Record and converts it to a CSVTweet
func ConvertRecordToTweet(r []string) CSVTweet {
	var result CSVTweet
	result.TweetID = r[0]
	result.InReplyToStatusID = r[1]
	result.InReplyToUserID = r[2]
	result.Timestamp = r[3]
	result.Source = r[4]
	result.Text = r[5]
	result.RetweetedStatusID = r[6]
	result.RetweetedStatusUserID = r[7]
	result.RetweetedStatusTimestamp = r[8]
	if r[9] != "" {
		result.ExpandedUrls = strings.Split(r[9], ",")
	}
	return result
}
