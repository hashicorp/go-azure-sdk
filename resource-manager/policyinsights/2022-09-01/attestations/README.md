
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2022-09-01/attestations` Documentation

The `attestations` SDK allows for interaction with Azure Resource Manager `policyinsights` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2022-09-01/attestations"
```


### Client Initialization

```go
client := attestations.NewAttestationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AttestationsClient.CreateOrUpdateAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.CreateOrUpdateAtResourceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.CreateOrUpdateAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.CreateOrUpdateAtResourceGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.CreateOrUpdateAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.CreateOrUpdateAtSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.DeleteAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

read, err := client.DeleteAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.DeleteAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

read, err := client.DeleteAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.DeleteAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

read, err := client.DeleteAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.GetAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

read, err := client.GetAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.GetAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

read, err := client.GetAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.GetAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

read, err := client.GetAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.ListForResource`

```go
ctx := context.TODO()
id := commonids.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.ListForResource(ctx, id, attestations.DefaultListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceComplete(ctx, id, attestations.DefaultListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AttestationsClient.ListForResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListForResourceGroup(ctx, id, attestations.DefaultListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ListForResourceGroupComplete(ctx, id, attestations.DefaultListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AttestationsClient.ListForSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListForSubscription(ctx, id, attestations.DefaultListForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ListForSubscriptionComplete(ctx, id, attestations.DefaultListForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
