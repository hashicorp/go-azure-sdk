
## `github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/credits` Documentation

The `credits` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2023-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2023-11-01/credits"
```


### Client Initialization

```go
client := credits.NewCreditsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CreditsClient.Get`

```go
ctx := context.TODO()
id := credits.NewBillingProfileID("billingAccountIdValue", "billingProfileIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
