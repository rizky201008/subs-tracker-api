package test

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
	"log"
	"testing"
)

var Firebase *firebase.App = initFirebase()

func initFirebase() *firebase.App {
	opt := option.WithCredentialsFile("../substracker-65db9-firebase-adminsdk-na16n-ff2b63130f.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic(err)
	}
	// [END initialize_app_default_golang]

	return app
}

func TestValidateToken(t *testing.T) {
	ctx := context.Background()
	client, err := Firebase.Auth(ctx)
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	token, err := client.VerifyIDToken(ctx, "token_here")
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	assert.Nil(t, err)
	fmt.Println(token.Expires)
	fmt.Println(token.UID)
}
