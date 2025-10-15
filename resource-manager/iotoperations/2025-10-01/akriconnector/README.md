
## `github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnector` Documentation

The `akriconnector` SDK allows for interaction with Azure Resource Manager `iotoperations` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnector"
```


### Client Initialization

```go
client := akriconnector.NewAkriConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AkriConnectorClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := akriconnector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName", "connectorName")

payload := akriconnector.AkriConnectorResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AkriConnectorClient.Delete`

```go
ctx := context.TODO()
id := akriconnector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName", "connectorName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AkriConnectorClient.Get`

```go
ctx := context.TODO()
id := akriconnector.NewConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName", "connectorName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AkriConnectorClient.ListByTemplate`

```go
ctx := context.TODO()
id := akriconnector.NewAkriConnectorTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName")

// alternatively `client.ListByTemplate(ctx, id)` can be used to do batched pagination
items, err := client.ListByTemplateComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
