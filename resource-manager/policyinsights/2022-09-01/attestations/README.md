
## `github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2022-09-01/attestations` Documentation

The `attestations` SDK allows for interaction with the Azure Resource Manager Service `policyinsights` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/policyinsights/2022-09-01/attestations"
```


### Client Initialization

```go
client := attestations.NewAttestationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AttestationsClient.AttestationsCreateOrUpdateAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtResourceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.AttestationsCreateOrUpdateAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtResourceGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.AttestationsCreateOrUpdateAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

payload := attestations.Attestation{
	// ...
}


if err := client.AttestationsCreateOrUpdateAtSubscriptionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `AttestationsClient.AttestationsDeleteAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

read, err := client.AttestationsDeleteAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsDeleteAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

read, err := client.AttestationsDeleteAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsDeleteAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

read, err := client.AttestationsDeleteAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsGetAtResource`

```go
ctx := context.TODO()
id := attestations.NewScopedAttestationID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "attestationValue")

read, err := client.AttestationsGetAtResource(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsGetAtResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewProviderAttestationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "attestationValue")

read, err := client.AttestationsGetAtResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsGetAtSubscription`

```go
ctx := context.TODO()
id := attestations.NewAttestationID("12345678-1234-9876-4563-123456789012", "attestationValue")

read, err := client.AttestationsGetAtSubscription(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationsListForResource`

```go
ctx := context.TODO()
id := attestations.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.AttestationsListForResource(ctx, id, attestations.DefaultAttestationsListForResourceOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForResourceComplete(ctx, id, attestations.DefaultAttestationsListForResourceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AttestationsClient.AttestationsListForResourceGroup`

```go
ctx := context.TODO()
id := attestations.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.AttestationsListForResourceGroup(ctx, id, attestations.DefaultAttestationsListForResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForResourceGroupComplete(ctx, id, attestations.DefaultAttestationsListForResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AttestationsClient.AttestationsListForSubscription`

```go
ctx := context.TODO()
id := attestations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.AttestationsListForSubscription(ctx, id, attestations.DefaultAttestationsListForSubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.AttestationsListForSubscriptionComplete(ctx, id, attestations.DefaultAttestationsListForSubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
