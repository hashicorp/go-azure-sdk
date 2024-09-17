
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/validateoperation` Documentation

The `validateoperation` SDK allows for interaction with Azure Resource Manager `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/validateoperation"
```


### Client Initialization

```go
client := validateoperation.NewValidateOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ValidateOperationClient.Trigger`

```go
ctx := context.TODO()
id := validateoperation.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

payload := validateoperation.ValidateOperationRequestResource{
	// ...
}


if err := client.TriggerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
