
## `github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/grafanaplugin` Documentation

The `grafanaplugin` SDK allows for interaction with Azure Resource Manager `dashboard` (API Version `2024-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dashboard/2024-10-01/grafanaplugin"
```


### Client Initialization

```go
client := grafanaplugin.NewGrafanaPluginClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GrafanaPluginClient.GrafanaFetchAvailablePlugins`

```go
ctx := context.TODO()
id := grafanaplugin.NewGrafanaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "grafanaName")

// alternatively `client.GrafanaFetchAvailablePlugins(ctx, id)` can be used to do batched pagination
items, err := client.GrafanaFetchAvailablePluginsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
