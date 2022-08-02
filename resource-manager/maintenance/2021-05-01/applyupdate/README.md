
## `github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/applyupdate` Documentation

The `applyupdate` SDK allows for interaction with the Azure Resource Manager Service `maintenance` (API Version `2021-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/applyupdate"
```


### Client Initialization

```go
client := applyupdate.NewApplyUpdateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplyUpdateClient.ForResourceGroupList`

```go
ctx := context.TODO()
id := applyupdate.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ForResourceGroupList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplyUpdateClient.List`

```go
ctx := context.TODO()
id := applyupdate.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
