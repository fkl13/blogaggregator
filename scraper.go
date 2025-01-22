package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/fkl13/boot.dev/blogaggregator/internal/database"
	"github.com/google/uuid"
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
		pubDate := sql.NullTime{}
		if t, err := time.Parse(time.RFC822, item.PubDate); err == nil {
			pubDate.Time = t
			pubDate.Valid = true
		}

		params := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Description: item.Description,
			Url:         nextFeed.Url,
			PublishedAt: pubDate,
			FeedID:      nextFeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), params)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Fetched feed %s\n", nextFeed.Name)
}
