
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/poolchange` Documentation

The `poolchange` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-09-01/poolchange"
```


### Client Initialization

```go
client := poolchange.NewPoolChangeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PoolChangeClient.VolumesPoolChange`

```go
ctx := context.TODO()
id := poolchange.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue")

payload := poolchange.PoolChangeRequest{
	// ...
}


if err := client.VolumesPoolChangeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
