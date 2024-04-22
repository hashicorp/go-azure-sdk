
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-03-01/billingmeters` Documentation

The `billingmeters` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2024-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-03-01/billingmeters"
```


### Client Initialization

```go
client := billingmeters.NewBillingMetersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingMetersClient.Get`

```go
ctx := context.TODO()
id := billingmeters.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
