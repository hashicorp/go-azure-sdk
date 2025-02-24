
## `github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-03-01/accesstoken` Documentation

The `accesstoken` SDK allows for interaction with Azure Resource Manager `videoindexer` (API Version `2025-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2025-03-01/accesstoken"
```


### Client Initialization

```go
client := accesstoken.NewAccessTokenClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AccessTokenClient.GenerateAccessToken`

```go
ctx := context.TODO()
id := accesstoken.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := accesstoken.GenerateAccessTokenParameters{
	// ...
}


read, err := client.GenerateAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AccessTokenClient.GenerateRestrictedViewerAccessToken`

```go
ctx := context.TODO()
id := accesstoken.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountName")

payload := accesstoken.GenerateRestrictedViewerAccessTokenParameters{
	// ...
}


read, err := client.GenerateRestrictedViewerAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
