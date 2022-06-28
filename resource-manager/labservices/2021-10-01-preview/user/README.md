
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/user` Documentation

The `user` SDK allows for interaction with the Azure Resource Manager Service `labservices` (API Version `2021-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/user"
```


### Client Initialization

```go
client := user.NewUserClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `UserClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := user.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue")

payload := user.User{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `UserClient.Delete`

```go
ctx := context.TODO()
id := user.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `UserClient.Get`

```go
ctx := context.TODO()
id := user.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.Invite`

```go
ctx := context.TODO()
id := user.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue")

payload := user.InviteBody{
	// ...
}

future, err := client.Invite(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `UserClient.ListByLab`

```go
ctx := context.TODO()
id := user.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue")
// alternatively `client.ListByLab(ctx, id)` can be used to do batched pagination
items, err := client.ListByLabComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.Update`

```go
ctx := context.TODO()
id := user.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labValue", "userValue")

payload := user.UserUpdate{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
