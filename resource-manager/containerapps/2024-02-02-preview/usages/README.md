
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/usages` Documentation

The `usages` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2024-02-02-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/usages"
```


### Client Initialization

```go
client := usages.NewUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UsagesClient.List`

```go
ctx := context.TODO()
id := usages.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UsagesClient.ManagedEnvironmentUsagesList`

```go
ctx := context.TODO()
id := usages.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentName")

// alternatively `client.ManagedEnvironmentUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.ManagedEnvironmentUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
