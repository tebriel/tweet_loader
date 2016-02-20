package main

import (
	"github.com/golang/glog"
	"gopkg.in/olivere/elastic.v3"
	"time"
)

var index_name = "csvtweets"

func connect(url string) *elastic.Client {
	// Create a client
	glog.V(2).Infof("Connecting to ES at: %s", url)
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		glog.Fatalf("Couldn't connect to ES: %s", err)
	}

	return client
}

func MakeIndex(client *elastic.Client) {
	glog.V(2).Infof("Checking to see if index %s exists", index_name)
	exists, err := client.IndexExists(index_name).Do()
	if err != nil {
		glog.Fatalf("Had problems checking index existence: %s", err)
	}
	if !exists {
		glog.V(2).Info("Index didn't exist, creating it")
		// Index does not exist yet.
		client.CreateIndex(index_name).Do()
		glog.V(2).Info("Pushing mapping to Index")
		resp, err := client.PutMapping().Index(index_name).Type("tweet").BodyString(es_mapping).Do()
		if err != nil {
			glog.Fatalf("Couldn't push mapping to index: %s", err)
		} else {
			glog.V(2).Infof("PutMapping Response: %t", resp)
		}
	} else {
		glog.V(2).Info("Index existed, nothing to do here")
	}
}

func SendToES(client *elastic.Client, tweets []CSVTweet) {
	glog.V(2).Infof("Building Bulk Insert for %d Tweets", len(tweets))

	bulkRequest := client.Bulk()
	for idx, tweet := range tweets {
		if idx == 0 {
			continue
		}
		action := elastic.NewBulkIndexRequest().Index(index_name).Type("tweet").Id(tweet.TweetId).Doc(tweet)
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

func DoESWork(es_url string, tweets []CSVTweet) {
	client := connect(es_url)
	MakeIndex(client)
	SendToES(client, tweets)
}
