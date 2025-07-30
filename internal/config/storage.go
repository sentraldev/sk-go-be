package config

import (
	"log"
	"os"

	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"

	"google.golang.org/api/option"
)

var SC *storage.Client

func InitializeFirebaseStorage(opt option.ClientOption) {
	ctx := context.Background()

	firebaseConfig := firebase.Config{
		StorageBucket: os.Getenv("FIREBASE_STORAGE_BUCKET"),
	}
	app, err := firebase.NewApp(context.Background(), &firebaseConfig, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	// Get a Storage client
	SC, err = app.Storage(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Storage client: %v", err)
	}

	// Get a reference to the default bucket
	_, err = SC.DefaultBucket()
	if err != nil {
		log.Fatalf("Failed to get default bucket: %v", err)
	}

	log.Println("Firebase Storage initialized successfully")
}
