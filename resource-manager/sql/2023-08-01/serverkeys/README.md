
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/serverkeys` Documentation

The `serverkeys` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01/serverkeys"
```


### Client Initialization

```go
client := serverkeys.NewServerKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerKeysClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := serverkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "keyName")

payload := serverkeys.ServerKey{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerKeysClient.Delete`

```go
ctx := context.TODO()
id := serverkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "keyName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ServerKeysClient.Get`

```go
ctx := context.TODO()
id := serverkeys.NewKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName", "keyName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerKeysClient.ListByServer`

```go
ctx := context.TODO()
id := commonids.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverName")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
