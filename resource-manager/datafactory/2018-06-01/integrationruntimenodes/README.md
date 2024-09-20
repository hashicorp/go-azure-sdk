
## `github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimenodes` Documentation

The `integrationruntimenodes` SDK allows for interaction with Azure Resource Manager `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimenodes"
```


### Client Initialization

```go
client := integrationruntimenodes.NewIntegrationRuntimeNodesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IntegrationRuntimeNodesClient.Delete`

```go
ctx := context.TODO()
id := integrationruntimenodes.NewNodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "integrationRuntimeName", "nodeName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationRuntimeNodesClient.Get`

```go
ctx := context.TODO()
id := integrationruntimenodes.NewNodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "integrationRuntimeName", "nodeName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationRuntimeNodesClient.GetIPAddress`

```go
ctx := context.TODO()
id := integrationruntimenodes.NewNodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "integrationRuntimeName", "nodeName")

read, err := client.GetIPAddress(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IntegrationRuntimeNodesClient.Update`

```go
ctx := context.TODO()
id := integrationruntimenodes.NewNodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryName", "integrationRuntimeName", "nodeName")

payload := integrationruntimenodes.UpdateIntegrationRuntimeNodeRequest{
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
