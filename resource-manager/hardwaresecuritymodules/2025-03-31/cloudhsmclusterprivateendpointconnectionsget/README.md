
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointconnectionsget` Documentation

The `cloudhsmclusterprivateendpointconnectionsget` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointconnectionsget"
```


### Client Initialization

```go
client := cloudhsmclusterprivateendpointconnectionsget.NewCloudHSMClusterPrivateEndpointConnectionsGetClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClusterPrivateEndpointConnectionsGetClient.PrivateEndpointConnectionsListByCloudHsmCluster`

```go
ctx := context.TODO()
id := cloudhsmclusterprivateendpointconnectionsget.NewCloudHsmClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName")

// alternatively `client.PrivateEndpointConnectionsListByCloudHsmCluster(ctx, id)` can be used to do batched pagination
items, err := client.PrivateEndpointConnectionsListByCloudHsmClusterComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
