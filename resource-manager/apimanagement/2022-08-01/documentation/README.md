
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/documentation` Documentation

The `documentation` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/documentation"
```


### Client Initialization

```go
client := documentation.NewDocumentationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DocumentationClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := documentation.NewDocumentationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "documentationIdValue")

payload := documentation.DocumentationContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, documentation.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DocumentationClient.Delete`

```go
ctx := context.TODO()
id := documentation.NewDocumentationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "documentationIdValue")

read, err := client.Delete(ctx, id, documentation.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DocumentationClient.Get`

```go
ctx := context.TODO()
id := documentation.NewDocumentationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "documentationIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DocumentationClient.GetEntityTag`

```go
ctx := context.TODO()
id := documentation.NewDocumentationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "documentationIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DocumentationClient.ListByService`

```go
ctx := context.TODO()
id := documentation.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.ListByService(ctx, id)` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DocumentationClient.Update`

```go
ctx := context.TODO()
id := documentation.NewDocumentationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "documentationIdValue")

payload := documentation.DocumentationUpdateContract{
	// ...
}


read, err := client.Update(ctx, id, payload, documentation.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
