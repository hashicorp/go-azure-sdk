
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/gatewaylistdebugcredentials` Documentation

The `gatewaylistdebugcredentials` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-05-01-preview/gatewaylistdebugcredentials"
```


### Client Initialization

```go
client := gatewaylistdebugcredentials.NewGatewayListDebugCredentialsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatewayListDebugCredentialsClient.GatewayListDebugCredentials`

```go
ctx := context.TODO()
id := gatewaylistdebugcredentials.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue")

payload := gatewaylistdebugcredentials.GatewayListDebugCredentialsContract{
	// ...
}


read, err := client.GatewayListDebugCredentials(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
