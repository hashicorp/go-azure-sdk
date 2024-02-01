
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclustersavailableclusterregions` Documentation

The `eventhubsclustersavailableclusterregions` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2022-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2022-01-01-preview/eventhubsclustersavailableclusterregions"
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
