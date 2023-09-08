
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesection` Documentation

The `useronenotesection` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/v1.0/useronenotesection"
```


### Client Initialization

```go
client := useronenotesection.NewUserOnenoteSectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteSectionClient.CreateUserByIdOnenoteSection`

```go
ctx := context.TODO()
id := useronenotesection.NewUserID("userIdValue")

payload := useronenotesection.OnenoteSection{
	// ...
}


read, err := client.CreateUserByIdOnenoteSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.CreateUserByIdOnenoteSectionByIdCopyToNotebook`

```go
ctx := context.TODO()
id := useronenotesection.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

payload := useronenotesection.CreateUserByIdOnenoteSectionByIdCopyToNotebookRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionByIdCopyToNotebook(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.CreateUserByIdOnenoteSectionByIdCopyToSectionGroup`

```go
ctx := context.TODO()
id := useronenotesection.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

payload := useronenotesection.CreateUserByIdOnenoteSectionByIdCopyToSectionGroupRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionByIdCopyToSectionGroup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.DeleteUserByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := useronenotesection.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

read, err := client.DeleteUserByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.GetUserByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := useronenotesection.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteSectionById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.GetUserByIdOnenoteSectionCount`

```go
ctx := context.TODO()
id := useronenotesection.NewUserID("userIdValue")

read, err := client.GetUserByIdOnenoteSectionCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionClient.ListUserByIdOnenoteSections`

```go
ctx := context.TODO()
id := useronenotesection.NewUserID("userIdValue")

// alternatively `client.ListUserByIdOnenoteSections(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteSectionClient.UpdateUserByIdOnenoteSectionById`

```go
ctx := context.TODO()
id := useronenotesection.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

payload := useronenotesection.OnenoteSection{
	// ...
}


read, err := client.UpdateUserByIdOnenoteSectionById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
