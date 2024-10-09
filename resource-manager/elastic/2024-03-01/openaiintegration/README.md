
## `github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/openaiintegration` Documentation

The `openaiintegration` SDK allows for interaction with Azure Resource Manager `elastic` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/elastic/2024-03-01/openaiintegration"
```


### Client Initialization

```go
client := openaiintegration.NewOpenAIIntegrationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenAIIntegrationClient.OpenAICreateOrUpdate`

```go
ctx := context.TODO()
id := openaiintegration.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

payload := openaiintegration.OpenAIIntegrationRPModel{
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


### Example Usage: `OpenAIIntegrationClient.OpenAIDelete`

```go
ctx := context.TODO()
id := openaiintegration.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationClient.OpenAIGet`

```go
ctx := context.TODO()
id := openaiintegration.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationClient.OpenAIGetStatus`

```go
ctx := context.TODO()
id := openaiintegration.NewOpenAIIntegrationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName", "openAIIntegrationName")

read, err := client.OpenAIGetStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenAIIntegrationClient.OpenAIList`

```go
ctx := context.TODO()
id := openaiintegration.NewMonitorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "monitorName")

// alternatively `client.OpenAIList(ctx, id)` can be used to do batched pagination
items, err := client.OpenAIListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
