
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-06-01/scan` Documentation

The `scan` SDK allows for interaction with Azure Resource Manager `security` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-06-01/scan"
```


### Client Initialization

```go
client := scan.NewScanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ScanClient.DefenderForStorageCancelMalwareScan`

```go
ctx := context.TODO()
id := scan.NewScopedMalwareScanID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "scanId")

read, err := client.DefenderForStorageCancelMalwareScan(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScanClient.DefenderForStorageGetMalwareScan`

```go
ctx := context.TODO()
id := scan.NewScopedMalwareScanID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "scanId")

read, err := client.DefenderForStorageGetMalwareScan(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ScanClient.DefenderForStorageStartMalwareScan`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.DefenderForStorageStartMalwareScan(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
