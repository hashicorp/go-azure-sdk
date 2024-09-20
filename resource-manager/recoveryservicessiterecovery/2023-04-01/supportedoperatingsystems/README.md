
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-04-01/supportedoperatingsystems` Documentation

The `supportedoperatingsystems` SDK allows for interaction with Azure Resource Manager `recoveryservicessiterecovery` (API Version `2023-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicessiterecovery/2023-04-01/supportedoperatingsystems"
```


### Client Initialization

```go
client := supportedoperatingsystems.NewSupportedOperatingSystemsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SupportedOperatingSystemsClient.Get`

```go
ctx := context.TODO()
id := supportedoperatingsystems.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceName")

read, err := client.Get(ctx, id, supportedoperatingsystems.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
