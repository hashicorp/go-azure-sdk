
## `github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/sqlvirtualmachinetroubleshoot` Documentation

The `sqlvirtualmachinetroubleshoot` SDK allows for interaction with the Azure Resource Manager Service `sqlvirtualmachine` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sqlvirtualmachine/2023-10-01/sqlvirtualmachinetroubleshoot"
```


### Client Initialization

```go
client := sqlvirtualmachinetroubleshoot.NewSqlVirtualMachineTroubleshootClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SqlVirtualMachineTroubleshootClient.Troubleshoot`

```go
ctx := context.TODO()
id := sqlvirtualmachinetroubleshoot.NewSqlVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sqlVirtualMachineValue")

payload := sqlvirtualmachinetroubleshoot.SqlVMTroubleshooting{
	// ...
}


if err := client.TroubleshootThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
