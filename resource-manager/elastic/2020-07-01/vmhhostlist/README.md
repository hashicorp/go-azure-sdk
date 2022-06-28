
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmhhostlist` Documentation

The `vmhhostlist` SDK allows for interaction with the Azure Resource Manager Service `elastic` (API Version `2020-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2020-07-01/vmhhostlist"
```


### Client Initialization

```go
client := vmhhostlist.NewVMHHostListClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMHHostListClient.VMHostList`

```go
ctx := context.TODO()
id := vmhhostlist.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorValue")

// alternatively `client.VMHostList(ctx, id)` can be used to do batched pagination
items, err := client.VMHostListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
