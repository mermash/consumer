package main

import (
	"context"
	"fmt"

	// "../sess-client/client"
	apiClient "github.com/mermash/openapi"
)

func main() {

	cfg := apiClient.NewConfiguration()
	cfg.Host = "127.0.0.1:8080"
	cfg.Scheme = "http"
	client := apiClient.NewAPIClient(cfg)
	sessManager := client.AuthCheckerApi

	ctx := context.Background()

	//create session
	var name *string
	var useragent *string
	namestr := "mermash"
	useragentstr := "chrome"
	name = &namestr
	useragent = &useragentstr
	sessId, response, err := sessManager.AuthCheckerCreate(ctx).Body(
		apiClient.SessionSession{
			Login:     name,
			Useragent: useragent,
		},
	).Execute()
	fmt.Println("create sessId", sessId, err, response)

	//check session
	sess, _, err := sessManager.AuthCheckerCheckExecute(
		sessManager.AuthCheckerCheck(ctx, *sessId.ID),
	)
	fmt.Println("after create", sess, err)

	//delete session
	_, response, err = sessManager.AuthCheckerDelete(ctx).Body(
		apiClient.SessionSessionID{
			ID: sessId.ID,
		},
	).Execute()
	fmt.Println("delete sessionID", response, err)

	//check again
	sess, _, err = sessManager.AuthCheckerCheckExecute(
		sessManager.AuthCheckerCheck(ctx, *sessId.ID),
	)
	fmt.Println("after delete", sess, err)
}
