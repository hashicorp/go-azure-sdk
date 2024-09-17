
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-11-01/clusterextensions` Documentation

The `clusterextensions` SDK allows for interaction with Azure Resource Manager `kubernetesconfiguration` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-11-01/clusterextensions"
```


### Client Initialization

```go
client := clusterextensions.NewClusterExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ClusterExtensionsClient.ExtensionsCreate`

```go
ctx := context.TODO()
id := clusterextensions.NewScopedExtensionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "extensionValue")

payload := clusterextensions.Extension{
	// ...
}


if err := client.ExtensionsCreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsDelete`

```go
ctx := context.TODO()
id := clusterextensions.NewScopedExtensionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "extensionValue")

if err := client.ExtensionsDeleteThenPoll(ctx, id, clusterextensions.DefaultExtensionsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsGet`

```go
ctx := context.TODO()
id := clusterextensions.NewScopedExtensionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "extensionValue")

read, err := client.ExtensionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsList`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ExtensionsList(ctx, id)` can be used to do batched pagination
items, err := client.ExtensionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsUpdate`

```go
ctx := context.TODO()
id := clusterextensions.NewScopedExtensionID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "extensionValue")

payload := clusterextensions.PatchExtension{
	// ...
}


if err := client.ExtensionsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
