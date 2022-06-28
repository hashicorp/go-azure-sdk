
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersavailableclusterregions` Documentation

The `eventhubsclustersavailableclusterregions` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2018-01-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2018-01-01-preview/eventhubsclustersavailableclusterregions"
```


### Client Initialization

```go
client := eventhubsclustersavailableclusterregions.NewEventHubsClustersAvailableClusterRegionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `EventHubsClustersAvailableClusterRegionsClient.ClustersListAvailableClusterRegion`

```go
ctx := context.TODO()
id := eventhubsclustersavailableclusterregions.NewSubscriptionID()
read, err := client.ClustersListAvailableClusterRegion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
