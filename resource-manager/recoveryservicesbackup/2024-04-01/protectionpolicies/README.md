
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/protectionpolicies` Documentation

The `protectionpolicies` SDK allows for interaction with the Azure Resource Manager Service `recoveryservicesbackup` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservicesbackup/2024-04-01/protectionpolicies"
```


### Client Initialization

```go
client := protectionpolicies.NewProtectionPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ProtectionPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := protectionpolicies.NewBackupPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupPolicyValue")

payload := protectionpolicies.ProtectionPolicyResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, protectionpolicies.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ProtectionPoliciesClient.Delete`

```go
ctx := context.TODO()
id := protectionpolicies.NewBackupPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupPolicyValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ProtectionPoliciesClient.Get`

```go
ctx := context.TODO()
id := protectionpolicies.NewBackupPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue", "backupPolicyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
