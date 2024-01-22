
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/address` Documentation

The `address` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/address"
```


### Client Initialization

```go
client := address.NewAddressClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AddressClient.Validate`

```go
ctx := context.TODO()

payload := address.AddressDetails{
	// ...
}


read, err := client.Validate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
