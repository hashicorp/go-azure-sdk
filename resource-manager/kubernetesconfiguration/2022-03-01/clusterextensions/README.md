
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/clusterextensions` Documentation

The `clusterextensions` SDK allows for interaction with the Azure Resource Manager Service `kubernetesconfiguration` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/clusterextensions"
```


### Client Initialization

```go
client := clusterextensions.NewClusterExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsCreate`

```go
ctx := context.TODO()
id := clusterextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")

payload := clusterextensions.Extension{
	// ...
}

future, err := client.ExtensionsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsDelete`

```go
ctx := context.TODO()
id := clusterextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")
future, err := client.ExtensionsDelete(ctx, id, clusterextensions.DefaultExtensionsDeleteOperationOptions())
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ClusterExtensionsClient.ExtensionsGet`

```go
ctx := context.TODO()
id := clusterextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")
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
id := clusterextensions.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue")
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
id := clusterextensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")

payload := clusterextensions.PatchExtension{
	// ...
}

future, err := client.ExtensionsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
