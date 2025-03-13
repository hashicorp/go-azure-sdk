
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/capabilityhost` Documentation

The `capabilityhost` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/capabilityhost"
```


### Client Initialization

```go
client := capabilityhost.NewCapabilityHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CapabilityHostClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := capabilityhost.NewCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "capabilityHostName")

payload := capabilityhost.CapabilityHostResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CapabilityHostClient.Delete`

```go
ctx := context.TODO()
id := capabilityhost.NewCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "capabilityHostName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CapabilityHostClient.Get`

```go
ctx := context.TODO()
id := capabilityhost.NewCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "capabilityHostName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
