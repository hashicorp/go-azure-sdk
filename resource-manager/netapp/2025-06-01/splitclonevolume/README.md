
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/splitclonevolume` Documentation

The `splitclonevolume` SDK allows for interaction with Azure Resource Manager `netapp` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2025-06-01/splitclonevolume"
```


### Client Initialization

```go
client := splitclonevolume.NewSplitCloneVolumeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SplitCloneVolumeClient.VolumesSplitCloneFromParent`

```go
ctx := context.TODO()
id := splitclonevolume.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountName", "capacityPoolName", "volumeName")

if err := client.VolumesSplitCloneFromParentThenPoll(ctx, id); err != nil {
	// handle the error
}
```
