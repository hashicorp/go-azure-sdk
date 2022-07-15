
## `github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/vaults` Documentation

The `vaults` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2022-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/netapp/2022-01-01/vaults"
```


### Client Initialization

```go
client := vaults.NewVaultsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VaultsClient.List`

```go
ctx := context.TODO()
id := vaults.NewNetAppAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
