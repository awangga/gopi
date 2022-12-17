# gopi
Golang Google Api. To access Google API, please download client_secret_file.json from Credentials in APIs & Services Menu in Your Google Cloud Console Project.
![Credentials in APIs & Services Menu](/docs/img/creds.jpg "Credentials Json Location")

# The flow
service + body = execute

# For Contributor
Testing before commit and push
```sh
go test
```

Release Version
```sh
git tag v0.1.2
git push origin --tags
go list -m github.com/awangga/gogis@v0.1.2
```

## Service
Open google api service with json credentials file and token.pickle, please run in localhost first to generate token.pickle with user confirmation. after that you may put token.pickle in your server.

## Body
Generate json or dictionary for data post to Google API

## Execute
Sending body into Google API and get response

## Example
First thing is import google api module and others helpers you need, after that please define apiscope,jsonsecfile and tokenpickel
```python
from googleapi import service,body,execute
from helper import email

apiscope=['https://www.googleapis.com/auth/spreadsheets', 'https://www.googleapis.com/auth/documents', 'https://www.googleapis.com/auth/drive','https://www.googleapis.com/auth/blogger','https://www.googleapis.com/auth/gmail.send','https://www.googleapis.com/auth/gmail.readonly']
jsonsecfile='client_secret_file.json'
tokenpickle='token.pickle'
```

### Sending email
First we create Mime Text Email with helper library : 
```python
msg=email.createMessage('Rolly Maulana Awangga <awangga@ulbi.ac.id>','rolly@awang.ga','my info',"hello gaes","plain")
```
After that just passing the argument with the variabel above with flow : service -> body -> execute
```python
srv=service.Open('gmail',apiscope,jsonsecfile,tokenpickle)

json=body.GmailSend(msg)

resp=execute.GmailSend(srv,json)
```
Thats all. If u want to catch response from google API just use resp (json format).