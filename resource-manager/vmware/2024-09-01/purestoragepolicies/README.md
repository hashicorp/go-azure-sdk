
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/purestoragepolicies` Documentation

The `purestoragepolicies` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/purestoragepolicies"
```


### Client Initialization

```go
client := purestoragepolicies.NewPureStoragePoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PureStoragePoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := purestoragepolicies.NewPureStoragePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "pureStoragePolicyName")

payload := purestoragepolicies.PureStoragePolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `PureStoragePoliciesClient.Delete`

```go
ctx := context.TODO()
id := purestoragepolicies.NewPureStoragePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "pureStoragePolicyName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PureStoragePoliciesClient.Get`

```go
ctx := context.TODO()
id := purestoragepolicies.NewPureStoragePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "pureStoragePolicyName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PureStoragePoliciesClient.List`

```go
ctx := context.TODO()
id := purestoragepolicies.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
