
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclustersavailableclusterregions` Documentation

The `eventhubsclustersavailableclusterregions` SDK allows for interaction with Azure Resource Manager `eventhub` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclustersavailableclusterregions"
```


### Client Initialization

```go
client := eventhubsclustersavailableclusterregions.NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventHubsClustersAvailableClusterRegionsClient.ClustersListAvailableClusterRegion`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ClustersListAvailableClusterRegion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
