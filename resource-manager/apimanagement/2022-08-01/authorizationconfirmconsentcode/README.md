
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationconfirmconsentcode` Documentation

The `authorizationconfirmconsentcode` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/authorizationconfirmconsentcode"
```


### Client Initialization

```go
client := authorizationconfirmconsentcode.NewAuthorizationConfirmConsentCodeClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AuthorizationConfirmConsentCodeClient.AuthorizationConfirmConsentCode`

```go
ctx := context.TODO()
id := authorizationconfirmconsentcode.NewAuthorizationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "authorizationProviderIdValue", "authorizationIdValue")

payload := authorizationconfirmconsentcode.AuthorizationConfirmConsentCodeRequestContract{
	// ...
}


read, err := client.AuthorizationConfirmConsentCode(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
