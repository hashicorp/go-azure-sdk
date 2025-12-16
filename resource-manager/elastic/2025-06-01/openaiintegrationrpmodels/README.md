
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/openaiintegrationrpmodels` Documentation

The `openaiintegrationrpmodels` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2025-06-01/openaiintegrationrpmodels"
```


### Client Initialization

```go
client := openaiintegrationrpmodels.NewOpenAIIntegrationRPModelsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenAIIntegrationRPModelsClient.OpenAICreateOrUpdate`

```go
ctx := context.TODO()
id := openaiintegrationrpmodels.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

payload := openaiintegrationrpmodels.OpenAIIntegrationRPModel{
	// ...
}


read, err := client.OpenAICreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationRPModelsClient.OpenAIDelete`

```go
ctx := context.TODO()
id := openaiintegrationrpmodels.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationRPModelsClient.OpenAIGet`

```go
ctx := context.TODO()
id := openaiintegrationrpmodels.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationRPModelsClient.OpenAIGetStatus`

```go
ctx := context.TODO()
id := openaiintegrationrpmodels.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIGetStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationRPModelsClient.OpenAIList`

```go
ctx := context.TODO()
id := openaiintegrationrpmodels.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.OpenAIList(ctx, id)` can be used to do batched pagination
items, err := client.OpenAIListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
