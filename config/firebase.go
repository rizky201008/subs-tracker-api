package config

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var Firebase *firebase.App

func InitFirebase() {
	// [START initialize_app_default_golang]
	opt := option.WithCredentialsFile("./substracker-65db9-firebase-adminsdk-na16n-ff2b63130f.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	// [END initialize_app_default_golang]

	Firebase = app
}
