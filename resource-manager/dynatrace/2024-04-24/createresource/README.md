
## `github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/createresource` Documentation

The `createresource` SDK allows for interaction with Azure Resource Manager `dynatrace` (API Version `2024-04-24`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dynatrace/2024-04-24/createresource"
```


### Client Initialization

```go
client := createresource.NewCreateResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CreateResourceClient.CreationSupportedGet`

```go
ctx := context.TODO()
id := createresource.NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId")

// alternatively `client.CreationSupportedGet(ctx, id)` can be used to do batched pagination
items, err := client.CreationSupportedGetComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CreateResourceClient.CreationSupportedList`

```go
ctx := context.TODO()
id := createresource.NewSubscriptionStatusID("12345678-1234-9876-4563-123456789012", "dynatraceEnvironmentId")

// alternatively `client.CreationSupportedList(ctx, id)` can be used to do batched pagination
items, err := client.CreationSupportedListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
