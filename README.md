# Tweet Loader #

If you export all your tweets from the settings page in Twitter, maybe you then want to load them
into ElasticSearch. I don't know if you want to do this, what am I, a wizard?

But if you do, then you can!

## How To ##

1. `export ES_URL=http://your.es.instance:9200 FILENAME=./path/to/tweets.csv`
1. `make run`

## What Do ##

This will create an index called "csvtweets" at the url specified, set up the mapping, then jam all
those tweets into the index. Right now it does it in one huge bulk operation (TODO: Let's batch the
bulk inserts).

## Requirements ##

1. Exported tweets.csv from twitter.com. Example lines look like this:

```csv
"tweet_id","in_reply_to_status_id","in_reply_to_user_id","timestamp","source","text","retweeted_status_id","retweeted_status_user_id","retweeted_status_timestamp","expanded_urls"
"700807342384746496","","15782607","2016-02-19 22:21:05 +0000","<a href=""http://tapbots.com/tweetbot"" rel=""nofollow"">Tweetbot for iΟS</a>","@jordansissel you were busy so I didn’t want to interrupt. But wanted to again say thanks for all your community work with #Logstash","","","",""
"700806305519251457","700802602011742209","16144388","2016-02-19 22:16:58 +0000","<a href=""http://tapbots.com/tweetbot"" rel=""nofollow"">Tweetbot for iΟS</a>","@shapr so it’s a problem, on some level that there isn’t huge adoption, so it’s hard to gain adoption. 5/5","","","",""
"700806106050727936","700802602011742209","16144388","2016-02-19 22:16:10 +0000","<a href=""http://tapbots.com/tweetbot"" rel=""nofollow"">Tweetbot for iΟS</a>","@shapr it’s not unmanageable, but it has a much higher overall barrier to entry than, say, go which is c++ like. ?/?","","","",""
```

Note I don't check to make sure the header is there, I just skip the first line. So if you're
manually creating this for some reason (sure, it's your time you're wasting) add that header.

2. Elasticsearch Instance. That seems pretty self explanatory, www.elastic.co if not.
3. I dunno, desire to play around with my weird script.
4. Golang environment, definitely a golang environment.
