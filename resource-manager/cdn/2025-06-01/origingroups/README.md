
## `github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/origingroups` Documentation

The `origingroups` SDK allows for interaction with Azure Resource Manager `cdn` (API Version `2025-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cdn/2025-06-01/origingroups"
```


### Client Initialization

```go
client := origingroups.NewOriginGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OriginGroupsClient.Create`

```go
ctx := context.TODO()
id := origingroups.NewEndpointOriginGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName", "endpointName", "originGroupName")

payload := origingroups.OriginGroup{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `OriginGroupsClient.Delete`

```go
ctx := context.TODO()
id := origingroups.NewEndpointOriginGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName", "endpointName", "originGroupName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OriginGroupsClient.Get`

```go
ctx := context.TODO()
id := origingroups.NewEndpointOriginGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName", "endpointName", "originGroupName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OriginGroupsClient.ListByEndpoint`

```go
ctx := context.TODO()
id := origingroups.NewEndpointID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName", "endpointName")

// alternatively `client.ListByEndpoint(ctx, id)` can be used to do batched pagination
items, err := client.ListByEndpointComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OriginGroupsClient.Update`

```go
ctx := context.TODO()
id := origingroups.NewEndpointOriginGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "profileName", "endpointName", "originGroupName")

payload := origingroups.OriginGroupUpdateParameters{
	// ...
}


if err := client.UpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
