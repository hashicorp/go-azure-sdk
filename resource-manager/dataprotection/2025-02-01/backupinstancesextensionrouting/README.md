
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-02-01/backupinstancesextensionrouting` Documentation

The `backupinstancesextensionrouting` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-02-01/backupinstancesextensionrouting"
```


### Client Initialization

```go
client := backupinstancesextensionrouting.NewBackupInstancesExtensionRoutingClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BackupInstancesExtensionRoutingClient.List`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
