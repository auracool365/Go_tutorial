package main

import (
	"fmt"
	"sync"
	"time"
)

type Article struct {
	Title   string
	Content string
}

type ArticleStore struct {
	articles map[string]Article
	rwMutex  sync.RWMutex
}

func (as *ArticleStore) ReadArticle(title string) (Article, bool) {
	as.rwMutex.RLock()
	defer as.rwMutex.RUnlock()

	article, exists := as.articles[title]
	return article, exists
}

func (as *ArticleStore) WriteArticle(title, content string) {
	as.rwMutex.Lock()
	defer as.rwMutex.Unlock()

	// Simulate processing time for writing/updating an article
	time.Sleep(time.Millisecond * 100)

	as.articles[title] = Article{Title: title, Content: content}
	fmt.Printf("Article '%s' updated.\n", title)
}

func main() {
	articleSystem := ArticleStore{
		articles: make(map[string]Article),
	}

	// Concurrent readers
	for i := range 3 {
		go func(userNum int) {
			title := "Introduction to Go"
			article, exists := articleSystem.ReadArticle(title)
			if exists {
				fmt.Printf("User %d is reading '%s': %s\n", userNum, title, article.Content)
			} else {
				fmt.Printf("User %d couldn't find article '%s'.\n", userNum, title)
			}
		}(i + 1)
	}

	// Exclusive writer
	time.Sleep(50 * time.Millisecond) // Give readers a head start
	go func() {
		title := "Introduction to Go"
		newContent := "This is an updated introduction to Go programming."
		articleSystem.WriteArticle(title, newContent)
	}()

	time.Sleep(500 * time.Millisecond) // Allow time for goroutines to finish

	// Simulating more readers after the update
	for i := 3; i < 6; i++ {
		go func(userNum int) {
			title := "Introduction to Go"
			article, exists := articleSystem.ReadArticle(title)
			if exists {
				fmt.Printf("User %d is reading '%s': %s\n", userNum, title, article.Content)
			} else {
				fmt.Printf("User %d couldn't find article '%s'.\n", userNum, title)
			}
		}(i + 1)
	}

	time.Sleep(500 * time.Millisecond) // Allow time for additional readers

	// Display final state of the article
	finalArticle, exists := articleSystem.ReadArticle("Introduction to Go")
	if exists {
		fmt.Printf("Final state of the article: %s\n", finalArticle.Content)
	} else {
		fmt.Println("Article 'Introduction to Go' not found.")
	}
}
