
## `github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/vmhost` Documentation

The `vmhost` SDK allows for interaction with Azure Resource Manager `logz` (API Version `2020-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/logz/2020-10-01/vmhost"
```


### Client Initialization

```go
client := vmhost.NewVMHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VMHostClient.MonitorListVMHostUpdate`

```go
ctx := context.TODO()
id := vmhost.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

payload := vmhost.VMHostUpdateRequest{
	// ...
}


// alternatively `client.MonitorListVMHostUpdate(ctx, id, payload)` can be used to do batched pagination
items, err := client.MonitorListVMHostUpdateComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMHostClient.MonitorListVMHosts`

```go
ctx := context.TODO()
id := vmhost.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.MonitorListVMHosts(ctx, id)` can be used to do batched pagination
items, err := client.MonitorListVMHostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMHostClient.MonitorVMHostPayload`

```go
ctx := context.TODO()
id := vmhost.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

read, err := client.MonitorVMHostPayload(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VMHostClient.SubAccountListVMHostUpdate`

```go
ctx := context.TODO()
id := vmhost.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "subAccountName")

payload := vmhost.VMHostUpdateRequest{
	// ...
}


// alternatively `client.SubAccountListVMHostUpdate(ctx, id, payload)` can be used to do batched pagination
items, err := client.SubAccountListVMHostUpdateComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMHostClient.SubAccountListVMHosts`

```go
ctx := context.TODO()
id := vmhost.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "subAccountName")

// alternatively `client.SubAccountListVMHosts(ctx, id)` can be used to do batched pagination
items, err := client.SubAccountListVMHostsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VMHostClient.SubAccountVMHostPayload`

```go
ctx := context.TODO()
id := vmhost.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "subAccountName")

read, err := client.SubAccountVMHostPayload(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
