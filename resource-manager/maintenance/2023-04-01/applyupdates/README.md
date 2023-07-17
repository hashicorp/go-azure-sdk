
## `github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/applyupdates` Documentation

The `applyupdates` SDK allows for interaction with the Azure Resource Manager Service `maintenance` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/applyupdates"
```


### Client Initialization

```go
client := applyupdates.NewApplyUpdatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplyUpdatesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := applyupdates.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.CreateOrUpdate(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplyUpdatesClient.CreateOrUpdateParent`

```go
ctx := context.TODO()
id := applyupdates.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.CreateOrUpdateParent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplyUpdatesClient.Get`

```go
ctx := context.TODO()
id := applyupdates.NewScopedApplyUpdateID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "applyUpdateValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplyUpdatesClient.GetParent`

```go
ctx := context.TODO()
id := applyupdates.NewScopedApplyUpdateID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "applyUpdateValue")

read, err := client.GetParent(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
