package main

import (
	"github.com/golang/glog"
	"gopkg.in/olivere/elastic.v3"
	"time"
)

var indexName = "csvtweets"

func connect(url string) *elastic.Client {
	// Create a client
	glog.V(2).Infof("Connecting to ES at: %s", url)
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		glog.Fatalf("Couldn't connect to ES: %s", err)
	}

	return client
}

// MakeIndex creates the needed elasticsearch Index
func MakeIndex(client *elastic.Client) {
	glog.V(2).Infof("Checking to see if index %s exists", indexName)
	exists, err := client.IndexExists(indexName).Do()
	if err != nil {
		glog.Fatalf("Had problems checking index existence: %s", err)
	}
	if !exists {
		glog.V(2).Info("Index didn't exist, creating it")
		// Index does not exist yet.
		client.CreateIndex(indexName).Do()
		glog.V(2).Info("Pushing mapping to Index")
		resp, err := client.PutMapping().Index(indexName).Type("tweet").BodyString(esMapping).Do()
		if err != nil {
			glog.Fatalf("Couldn't push mapping to index: %s", err)
		} else {
			glog.V(2).Infof("PutMapping Response: %t", resp)
		}
	} else {
		glog.V(2).Info("Index existed, nothing to do here")
	}
}

// SendToES does a bulk index of tweets into ElasticSearch
func SendToES(client *elastic.Client, tweets []CSVTweet) {
	glog.V(2).Infof("Building Bulk Index for %d Tweets", len(tweets))

	bulkRequest := client.Bulk()
	for idx, tweet := range tweets {
		if idx == 0 {
			continue
		}
		action := elastic.NewBulkIndexRequest().Index(indexName).Type("tweet").Id(tweet.TweetID).Doc(tweet)
		bulkRequest.Add(action)
	}

	t := time.Now()
	glog.V(2).Info("Bulk Inserting tweets into ES")
	_, err := bulkRequest.Do()
	if err != nil {
		glog.Fatalf("Couldn't bulk load documents: %s", err)
	}
	glog.V(2).Infof("Bulk Insert Complete, taking %ds", time.Since(t)/time.Second)

}

// DoESWork sticthes together all of the ES functionality into a single workflow
func DoESWork(esURL string, tweets []CSVTweet) {
	client := connect(esURL)
	MakeIndex(client)
	SendToES(client, tweets)
}
