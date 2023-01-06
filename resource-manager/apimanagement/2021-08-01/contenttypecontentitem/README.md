
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/contenttypecontentitem` Documentation

The `contenttypecontentitem` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2021-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2021-08-01/contenttypecontentitem"
```


### Client Initialization

```go
client := contenttypecontentitem.NewContentTypeContentItemClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContentTypeContentItemClient.ContentItemCreateOrUpdate`

```go
ctx := context.TODO()
id := contenttypecontentitem.NewContentItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "contentTypeIdValue", "contentItemIdValue")

payload := contenttypecontentitem.ContentItemContract{
	// ...
}


read, err := client.ContentItemCreateOrUpdate(ctx, id, payload, contenttypecontentitem.DefaultContentItemCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTypeContentItemClient.ContentItemDelete`

```go
ctx := context.TODO()
id := contenttypecontentitem.NewContentItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "contentTypeIdValue", "contentItemIdValue")

read, err := client.ContentItemDelete(ctx, id, contenttypecontentitem.DefaultContentItemDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTypeContentItemClient.ContentItemGet`

```go
ctx := context.TODO()
id := contenttypecontentitem.NewContentItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "contentTypeIdValue", "contentItemIdValue")

read, err := client.ContentItemGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTypeContentItemClient.ContentItemGetEntityTag`

```go
ctx := context.TODO()
id := contenttypecontentitem.NewContentItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "contentTypeIdValue", "contentItemIdValue")

read, err := client.ContentItemGetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContentTypeContentItemClient.ContentItemListByService`

```go
ctx := context.TODO()
id := contenttypecontentitem.NewContentTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "contentTypeIdValue")

// alternatively `client.ContentItemListByService(ctx, id)` can be used to do batched pagination
items, err := client.ContentItemListByServiceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
