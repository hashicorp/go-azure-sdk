
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewaygeneratetoken` Documentation

The `gatewaygeneratetoken` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/gatewaygeneratetoken"
```


### Client Initialization

```go
client := gatewaygeneratetoken.NewGatewayGenerateTokenClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GatewayGenerateTokenClient.GatewayGenerateToken`

```go
ctx := context.TODO()
id := gatewaygeneratetoken.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "gatewayIdValue")

payload := gatewaygeneratetoken.GatewayTokenRequestContract{
	// ...
}


read, err := client.GatewayGenerateToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
