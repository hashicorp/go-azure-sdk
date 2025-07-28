
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-06-01/contentproducttemplates` Documentation

The `contentproducttemplates` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2025-06-01/contentproducttemplates"
```


### Client Initialization

```go
client := contentproducttemplates.NewContentProductTemplatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContentProductTemplatesClient.ProductTemplateGet`

```go
ctx := context.TODO()
id := contentproducttemplates.NewContentProductTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "templateId")

read, err := client.ProductTemplateGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentProductTemplatesClient.ProductTemplatesList`

```go
ctx := context.TODO()
id := contentproducttemplates.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.ProductTemplatesList(ctx, id, contentproducttemplates.DefaultProductTemplatesListOperationOptions())` can be used to do batched pagination
items, err := client.ProductTemplatesListComplete(ctx, id, contentproducttemplates.DefaultProductTemplatesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
