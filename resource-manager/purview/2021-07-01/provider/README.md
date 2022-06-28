
## `github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/provider` Documentation

The `provider` SDK allows for interaction with the Azure Resource Manager Service `purview` (API Version `2021-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/purview/2021-07-01/provider"
```


### Client Initialization

```go
client := provider.NewProviderClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ProviderClient.AccountsCheckNameAvailability`

```go
ctx := context.TODO()
id := provider.NewSubscriptionID()

payload := provider.CheckNameAvailabilityRequest{
	// ...
}

read, err := client.AccountsCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
