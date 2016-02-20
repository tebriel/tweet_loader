package main

import (
	"encoding/csv"
	"github.com/golang/glog"
	"os"
	"strconv"
	"strings"
)

type CSVTweet struct {
	TweetId                  int      `json:"tweet_id"`
	InReplyToStatusId        int      `json:"in_reply_to_status_id"`
	InReplyToUserId          int      `json:"in_reply_to_user_id"`
	Timestamp                string   `json:"timestamp"`
	Source                   string   `json:"source"`
	Text                     string   `json:"text"`
	RetweetedStatusId        int      `json:"retweeted_status_id"`
	RetweetedStatusUserId    int      `json:"retweeted_status_user_id"`
	RetweetedStatusTimestamp string   `json:"retweeted_status_timestamp"`
	ExpandedUrls             []string `json:"expanded_urls"`
}

func ReadCSV(filename string) [][]string {
	glog.V(2).Infof("Opening %s", filename)
	csv_file, err := os.Open(filename)
	if err != nil {
		glog.Fatalf("Couldn't read file %s because of %s", filename, err)
	}
	csvreader := csv.NewReader(csv_file)
	glog.V(2).Infof("Reading %s", filename)
	records, err := csvreader.ReadAll()
	if err != nil {
		glog.Fatalf("Couldn't read records: %s", err)
	}
	glog.V(2).Infof("Found %d lines in the file", len(records))
	return records
}

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

func ConvertRecordToTweet(r []string) CSVTweet {
	var result CSVTweet
	result.TweetId, _ = strconv.Atoi(r[0])
	result.InReplyToStatusId, _ = strconv.Atoi(r[1])
	result.InReplyToUserId, _ = strconv.Atoi(r[2])
	result.Timestamp = r[3]
	result.Source = r[4]
	result.Text = r[5]
	result.RetweetedStatusId, _ = strconv.Atoi(r[6])
	result.RetweetedStatusUserId, _ = strconv.Atoi(r[7])
	result.RetweetedStatusTimestamp = r[8]
	if r[9] != "" {
		result.ExpandedUrls = strings.Split(r[9], ",")
	}
	return result
}
