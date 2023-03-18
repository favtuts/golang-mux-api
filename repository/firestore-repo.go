package repository

import (
	"context"
	"golang-mux-api/entity"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct{}

// NewFirestoreRepository creates a new repo
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-reviews-62551"
	collectionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	iter := client.Collection(collectionName).Documents(ctx)

	/*
		p, err := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		log.Println(p.Data()["Title"].(string))

		p1, err := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		log.Println(p1.Data()["Title"].(string))

		p2, err := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		log.Println(p2.Data()["Title"].(string))

		p3, err := iter.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			//return posts, nil
		} else {
			log.Println(p3.Data()["Title"].(string))
		}
	*/

	for {
		doc, err := iter.Next()
		//Ref: https://stackoverflow.com/questions/52626974/how-to-return-document-index-name-value
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// Delete: TODO
func (r *repo) Delete(post *entity.Post) error {
	return nil
}
