
## `github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/discoveryrules` Documentation

The `discoveryrules` SDK allows for interaction with Azure Resource Manager `cloudhealth` (API Version `2026-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cloudhealth/2026-11-01/discoveryrules"
```


### Client Initialization

```go
client := discoveryrules.NewDiscoveryRulesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DiscoveryRulesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := discoveryrules.NewDiscoveryruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "discoveryruleName")

payload := discoveryrules.DiscoveryRule{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveryRulesClient.Delete`

```go
ctx := context.TODO()
id := discoveryrules.NewDiscoveryruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "discoveryruleName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `DiscoveryRulesClient.Get`

```go
ctx := context.TODO()
id := discoveryrules.NewDiscoveryruleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName", "discoveryruleName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DiscoveryRulesClient.ListByHealthModel`

```go
ctx := context.TODO()
id := discoveryrules.NewHealthmodelID("12345678-1234-9876-4563-123456789012", "example-resource-group", "healthmodelName")

// alternatively `client.ListByHealthModel(ctx, id, discoveryrules.DefaultListByHealthModelOperationOptions())` can be used to do batched pagination
items, err := client.ListByHealthModelComplete(ctx, id, discoveryrules.DefaultListByHealthModelOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
