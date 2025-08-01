
## `github.com/hashicorp/go-azure-sdk/resource-manager/storage/2025-01-01/objectreplicationpolicies` Documentation

The `objectreplicationpolicies` SDK allows for interaction with Azure Resource Manager `storage` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/storage/2025-01-01/objectreplicationpolicies"
```


### Client Initialization

```go
client := objectreplicationpolicies.NewObjectReplicationPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ObjectReplicationPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := objectreplicationpolicies.NewObjectReplicationPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "objectReplicationPolicyId")

payload := objectreplicationpolicies.ObjectReplicationPolicy{
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


### Example Usage: `ObjectReplicationPoliciesClient.Delete`

```go
ctx := context.TODO()
id := objectreplicationpolicies.NewObjectReplicationPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "objectReplicationPolicyId")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ObjectReplicationPoliciesClient.Get`

```go
ctx := context.TODO()
id := objectreplicationpolicies.NewObjectReplicationPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName", "objectReplicationPolicyId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ObjectReplicationPoliciesClient.List`

```go
ctx := context.TODO()
id := commonids.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageAccountName")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
