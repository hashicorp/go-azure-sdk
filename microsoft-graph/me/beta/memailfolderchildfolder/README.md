
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfolderchildfolder` Documentation

The `memailfolderchildfolder` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/memailfolderchildfolder"
```


### Client Initialization

```go
client := memailfolderchildfolder.NewMeMailFolderChildFolderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeMailFolderChildFolderClient.CreateMeMailFolderByIdChildFolder`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderID("mailFolderIdValue")

payload := memailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.CreateMeMailFolderByIdChildFolderByIdCopy`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

payload := memailfolderchildfolder.CreateMeMailFolderByIdChildFolderByIdCopyRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.CreateMeMailFolderByIdChildFolderByIdMove`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

payload := memailfolderchildfolder.CreateMeMailFolderByIdChildFolderByIdMoveRequest{
	// ...
}


read, err := client.CreateMeMailFolderByIdChildFolderByIdMove(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.DeleteMeMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

read, err := client.DeleteMeMailFolderByIdChildFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.GetMeMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

read, err := client.GetMeMailFolderByIdChildFolderById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.GetMeMailFolderByIdChildFolderCount`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderID("mailFolderIdValue")

read, err := client.GetMeMailFolderByIdChildFolderCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeMailFolderChildFolderClient.ListMeMailFolderByIdChildFolders`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderID("mailFolderIdValue")

// alternatively `client.ListMeMailFolderByIdChildFolders(ctx, id)` can be used to do batched pagination
items, err := client.ListMeMailFolderByIdChildFoldersComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeMailFolderChildFolderClient.UpdateMeMailFolderByIdChildFolderById`

```go
ctx := context.TODO()
id := memailfolderchildfolder.NewMeMailFolderChildFolderID("mailFolderIdValue", "mailFolderId1Value")

payload := memailfolderchildfolder.MailFolder{
	// ...
}


read, err := client.UpdateMeMailFolderByIdChildFolderById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
