# gopi
Golang Google Api. To access Google API, please download client_secret_file.json from Credentials in APIs & Services Menu in Your Google Cloud Console Project.
![Credentials in APIs & Services Menu](/docs/img/creds.jpg "Credentials Json Location")

# The flow
getClient + NewService + [body] = Do()

[] means optional

# For Contributor
Environtment Variables
```sh
GOPROXY=proxy.golang.org
```

Testing before commit and push
```sh
go test
git tag
```

Release Version
```sh
git tag v0.1.2
git push origin --tags
go list -m github.com/awangga/gopi@v0.1.2
```

## getClient
Open google api service with json credentials file and tokenfile, please run in localhost first to generate token.json with user confirmation. after that you may put token.json in your server.

## NewService
Select service for docs,mail,drive etc.

## Body
Generate json or dictionary for data post to Google API

## DO()
Sending request with or without body into Google API and get response

## Example
First thing is import google api module and others helpers you need, after that please define apiscope,jsonsecfile and tokenfile
```go
var apiscope = []string{"https://www.googleapis.com/auth/spreadsheets", "https://www.googleapis.com/auth/documents", "https://www.googleapis.com/auth/drive", "https://www.googleapis.com/auth/blogger", "https://www.googleapis.com/auth/gmail.send", "https://www.googleapis.com/auth/gmail.readonly"}

const jsonsecfile = "credentials.json"
const tokenfile = "token.json"
```

### Reading google docs
First import library : 
```go
import "github.com/awangga/gopi"
```
After that use in your main package
```go
	ctx := context.Background()
	client := gopi.GetClient(jsonsecfile, tokenfile, apiscope...)
	srv, err := docs.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
	}

	// Prints the title of the requested doc:
	// https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
	docId := "195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE"
	doc, err := srv.Documents.Get(docId).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}
	fmt.Printf("The title of the doc is: %s\n", doc.Title)

```
Thats all. If u want to catch response from google API just use doc (json format).
