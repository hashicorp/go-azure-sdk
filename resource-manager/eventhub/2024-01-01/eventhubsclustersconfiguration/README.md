
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersconfiguration` Documentation

The `eventhubsclustersconfiguration` SDK allows for interaction with Azure Resource Manager `eventhub` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2024-01-01/eventhubsclustersconfiguration"
```


### Client Initialization

```go
client := eventhubsclustersconfiguration.NewEventHubsClustersConfigurationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventHubsClustersConfigurationClient.ConfigurationGet`

```go
ctx := context.TODO()
id := eventhubsclustersconfiguration.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

read, err := client.ConfigurationGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventHubsClustersConfigurationClient.ConfigurationPatch`

```go
ctx := context.TODO()
id := eventhubsclustersconfiguration.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := eventhubsclustersconfiguration.ClusterQuotaConfigurationProperties{
	// ...
}


read, err := client.ConfigurationPatch(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
