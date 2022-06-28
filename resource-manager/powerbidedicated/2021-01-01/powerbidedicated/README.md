
## `github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/powerbidedicated` Documentation

The `powerbidedicated` SDK allows for interaction with the Azure Resource Manager Service `powerbidedicated` (API Version `2021-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/powerbidedicated/2021-01-01/powerbidedicated"
```


### Client Initialization

```go
client := powerbidedicated.NewPowerBIDedicatedClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PowerBIDedicatedClient.CapacitiesListSkus`

```go
ctx := context.TODO()
id := powerbidedicated.NewSubscriptionID()

read, err := client.CapacitiesListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
