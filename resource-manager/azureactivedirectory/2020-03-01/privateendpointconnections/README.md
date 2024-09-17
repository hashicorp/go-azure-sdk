
## `github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privateendpointconnections` Documentation

The `privateendpointconnections` SDK allows for interaction with Azure Resource Manager `azureactivedirectory` (API Version `2020-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privateendpointconnections"
```


### Client Initialization

```go
client := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateEndpointConnectionsClient.Create`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue", "privateEndpointConnectionValue")

payload := privateendpointconnections.PrivateEndpointConnection{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Delete`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue", "privateEndpointConnectionValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateEndpointConnectionsClient.Get`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue", "privateEndpointConnectionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateEndpointConnectionsClient.ListByPolicyName`

```go
ctx := context.TODO()
id := privateendpointconnections.NewPrivateLinkForAzureAdID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateLinkForAzureAdValue")

// alternatively `client.ListByPolicyName(ctx, id)` can be used to do batched pagination
items, err := client.ListByPolicyNameComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
