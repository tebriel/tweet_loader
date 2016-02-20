package main

import (
	"testing"
)

var test_filename = "./fixtures/example.csv"

var header_row = []string{"tweet_id", "in_reply_to_status_id", "in_reply_to_user_id", "timestamp", "source", "text", "retweeted_status_id", "retweeted_status_user_id", "retweeted_status_timestamp", "expanded_urls"}

var record_no_urls = []string{"700467274776387584", "700450215577255936", "14838543", "2016-02-18 23:49:47 +0000", "<a href=\"\"http://tapbots.com/tweetbot\"\" rel=\"\"nofollow\"\">Tweetbot for iΟS</a>", "@debadair @elastic #golang for me", "", "", "", ""}

var record_with_urls = []string{"700467274776387584", "700450215577255936", "14838543", "2016-02-18 23:49:47 +0000", "<a href=\"\"http://tapbots.com/tweetbot\"\" rel=\"\"nofollow\"\">Tweetbot for iΟS</a>", "@debadair @elastic #golang for me", "", "", "", "http://twitter.com/GoatUserStories/status/700322632114184193/photo/1,http://twitter.com/GoatUserStories/status/700322632114184193/photo/1"}

func TestConvertRecordToTweet(t *testing.T) {
	tweet := ConvertRecordToTweet(record_no_urls)
	expected_id := "700467274776387584"
	actual_id := tweet.TweetId
	expected_timestamp := "2016-02-18 23:49:47 +0000"
	actual_timestamp := tweet.Timestamp
	if expected_id != actual_id {
		t.Errorf("Expected tweet id to be %d but was %d", expected_id, actual_id)
	} else if expected_timestamp != actual_timestamp {
		t.Errorf("Expected tweet timestamp to be %s but was %s", expected_timestamp, actual_timestamp)
	}
}

func TestConvertRecordToTweetUrls(t *testing.T) {
	tweet := ConvertRecordToTweet(record_with_urls)
	expected_length := 2
	actual_length := len(tweet.ExpandedUrls)
	if actual_length != expected_length {
		t.Errorf("Expected %d tweets but had %d", expected_length, actual_length)
	}
}

func TestMarshallRecords(t *testing.T) {
	var records = [][]string{header_row, record_no_urls, record_with_urls}
	tweets := MarshallRecords(records)
	expected_length := 2
	actual_length := len(tweets)
	if actual_length != expected_length {
		t.Errorf("Expected %d tweets, but had %d", expected_length, actual_length)
	}
	expected_id := "700467274776387584"
	actual_id := tweets[0].TweetId
	if actual_id != expected_id {
		t.Errorf("Expected the first tweet to have id of %d but instead had %d", expected_id, actual_id)
	}
}

func TestReadCSV(t *testing.T) {
	records := ReadCSV(test_filename)
	expected_length := 10
	actual_length := len(records)
	if actual_length != expected_length {
		t.Errorf("Expected there to be %d lines but there instead were %d", expected_length, actual_length)
	}
}
