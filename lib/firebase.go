package lib

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var ctx = context.Background()
var serviceAccount = option.WithCredentialsFile("./private/adminsdk-priv-key.json")

func Firestore() (*firestore.Client, error) {
	// Initialize Firebase
	app, err := firebase.NewApp(ctx, nil, serviceAccount)
	if err != nil {
		return nil, err
	}

	// Initialize Firestore Client
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
