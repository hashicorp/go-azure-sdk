
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-04-01/datamove` Documentation

The `datamove` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2023-04-01/datamove"
```


### Client Initialization

```go
client := datamove.NewDataMoveClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataMoveClient.BMSPrepareDataMove`

```go
ctx := context.TODO()
id := datamove.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := datamove.PrepareDataMoveRequest{
	// ...
}


if err := client.BMSPrepareDataMoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DataMoveClient.BMSTriggerDataMove`

```go
ctx := context.TODO()
id := datamove.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := datamove.TriggerDataMoveRequest{
	// ...
}


if err := client.BMSTriggerDataMoveThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
