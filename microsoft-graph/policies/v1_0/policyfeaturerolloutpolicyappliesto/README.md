
## `github.com/hashicorp/go-azure-sdk/resource-manager/policies/v1.0/policyfeaturerolloutpolicyappliesto` Documentation

The `policyfeaturerolloutpolicyappliesto` SDK allows for interaction with the Azure Resource Manager Service `policies` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/policies/v1.0/policyfeaturerolloutpolicyappliesto"
```


### Client Initialization

```go
client := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyAppliesToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.AddPolicyFeatureRolloutPolicyByIdAppliesToRef`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := policyfeaturerolloutpolicyappliesto.ReferenceCreate{
	// ...
}


read, err := client.AddPolicyFeatureRolloutPolicyByIdAppliesToRef(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.CreatePolicyFeatureRolloutPolicyByIdAppliesTo`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := policyfeaturerolloutpolicyappliesto.DirectoryObject{
	// ...
}


read, err := client.CreatePolicyFeatureRolloutPolicyByIdAppliesTo(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.GetPolicyFeatureRolloutPolicyByIdAppliesToAvailableExtensionProperties`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.GetPolicyFeatureRolloutPolicyByIdAppliesToAvailableExtensionProperties(ctx, id)` can be used to do batched pagination
items, err := client.GetPolicyFeatureRolloutPolicyByIdAppliesToAvailableExtensionPropertiesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.GetPolicyFeatureRolloutPolicyByIdAppliesToByIds`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.GetPolicyFeatureRolloutPolicyByIdAppliesToByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetPolicyFeatureRolloutPolicyByIdAppliesToByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.GetPolicyFeatureRolloutPolicyByIdAppliesToCount`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

read, err := client.GetPolicyFeatureRolloutPolicyByIdAppliesToCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.ListPolicyFeatureRolloutPolicyByIdAppliesToRefs`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.ListPolicyFeatureRolloutPolicyByIdAppliesToRefs(ctx, id)` can be used to do batched pagination
items, err := client.ListPolicyFeatureRolloutPolicyByIdAppliesToRefsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.ListPolicyFeatureRolloutPolicyByIdAppliesTos`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.ListPolicyFeatureRolloutPolicyByIdAppliesTos(ctx, id)` can be used to do batched pagination
items, err := client.ListPolicyFeatureRolloutPolicyByIdAppliesTosComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.RemovePolicyFeatureRolloutPolicyByIdAppliesToByIdRef`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue")

read, err := client.RemovePolicyFeatureRolloutPolicyByIdAppliesToByIdRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PolicyFeatureRolloutPolicyAppliesToClient.ValidatePolicyFeatureRolloutPolicyByIdAppliesToProperty`

```go
ctx := context.TODO()
id := policyfeaturerolloutpolicyappliesto.NewPolicyFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := policyfeaturerolloutpolicyappliesto.ValidatePolicyFeatureRolloutPolicyByIdAppliesToPropertyRequest{
	// ...
}


read, err := client.ValidatePolicyFeatureRolloutPolicyByIdAppliesToProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
