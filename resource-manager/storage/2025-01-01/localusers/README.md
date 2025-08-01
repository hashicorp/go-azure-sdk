
## `github.com/hashicorp/go-azure-sdk/resource-manager/storage/2025-01-01/localusers` Documentation

The `localusers` SDK allows for interaction with Azure Resource Manager `storage` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/storage/2025-01-01/localusers"
```


### Client Initialization

```go
client := localusers.NewLocalUsersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocalUsersClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := localusers.NewLocalUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "localUserName")

payload := localusers.LocalUser{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocalUsersClient.Delete`

```go
ctx := context.TODO()
id := localusers.NewLocalUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "localUserName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocalUsersClient.Get`

```go
ctx := context.TODO()
id := localusers.NewLocalUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "localUserName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocalUsersClient.List`

```go
ctx := context.TODO()
id := commonids.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName")

// alternatively `client.List(ctx, id, localusers.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, localusers.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `LocalUsersClient.ListKeys`

```go
ctx := context.TODO()
id := localusers.NewLocalUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "localUserName")

read, err := client.ListKeys(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `LocalUsersClient.RegeneratePassword`

```go
ctx := context.TODO()
id := localusers.NewLocalUserID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "localUserName")

read, err := client.RegeneratePassword(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
