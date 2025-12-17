
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-10-01/managednamespaces` Documentation

The `managednamespaces` SDK allows for interaction with Azure Resource Manager `containerservice` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerservice/2025-10-01/managednamespaces"
```


### Client Initialization

```go
client := managednamespaces.NewManagedNamespacesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedNamespacesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "managedNamespaceName")

payload := managednamespaces.ManagedNamespace{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNamespacesClient.Delete`

```go
ctx := context.TODO()
id := managednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "managedNamespaceName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedNamespacesClient.Get`

```go
ctx := context.TODO()
id := managednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "managedNamespaceName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedNamespacesClient.ListByManagedCluster`

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


### Example Usage: `ManagedNamespacesClient.ListCredential`

```go
ctx := context.TODO()
id := managednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "managedNamespaceName")

read, err := client.ListCredential(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedNamespacesClient.Update`

```go
ctx := context.TODO()
id := managednamespaces.NewManagedNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedClusterName", "managedNamespaceName")

payload := managednamespaces.TagsObject{
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
