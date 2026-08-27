
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-05-01/identitybindings` Documentation

The `identitybindings` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2026-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2026-05-01/identitybindings"
```


### Client Initialization

```go
client := identitybindings.NewIdentityBindingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityBindingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := identitybindings.NewIdentityBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "identityBindingName")

payload := identitybindings.IdentityBinding{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `IdentityBindingsClient.Delete`

```go
ctx := context.TODO()
id := identitybindings.NewIdentityBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "identityBindingName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `IdentityBindingsClient.Get`

```go
ctx := context.TODO()
id := identitybindings.NewIdentityBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "identityBindingName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityBindingsClient.ListByManagedCluster`

```go
ctx := context.TODO()
id := commonids.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName")

// alternatively `client.ListByManagedCluster(ctx, id)` can be used to do batched pagination
items, err := client.ListByManagedClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
