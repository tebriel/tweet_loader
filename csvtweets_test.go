package main

import (
	"testing"
)

var testFilename = "./fixtures/example.csv"

var headerRow = []string{"tweet_id", "in_reply_to_status_id", "in_reply_to_user_id", "timestamp", "source", "text", "retweeted_status_id", "retweeted_status_user_id", "retweeted_status_timestamp", "expanded_urls"}

var recordNoUrls = []string{"700467274776387584", "700450215577255936", "14838543", "2016-02-18 23:49:47 +0000", "<a href=\"\"http://tapbots.com/tweetbot\"\" rel=\"\"nofollow\"\">Tweetbot for iΟS</a>", "@debadair @elastic #golang for me", "", "", "", ""}

var recordWithUrls = []string{"700467274776387584", "700450215577255936", "14838543", "2016-02-18 23:49:47 +0000", "<a href=\"\"http://tapbots.com/tweetbot\"\" rel=\"\"nofollow\"\">Tweetbot for iΟS</a>", "@debadair @elastic #golang for me", "", "", "", "http://twitter.com/GoatUserStories/status/700322632114184193/photo/1,http://twitter.com/GoatUserStories/status/700322632114184193/photo/1"}

func TestConvertRecordToTweet(t *testing.T) {
	tweet := ConvertRecordToTweet(recordNoUrls)
	expectedID := "700467274776387584"
	actualID := tweet.TweetID
	expectedTimestamp := "2016-02-18 23:49:47 +0000"
	actualTimestamp := tweet.Timestamp
	if expectedID != actualID {
		t.Errorf("Expected tweet id to be %s but was %s", expectedID, actualID)
	} else if expectedTimestamp != actualTimestamp {
		t.Errorf("Expected tweet timestamp to be %s but was %s", expectedTimestamp, actualTimestamp)
	}
}

func TestConvertRecordToTweetUrls(t *testing.T) {
	tweet := ConvertRecordToTweet(recordWithUrls)
	expectedLength := 2
	actualLength := len(tweet.ExpandedUrls)
	if actualLength != expectedLength {
		t.Errorf("Expected %d tweets but had %d", expectedLength, actualLength)
	}
}

func TestMarshallRecords(t *testing.T) {
	var records = [][]string{headerRow, recordNoUrls, recordWithUrls}
	tweets := MarshallRecords(records)
	expectedLength := 2
	actualLength := len(tweets)
	if actualLength != expectedLength {
		t.Errorf("Expected %d tweets, but had %d", expectedLength, actualLength)
	}
	expectedID := "700467274776387584"
	actualID := tweets[0].TweetID
	if actualID != expectedID {
		t.Errorf("Expected the first tweet to have id of %s but instead had %s", expectedID, actualID)
	}
}

func TestReadCSV(t *testing.T) {
	records := ReadCSV(testFilename)
	expectedLength := 10
	actualLength := len(records)
	if actualLength != expectedLength {
		t.Errorf("Expected there to be %d lines but there instead were %d", expectedLength, actualLength)
	}
}
