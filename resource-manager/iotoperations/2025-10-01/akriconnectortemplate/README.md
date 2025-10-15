
## `github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnectortemplate` Documentation

The `akriconnectortemplate` SDK allows for interaction with Azure Resource Manager `iotoperations` (API Version `2025-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/iotoperations/2025-10-01/akriconnectortemplate"
```


### Client Initialization

```go
client := akriconnectortemplate.NewAkriConnectorTemplateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AkriConnectorTemplateClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := akriconnectortemplate.NewAkriConnectorTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName")

payload := akriconnectortemplate.AkriConnectorTemplateResource{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AkriConnectorTemplateClient.Delete`

```go
ctx := context.TODO()
id := akriconnectortemplate.NewAkriConnectorTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `AkriConnectorTemplateClient.Get`

```go
ctx := context.TODO()
id := akriconnectortemplate.NewAkriConnectorTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName", "akriConnectorTemplateName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AkriConnectorTemplateClient.ListByInstanceResource`

```go
ctx := context.TODO()
id := akriconnectortemplate.NewInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instanceName")

// alternatively `client.ListByInstanceResource(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
