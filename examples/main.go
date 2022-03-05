package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jamfproapi "github.com/kenchan0130/jamf-pro-api-client-go"
)

const (
	USER_NAME = ""
	PASSWORD  = ""

	URL = "https://yourserver.jamfcloud.com/api"
)

func main() {
	ctx := context.Background()
	token := getToken()
	c, err := jamfproapi.NewClientWithResponses(URL, jamfproapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *token))
		return nil
	}))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	res, err := c.GetUserWithResponse(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	accounts, _ := json.Marshal(res.JSON200)
	fmt.Printf("Accounts - \n%s\n", string(accounts))
}

func getToken() *string {
	ctx := context.Background()
	c, err := jamfproapi.NewClientWithResponses(URL, jamfproapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(USER_NAME, PASSWORD)
		return nil
	}))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := c.PostAuthTokensWithResponse(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return res.JSON200.Token
}
