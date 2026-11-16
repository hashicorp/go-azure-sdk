
## `github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/authentications` Documentation

The `authentications` SDK allows for interaction with Azure Resource Manager `horizondb` (API Version `2026-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/horizondb/2026-10-01-preview/authentications"
```


### Client Initialization

```go
client := authentications.NewAuthenticationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthenticationsClient.List`

```go
ctx := context.TODO()
id := authentications.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
