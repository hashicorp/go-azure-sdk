
## `github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-09-01/contenttemplates` Documentation

The `contenttemplates` SDK allows for interaction with Azure Resource Manager `securityinsights` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/securityinsights/2024-09-01/contenttemplates"
```


### Client Initialization

```go
client := contenttemplates.NewContentTemplatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContentTemplatesClient.ContentTemplateDelete`

```go
ctx := context.TODO()
id := contenttemplates.NewContentTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "templateId")

read, err := client.ContentTemplateDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTemplatesClient.ContentTemplateGet`

```go
ctx := context.TODO()
id := contenttemplates.NewContentTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "templateId")

read, err := client.ContentTemplateGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTemplatesClient.ContentTemplateInstall`

```go
ctx := context.TODO()
id := contenttemplates.NewContentTemplateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "templateId")

payload := contenttemplates.TemplateModel{
	// ...
}


read, err := client.ContentTemplateInstall(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTemplatesClient.List`

```go
ctx := context.TODO()
id := contenttemplates.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.List(ctx, id, contenttemplates.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, contenttemplates.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
