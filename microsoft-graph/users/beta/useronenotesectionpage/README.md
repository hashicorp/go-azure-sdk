
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotesectionpage` Documentation

The `useronenotesectionpage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotesectionpage"
```


### Client Initialization

```go
client := useronenotesectionpage.NewUserOnenoteSectionPageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenoteSectionPageClient.CreateUserByIdOnenoteSectionByIdPage`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

payload := useronenotesectionpage.OnenotePage{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionByIdPage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.CreateUserByIdOnenoteSectionByIdPageByIdCopyToSection`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionPageID("userIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectionpage.CreateUserByIdOnenoteSectionByIdPageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionByIdPageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.CreateUserByIdOnenoteSectionByIdPageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionPageID("userIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectionpage.CreateUserByIdOnenoteSectionByIdPageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateUserByIdOnenoteSectionByIdPageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.DeleteUserByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionPageID("userIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.DeleteUserByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.GetUserByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionPageID("userIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

read, err := client.GetUserByIdOnenoteSectionByIdPageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.GetUserByIdOnenoteSectionByIdPageCount`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

read, err := client.GetUserByIdOnenoteSectionByIdPageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenoteSectionPageClient.ListUserByIdOnenoteSectionByIdPages`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionID("userIdValue", "onenoteSectionIdValue")

// alternatively `client.ListUserByIdOnenoteSectionByIdPages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenoteSectionByIdPagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenoteSectionPageClient.UpdateUserByIdOnenoteSectionByIdPageById`

```go
ctx := context.TODO()
id := useronenotesectionpage.NewUserOnenoteSectionPageID("userIdValue", "onenoteSectionIdValue", "onenotePageIdValue")

payload := useronenotesectionpage.OnenotePage{
	// ...
}


read, err := client.UpdateUserByIdOnenoteSectionByIdPageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
