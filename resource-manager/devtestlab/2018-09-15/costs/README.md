
## `github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/costs` Documentation

The `costs` SDK allows for interaction with Azure Resource Manager `devtestlab` (API Version `2018-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/costs"
```


### Client Initialization

```go
client := costs.NewCostsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CostsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := costs.NewCostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "name")

payload := costs.LabCost{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CostsClient.Get`

```go
ctx := context.TODO()
id := costs.NewCostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "labName", "name")

read, err := client.Get(ctx, id, costs.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
