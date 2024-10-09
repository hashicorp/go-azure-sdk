
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/users` Documentation

The `users` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/users"
```


### Client Initialization

```go
client := users.NewUsersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := users.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

payload := users.User{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `UsersClient.Delete`

```go
ctx := context.TODO()
id := users.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `UsersClient.Get`

```go
ctx := context.TODO()
id := users.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

read, err := client.Get(ctx, id, users.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UsersClient.List`

```go
ctx := context.TODO()
id := users.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName")

// alternatively `client.List(ctx, id, users.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, users.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UsersClient.Update`

```go
ctx := context.TODO()
id := users.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "userName")

payload := users.UpdateResource{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
