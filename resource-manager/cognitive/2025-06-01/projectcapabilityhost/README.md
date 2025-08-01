
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/projectcapabilityhost` Documentation

The `projectcapabilityhost` SDK allows for interaction with Azure Resource Manager `cognitive` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2025-06-01/projectcapabilityhost"
```


### Client Initialization

```go
client := projectcapabilityhost.NewProjectCapabilityHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProjectCapabilityHostClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := projectcapabilityhost.NewProjectCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "capabilityHostName")

payload := projectcapabilityhost.CapabilityHostResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ProjectCapabilityHostClient.Delete`

```go
ctx := context.TODO()
id := projectcapabilityhost.NewProjectCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "capabilityHostName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ProjectCapabilityHostClient.Get`

```go
ctx := context.TODO()
id := projectcapabilityhost.NewProjectCapabilityHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName", "projectName", "capabilityHostName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
