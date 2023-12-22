
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-07-01/filelocks` Documentation

The `filelocks` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2023-07-01/filelocks"
```


### Client Initialization

```go
client := filelocks.NewFileLocksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FileLocksClient.VolumesBreakFileLocks`

```go
ctx := context.TODO()
id := filelocks.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue")

payload := filelocks.BreakFileLocksRequest{
	// ...
}


if err := client.VolumesBreakFileLocksThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
