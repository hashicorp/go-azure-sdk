
## `github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2024-06-01/snapshots` Documentation

The `snapshots` SDK allows for interaction with Azure Resource Manager `appconfiguration` (API Version `2024-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/appconfiguration/2024-06-01/snapshots"
```


### Client Initialization

```go
client := snapshots.NewSnapshotsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SnapshotsClient.Create`

```go
ctx := context.TODO()
id := snapshots.NewSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreName", "snapshotName")

payload := snapshots.Snapshot{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `SnapshotsClient.Get`

```go
ctx := context.TODO()
id := snapshots.NewSnapshotID("12345678-1234-9876-4563-123456789012", "example-resource-group", "configurationStoreName", "snapshotName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
