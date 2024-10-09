
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/virtualmachinesizes` Documentation

The `virtualmachinesizes` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2022-03-01/virtualmachinesizes"
```


### Client Initialization

```go
client := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachineSizesClient.List`

```go
ctx := context.TODO()
id := virtualmachinesizes.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
