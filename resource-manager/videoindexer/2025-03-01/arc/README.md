
## `github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-03-01/arc` Documentation

The `arc` SDK allows for interaction with Azure Resource Manager `videoindexer` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-03-01/arc"
```


### Client Initialization

```go
client := arc.NewArcClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ArcClient.GenerateExtensionAccessToken`

```go
ctx := context.TODO()
id := arc.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := arc.GenerateExtensionAccessTokenParameters{
	// ...
}


read, err := client.GenerateExtensionAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ArcClient.GenerateExtensionRestrictedViewerAccessToken`

```go
ctx := context.TODO()
id := arc.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := arc.GenerateExtensionRestrictedViewerAccessTokenParameters{
	// ...
}


read, err := client.GenerateExtensionRestrictedViewerAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
