package main

import (
	"encoding/csv"
	"flag"
	"github.com/golang/glog"
	"gopkg.in/olivere/elastic.v3"
	"os"
	"strconv"
	"time"
)

var field_mapping = []string{"TweetId", "InReplyToStatusId", "InReplyToUserId", "Timestamp", "Source", "Text", "RetweetedStatusId", "RetweetedStatusUserId", "RetweetedStatusTimestamp", "ExpandedUrls"}

type CSVTweet struct {
	TweetId                  int    `json:"tweet_id"`
	InReplyToStatusId        int    `json:"in_reply_to_status_id"`
	InReplyToUserId          int    `json:"in_reply_to_user_id"`
	Timestamp                string `json:"timestamp"`
	Source                   string `json:"source"`
	Text                     string `json:"text"`
	RetweetedStatusId        int    `json:"retweeted_status_id"`
	RetweetedStatusUserId    int    `json:"retweeted_status_user_id"`
	RetweetedStatusTimestamp string `json:"retweeted_status_timestamp"`
	ExpandedUrls             string `json:"expanded_urls"`
}

func ReadCSV(filename string) [][]string {
	csv_file, err := os.Open(filename)
	if err != nil {
		glog.Fatalf("Couldn't read file %s because of %s", filename, err)
	}
	csvreader := csv.NewReader(csv_file)
	records, err := csvreader.ReadAll()
	if err != nil {
		glog.Fatalf("Couldn't read records: %s", err)
	}
	glog.Infof("Record 1: %s, %d fields", records[1], len(records[1]))
	return records
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
	result.ExpandedUrls = r[9]
	return result
}

func connect(url string) *elastic.Client {
	// Create a client
	glog.V(2).Infof("Connecting to ES at: %s", url)
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		glog.Fatalf("Couldn't connect to ES: %s", err)
	}

	return client
}

func SendToES(es_url string, tweets [][]string) {
	client := connect(es_url)
	glog.V(2).Infof("Building Bulk Insert for %d Tweets", len(tweets))

	bulkRequest := client.Bulk()
	for idx, tweet := range tweets {
		if idx == 0 {
			continue
		}
		tweet_r := ConvertRecordToTweet(tweet)
		action := elastic.NewBulkIndexRequest().Index("csvtweets").Type("tweet").Id(strconv.Itoa(tweet_r.TweetId)).Doc(tweet_r)
		bulkRequest.Add(action)
	}

	t := time.Now()
	glog.V(2).Info("Inserting bulk tweets into ES")
	_, err := bulkRequest.Do()
	if err != nil {
		glog.Fatalf("Couldn't bulk load documents: %s", err)
	}
	glog.V(2).Infof("Bulk Insert Complete, taking %ds", time.Since(t)/time.Second)

}

func main() {
	flag.Parse()
	records := ReadCSV("tweets.csv")
	SendToES("http://frodux.in:9200", records)

	glog.Infof("Record: %s", ConvertRecordToTweet(records[1]).Text)
}
