
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mailfolder` Documentation

The `mailfolder` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/mailfolder"
```


### Client Initialization

```go
client := mailfolder.NewMailFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MailFolderClient.CreateMailFolder`

```go
ctx := context.TODO()

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.CreateMailFolder(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.CreateMailFolderCopy`

```go
ctx := context.TODO()
id := mailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := mailfolder.CreateMailFolderCopyRequest{
	// ...
}


read, err := client.CreateMailFolderCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.CreateMailFolderMove`

```go
ctx := context.TODO()
id := mailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := mailfolder.CreateMailFolderMoveRequest{
	// ...
}


read, err := client.CreateMailFolderMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.DeleteMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewMeMailFolderID("mailFolderIdValue")

read, err := client.DeleteMailFolder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.GetMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewMeMailFolderID("mailFolderIdValue")

read, err := client.GetMailFolder(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.GetMailFolderCount`

```go
ctx := context.TODO()


read, err := client.GetMailFolderCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MailFolderClient.ListMailFolders`

```go
ctx := context.TODO()


// alternatively `client.ListMailFolders(ctx)` can be used to do batched pagination
items, err := client.ListMailFoldersComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MailFolderClient.UpdateMailFolder`

```go
ctx := context.TODO()
id := mailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := mailfolder.MailFolder{
	// ...
}


read, err := client.UpdateMailFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
