
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotepage` Documentation

The `useronenotepage` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/useronenotepage"
```


### Client Initialization

```go
client := useronenotepage.NewUserOnenotePageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserOnenotePageClient.CreateUserByIdOnenotePage`

```go
ctx := context.TODO()
id := useronenotepage.NewUserID("userIdValue")

payload := useronenotepage.OnenotePage{
	// ...
}


read, err := client.CreateUserByIdOnenotePage(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.CreateUserByIdOnenotePageByIdCopyToSection`

```go
ctx := context.TODO()
id := useronenotepage.NewUserOnenotePageID("userIdValue", "onenotePageIdValue")

payload := useronenotepage.CreateUserByIdOnenotePageByIdCopyToSectionRequest{
	// ...
}


read, err := client.CreateUserByIdOnenotePageByIdCopyToSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.CreateUserByIdOnenotePageByIdOnenotePatchContent`

```go
ctx := context.TODO()
id := useronenotepage.NewUserOnenotePageID("userIdValue", "onenotePageIdValue")

payload := useronenotepage.CreateUserByIdOnenotePageByIdOnenotePatchContentRequest{
	// ...
}


read, err := client.CreateUserByIdOnenotePageByIdOnenotePatchContent(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.DeleteUserByIdOnenotePageById`

```go
ctx := context.TODO()
id := useronenotepage.NewUserOnenotePageID("userIdValue", "onenotePageIdValue")

read, err := client.DeleteUserByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.GetUserByIdOnenotePageById`

```go
ctx := context.TODO()
id := useronenotepage.NewUserOnenotePageID("userIdValue", "onenotePageIdValue")

read, err := client.GetUserByIdOnenotePageById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.GetUserByIdOnenotePageCount`

```go
ctx := context.TODO()
id := useronenotepage.NewUserID("userIdValue")

read, err := client.GetUserByIdOnenotePageCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserOnenotePageClient.ListUserByIdOnenotePages`

```go
ctx := context.TODO()
id := useronenotepage.NewUserID("userIdValue")

// alternatively `client.ListUserByIdOnenotePages(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdOnenotePagesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserOnenotePageClient.UpdateUserByIdOnenotePageById`

```go
ctx := context.TODO()
id := useronenotepage.NewUserOnenotePageID("userIdValue", "onenotePageIdValue")

payload := useronenotepage.OnenotePage{
	// ...
}


read, err := client.UpdateUserByIdOnenotePageById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
