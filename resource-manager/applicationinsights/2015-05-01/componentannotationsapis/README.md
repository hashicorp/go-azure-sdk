
## `github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentannotationsapis` Documentation

The `componentannotationsapis` SDK allows for interaction with Azure Resource Manager `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/componentannotationsapis"
```


### Client Initialization

```go
client := componentannotationsapis.NewComponentAnnotationsAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComponentAnnotationsAPIsClient.AnnotationsCreate`

```go
ctx := context.TODO()
id := componentannotationsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

payload := componentannotationsapis.Annotation{
	// ...
}


read, err := client.AnnotationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentAnnotationsAPIsClient.AnnotationsDelete`

```go
ctx := context.TODO()
id := componentannotationsapis.NewAnnotationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "annotationId")

read, err := client.AnnotationsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentAnnotationsAPIsClient.AnnotationsGet`

```go
ctx := context.TODO()
id := componentannotationsapis.NewAnnotationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName", "annotationId")

read, err := client.AnnotationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComponentAnnotationsAPIsClient.AnnotationsList`

```go
ctx := context.TODO()
id := componentannotationsapis.NewComponentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "componentName")

read, err := client.AnnotationsList(ctx, id, componentannotationsapis.DefaultAnnotationsListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
