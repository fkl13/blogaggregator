package main

import (
	"context"
	"fmt"
	"log"
)

func scrapeFeeds(s *state) {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("couldn't get next feed", err)
		return
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		log.Printf("couldn't mark feed %s as fetched: %v", nextFeed.Name, err)
		return
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		log.Printf("couldn't fetch %s feed: %v", nextFeed.Name, err)
		return
	}

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}
}
