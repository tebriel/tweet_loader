package main

// I thought about having this as a template, but it doesn't change. But I wanted it compiled into
// the app. If this is bad (yeah definitely probably is) oh well, too bad
const esMapping = `
{
  "tweet": {
    "properties": {
      "expanded_urls": {
        "type": "string",
        "analyzer": "whitespace"
      },
      "in_reply_to_status_id": {
        "type": "string"
      },
      "in_reply_to_user_id": {
        "type": "string"
      },
      "retweeted_status_id": {
        "type": "string"
      },
      "retweeted_status_timestamp": {
        "type": "date",
        "ignore_malformed": true,
        "format": "YYYY-MM-dd HH:mm:ss Z"
      },
      "retweeted_status_user_id": {
        "type": "string"
      },
      "source": {
        "type": "string"
      },
      "text": {
        "type": "string"
      },
      "timestamp": {
        "type": "date",
        "format": "YYYY-MM-dd HH:mm:ss Z",
        "ignore_malformed": true
      },
      "tweet_id": {
        "type": "string"
      }
    }
  }
}
`
