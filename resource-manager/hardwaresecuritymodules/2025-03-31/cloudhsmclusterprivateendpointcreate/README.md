
## `github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointcreate` Documentation

The `cloudhsmclusterprivateendpointcreate` SDK allows for interaction with Azure Resource Manager `hardwaresecuritymodules` (API Version `2025-03-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hardwaresecuritymodules/2025-03-31/cloudhsmclusterprivateendpointcreate"
```


### Client Initialization

```go
client := cloudhsmclusterprivateendpointcreate.NewCloudHSMClusterPrivateEndpointCreateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CloudHSMClusterPrivateEndpointCreateClient.CloudHsmClusterPrivateEndpointConnectionsCreate`

```go
ctx := context.TODO()
id := cloudhsmclusterprivateendpointcreate.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudHsmClusterName", "privateEndpointConnectionName")

payload := cloudhsmclusterprivateendpointcreate.PrivateEndpointConnection{
	// ...
}


read, err := client.CloudHsmClusterPrivateEndpointConnectionsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
