
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebooksection` Documentation

The `useronenotenotebooksection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotenotebooksection"
```


### Client Initialization

```go
client := useronenotenotebooksection.NewUserOnenoteNotebookSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteNotebookSectionClient.CreateUserByIdOnenoteNotebookByIdSection`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

payload := useronenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksection.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksection.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteNotebookByIdSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.DeleteUserByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.DeleteUserByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.GetUserByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.GetUserByIdOnenoteNotebookByIdSectionCount`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

read, err := client.GetUserByIdOnenoteNotebookByIdSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.ListUserByIdOnenoteNotebookByIdSections`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookID("userIdValue", "notebookIdValue")

// alternatively `client.ListUserByIdOnenoteNotebookByIdSections(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteNotebookByIdSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteNotebookSectionClient.UpdateUserByIdOnenoteNotebookByIdSectionById`

```go
ctx := context.TODO()
id := useronenotenotebooksection.NewUserOnenoteNotebookSectionID("userIdValue", "notebookIdValue", "onenoteSectionIdValue")

payload := useronenotenotebooksection.OnenoteSection{
	// ...
}


read, err := client.UpdateUserByIdOnenoteNotebookByIdSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
