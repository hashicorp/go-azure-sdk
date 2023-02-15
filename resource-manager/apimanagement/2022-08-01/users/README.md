
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/users` Documentation

The `users` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/users"
```


### Client Initialization

```go
client := users.NewUsersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsersClient.UserGenerateSsoUrl`

```go
ctx := context.TODO()
id := users.NewUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "userIdValue")

read, err := client.UserGenerateSsoUrl(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
