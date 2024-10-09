
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/customimages` Documentation

The `customimages` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/customimages"
```


### Client Initialization

```go
client := customimages.NewCustomImagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CustomImagesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := customimages.NewCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "customImageName")

payload := customimages.CustomImage{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `CustomImagesClient.Delete`

```go
ctx := context.TODO()
id := customimages.NewCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "customImageName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `CustomImagesClient.Get`

```go
ctx := context.TODO()
id := customimages.NewCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "customImageName")

read, err := client.Get(ctx, id, customimages.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CustomImagesClient.List`

```go
ctx := context.TODO()
id := customimages.NewLabID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName")

// alternatively `client.List(ctx, id, customimages.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, customimages.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CustomImagesClient.Update`

```go
ctx := context.TODO()
id := customimages.NewCustomImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "customImageName")

payload := customimages.UpdateResource{
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
