
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productwiki` Documentation

The `productwiki` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/productwiki"
```


### Client Initialization

```go
client := productwiki.NewProductWikiClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProductWikiClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

payload := productwiki.WikiContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, productwiki.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductWikiClient.Delete`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

read, err := client.Delete(ctx, id, productwiki.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductWikiClient.Get`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductWikiClient.GetEntityTag`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProductWikiClient.List`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ProductWikiClient.Update`

```go
ctx := context.TODO()
id := productwiki.NewProductID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "productIdValue")

payload := productwiki.WikiUpdateContract{
	// ...
}


read, err := client.Update(ctx, id, payload, productwiki.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
