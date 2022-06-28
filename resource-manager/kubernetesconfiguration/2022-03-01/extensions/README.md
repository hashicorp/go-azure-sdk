
## `github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensions` Documentation

The `extensions` SDK allows for interaction with the Azure Resource Manager Service `kubernetesconfiguration` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/kubernetesconfiguration/2022-03-01/extensions"
```


### Client Initialization

```go
client := extensions.NewExtensionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ExtensionsClient.Create`

```go
ctx := context.TODO()
id := extensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")

payload := extensions.Extension{
	// ...
}

future, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ExtensionsClient.Delete`

```go
ctx := context.TODO()
id := extensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")
future, err := client.Delete(ctx, id, extensions.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ExtensionsClient.Get`

```go
ctx := context.TODO()
id := extensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExtensionsClient.List`

```go
ctx := context.TODO()
id := extensions.NewProviderID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue")
// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExtensionsClient.Update`

```go
ctx := context.TODO()
id := extensions.NewExtensionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterRpValue", "clusterResourceValue", "clusterValue", "extensionValue")

payload := extensions.PatchExtension{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
