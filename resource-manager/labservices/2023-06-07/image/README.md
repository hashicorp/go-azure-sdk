
## `github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/image` Documentation

The `image` SDK allows for interaction with Azure Resource Manager `labservices` (API Version `2023-06-07`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/image"
```


### Client Initialization

```go
client := image.NewImageClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImageClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := image.NewImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanName", "imageName")

payload := image.Image{
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


### Example Usage: `ImageClient.Get`

```go
ctx := context.TODO()
id := image.NewImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanName", "imageName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ImageClient.ListByLabPlan`

```go
ctx := context.TODO()
id := image.NewLabPlanID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanName")

// alternatively `client.ListByLabPlan(ctx, id)` can be used to do batched pagination
items, err := client.ListByLabPlanComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ImageClient.Update`

```go
ctx := context.TODO()
id := image.NewImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labPlanName", "imageName")

payload := image.ImageUpdate{
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
