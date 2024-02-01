
## `github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/operations` Documentation

The `operations` SDK allows for interaction with the Azure Resource Manager Service `synapse` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/operations"
```


### Client Initialization

```go
client := operations.NewOperationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := operations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := operations.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
