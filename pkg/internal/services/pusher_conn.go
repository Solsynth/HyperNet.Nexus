package services

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

// ExtFire is a Firebase App client
var ExtFire *firebase.App

func SetupFirebase() error {
	opt := option.WithCredentialsFile(viper.GetString("firebase_credentials"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	} else {
		ExtFire = app
	}

	return nil
}

// ExtAPNS is an Apple Push Notification Services client
var ExtAPNS *apns2.Client

func SetupAPNS() error {
	authKey, err := token.AuthKeyFromFile(viper.GetString("apns_credentials"))
	if err != nil {
		return err
	}

	ExtAPNS = apns2.NewTokenClient(&token.Token{
		AuthKey: authKey,
		KeyID:   viper.GetString("apns_credentials_key"),
		TeamID:  viper.GetString("apns_credentials_team"),
	}).Production()

	return nil
}
