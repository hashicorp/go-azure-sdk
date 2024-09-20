
## `github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-02-01/templatespecs` Documentation

The `templatespecs` SDK allows for interaction with Azure Resource Manager `resources` (API Version `2022-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/resources/2022-02-01/templatespecs"
```


### Client Initialization

```go
client := templatespecs.NewTemplateSpecsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TemplateSpecsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := templatespecs.NewTemplateSpecID("12345678-1234-9876-4563-123456789012", "example-resource-group", "templateSpecName")

payload := templatespecs.TemplateSpec{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateSpecsClient.Delete`

```go
ctx := context.TODO()
id := templatespecs.NewTemplateSpecID("12345678-1234-9876-4563-123456789012", "example-resource-group", "templateSpecName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateSpecsClient.Get`

```go
ctx := context.TODO()
id := templatespecs.NewTemplateSpecID("12345678-1234-9876-4563-123456789012", "example-resource-group", "templateSpecName")

read, err := client.Get(ctx, id, templatespecs.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateSpecsClient.GetBuiltIn`

```go
ctx := context.TODO()
id := templatespecs.NewBuiltInTemplateSpecID("templateSpecName")

read, err := client.GetBuiltIn(ctx, id, templatespecs.DefaultGetBuiltInOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TemplateSpecsClient.ListBuiltIns`

```go
ctx := context.TODO()


// alternatively `client.ListBuiltIns(ctx, templatespecs.DefaultListBuiltInsOperationOptions())` can be used to do batched pagination
items, err := client.ListBuiltInsComplete(ctx, templatespecs.DefaultListBuiltInsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateSpecsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id, templatespecs.DefaultListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id, templatespecs.DefaultListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateSpecsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id, templatespecs.DefaultListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id, templatespecs.DefaultListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `TemplateSpecsClient.Update`

```go
ctx := context.TODO()
id := templatespecs.NewTemplateSpecID("12345678-1234-9876-4563-123456789012", "example-resource-group", "templateSpecName")

payload := templatespecs.TemplateSpecUpdateModel{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
