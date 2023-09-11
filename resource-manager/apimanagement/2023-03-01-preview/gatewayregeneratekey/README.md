
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/gatewayregeneratekey` Documentation

The `gatewayregeneratekey` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2023-03-01-preview/gatewayregeneratekey"
```


### Client Initialization

```go
client := gatewayregeneratekey.NewGatewayRegenerateKeyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatewayRegenerateKeyClient.GatewayRegenerateKey`

```go
ctx := context.TODO()
id := gatewayregeneratekey.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue")

payload := gatewayregeneratekey.GatewayKeyRegenerationRequestContract{
	// ...
}


read, err := client.GatewayRegenerateKey(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
