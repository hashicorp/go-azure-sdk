
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/findrestorabletimeranges` Documentation

The `findrestorabletimeranges` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/findrestorabletimeranges"
```


### Client Initialization

```go
client := findrestorabletimeranges.NewFindRestorableTimeRangesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `FindRestorableTimeRangesClient.RestorableTimeRangesFind`

```go
ctx := context.TODO()
id := findrestorabletimeranges.NewBackupInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "backupInstanceValue")

payload := findrestorabletimeranges.AzureBackupFindRestorableTimeRangesRequest{
	// ...
}


read, err := client.RestorableTimeRangesFind(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
