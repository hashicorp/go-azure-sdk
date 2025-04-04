
## `github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineextensionssetup` Documentation

The `machineextensionssetup` SDK allows for interaction with Azure Resource Manager `hybridcompute` (API Version `2025-01-13`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2025-01-13/machineextensionssetup"
```


### Client Initialization

```go
client := machineextensionssetup.NewMachineExtensionsSetupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MachineExtensionsSetupClient.SetupExtensions`

```go
ctx := context.TODO()
id := machineextensionssetup.NewMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "machineName")

payload := machineextensionssetup.SetupExtensionRequest{
	// ...
}


if err := client.SetupExtensionsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
