
## `github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2022-11-01/mhsmlistprivateendpointconnections` Documentation

The `mhsmlistprivateendpointconnections` SDK allows for interaction with the Azure Resource Manager Service `keyvault` (API Version `2022-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/keyvault/2022-11-01/mhsmlistprivateendpointconnections"
```


### Client Initialization

```go
client := mhsmlistprivateendpointconnections.NewMHSMListPrivateEndpointConnectionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MHSMListPrivateEndpointConnectionsClient.MHSMPrivateEndpointConnectionsListByResource`

```go
ctx := context.TODO()
id := mhsmlistprivateendpointconnections.NewManagedHSMID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedHSMValue")

// alternatively `client.MHSMPrivateEndpointConnectionsListByResource(ctx, id)` can be used to do batched pagination
items, err := client.MHSMPrivateEndpointConnectionsListByResourceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
