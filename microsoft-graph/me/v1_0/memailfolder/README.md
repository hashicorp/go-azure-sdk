
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolder` Documentation

The `memailfolder` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/v1.0/memailfolder"
```


### Client Initialization

```go
client := memailfolder.NewMeMailFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderClient.CreateMeMailFolder`

```go
ctx := context.TODO()

payload := memailfolder.MailFolder{
	// ...
}


read, err := client.CreateMeMailFolder(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.CreateMeMailFolderByIdCopy`

```go
ctx := context.TODO()
id := memailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := memailfolder.CreateMeMailFolderByIdCopyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.CreateMeMailFolderByIdMove`

```go
ctx := context.TODO()
id := memailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := memailfolder.CreateMeMailFolderByIdMoveRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.DeleteMeMailFolderById`

```go
ctx := context.TODO()
id := memailfolder.NewMeMailFolderID("mailFolderIdValue")

read, err := client.DeleteMeMailFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.GetMeMailFolderById`

```go
ctx := context.TODO()
id := memailfolder.NewMeMailFolderID("mailFolderIdValue")

read, err := client.GetMeMailFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.GetMeMailFolderCount`

```go
ctx := context.TODO()


read, err := client.GetMeMailFolderCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderClient.ListMeMailFolders`

```go
ctx := context.TODO()


// alternatively `client.ListMeMailFolders(ctx)` can be used to do batched pagination
items, err := client.ListMeMailFoldersComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeMailFolderClient.UpdateMeMailFolderById`

```go
ctx := context.TODO()
id := memailfolder.NewMeMailFolderID("mailFolderIdValue")

payload := memailfolder.MailFolder{
	// ...
}


read, err := client.UpdateMeMailFolderById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
