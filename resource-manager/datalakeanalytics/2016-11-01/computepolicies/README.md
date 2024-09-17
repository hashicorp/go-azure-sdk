
## `github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/computepolicies` Documentation

The `computepolicies` SDK allows for interaction with Azure Resource Manager `datalakeanalytics` (API Version `2016-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/datalakeanalytics/2016-11-01/computepolicies"
```


### Client Initialization

```go
client := computepolicies.NewComputePoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ComputePoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := computepolicies.NewComputePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "computePolicyValue")

payload := computepolicies.CreateOrUpdateComputePolicyParameters{
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


### Example Usage: `ComputePoliciesClient.Delete`

```go
ctx := context.TODO()
id := computepolicies.NewComputePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "computePolicyValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputePoliciesClient.Get`

```go
ctx := context.TODO()
id := computepolicies.NewComputePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "computePolicyValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ComputePoliciesClient.ListByAccount`

```go
ctx := context.TODO()
id := computepolicies.NewAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue")

// alternatively `client.ListByAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ComputePoliciesClient.Update`

```go
ctx := context.TODO()
id := computepolicies.NewComputePolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "accountValue", "computePolicyValue")

payload := computepolicies.UpdateComputePolicyParameters{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
