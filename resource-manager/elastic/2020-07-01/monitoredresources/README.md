
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitoredresources` Documentation

The `monitoredresources` SDK allows for interaction with the Azure Resource Manager Service `elastic` (API Version `2020-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/monitoredresources"
```


### Client Initialization

```go
client := monitoredresources.NewMonitoredResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitoredResourcesClient.List`

```go
ctx := context.TODO()
id := monitoredresources.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
