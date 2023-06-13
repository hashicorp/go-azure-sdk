
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/crossregionrestore` Documentation

The `crossregionrestore` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2023-01-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-01-15/crossregionrestore"
```


### Client Initialization

```go
client := crossregionrestore.NewCrossRegionRestoreClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CrossRegionRestoreClient.Trigger`

```go
ctx := context.TODO()
id := crossregionrestore.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

payload := crossregionrestore.CrossRegionRestoreRequest{
	// ...
}


if err := client.TriggerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
