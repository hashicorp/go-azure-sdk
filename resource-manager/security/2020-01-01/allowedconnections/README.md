
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/allowedconnections` Documentation

The `allowedconnections` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/allowedconnections"
```


### Client Initialization

```go
client := allowedconnections.NewAllowedConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AllowedConnectionsClient.Get`

```go
ctx := context.TODO()
id := allowedconnections.NewConnectionTypeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationValue", "External")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AllowedConnectionsClient.List`

```go
ctx := context.TODO()
id := allowedconnections.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AllowedConnectionsClient.ListByHomeRegion`

```go
ctx := context.TODO()
id := allowedconnections.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

// alternatively `client.ListByHomeRegion(ctx, id)` can be used to do batched pagination
items, err := client.ListByHomeRegionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
