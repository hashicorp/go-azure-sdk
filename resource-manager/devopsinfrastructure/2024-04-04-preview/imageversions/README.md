
## `github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/imageversions` Documentation

The `imageversions` SDK allows for interaction with the Azure Resource Manager Service `devopsinfrastructure` (API Version `2024-04-04-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devopsinfrastructure/2024-04-04-preview/imageversions"
```


### Client Initialization

```go
client := imageversions.NewImageVersionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ImageVersionsClient.ListByImage`

```go
ctx := context.TODO()
id := imageversions.NewImageID("12345678-1234-9876-4563-123456789012", "example-resource-group", "imageValue")

// alternatively `client.ListByImage(ctx, id)` can be used to do batched pagination
items, err := client.ListByImageComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
