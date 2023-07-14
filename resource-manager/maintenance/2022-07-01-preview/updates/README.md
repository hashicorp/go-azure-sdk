
## `github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/updates` Documentation

The `updates` SDK allows for interaction with the Azure Resource Manager Service `maintenance` (API Version `2022-07-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/updates"
```


### Client Initialization

```go
client := updates.NewUpdatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UpdatesClient.List`

```go
ctx := context.TODO()
id := updates.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UpdatesClient.ListParent`

```go
ctx := context.TODO()
id := updates.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.ListParent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
