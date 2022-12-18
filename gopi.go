package gopi

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// retrive service from api
func GetService(client *http.Client, srvtype string) interface{} {
	ctx := context.Background()
	var srv interface{}
	var err error
	switch srvtype {
	case "docs":
		srv, err = docs.NewService(ctx, option.WithHTTPClient(client))
	case "sheets":
		srv, err = sheets.NewService(ctx, option.WithHTTPClient(client))
	case "blogger":
		srv, err = blogger.NewService(ctx, option.WithHTTPClient(client))
	case "drive":
		srv, err = drive.NewService(ctx, option.WithHTTPClient(client))
	case "gmail":
		srv, err = gmail.NewService(ctx, option.WithHTTPClient(client))

	}
	if err != nil {
		log.Fatalf("Unable to retrieve Service client: %v", err)
	}
	return srv

}

// Retrieves a token, saves the token, then returns the generated client.
func GetClient(jsonsecfile string, tokFile string, apiscope ...string) *http.Client {
	config := getConfig(jsonsecfile, apiscope...)

	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func getConfig(jsonsecfile string, apiscope ...string) *oauth2.Config {
	b, err := os.ReadFile(jsonsecfile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, apiscope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	return config
}

// Requests a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Requests a token from the web, then returns the retrieved token.
func GenerateTokenFile(jsonsecfile string, tokFile string, apiscope ...string) {
	config := getConfig(jsonsecfile, apiscope...)
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	r, _ := io.Pipe()
	scanner := bufio.NewScanner(r)
	authCode := scanner.Text()

	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	saveToken(tokFile, tok)
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
