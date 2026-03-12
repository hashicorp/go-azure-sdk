
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/users` Documentation

The `users` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/users"
```


### Client Initialization

```go
client := users.NewUsersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsersClient.GetPublishingUser`

```go
ctx := context.TODO()


read, err := client.GetPublishingUser(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UsersClient.UpdatePublishingUser`

```go
ctx := context.TODO()

payload := users.User{
	// ...
}


read, err := client.UpdatePublishingUser(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
