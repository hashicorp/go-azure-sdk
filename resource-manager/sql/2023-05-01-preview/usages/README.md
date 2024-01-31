
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/usages` Documentation

The `usages` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-05-01-preview/usages"
```


### Client Initialization

```go
client := usages.NewUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsagesClient.ListByInstancePool`

```go
ctx := context.TODO()
id := usages.NewInstancePoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "instancePoolValue")

// alternatively `client.ListByInstancePool(ctx, id, usages.DefaultListByInstancePoolOperationOptions())` can be used to do batched pagination
items, err := client.ListByInstancePoolComplete(ctx, id, usages.DefaultListByInstancePoolOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
