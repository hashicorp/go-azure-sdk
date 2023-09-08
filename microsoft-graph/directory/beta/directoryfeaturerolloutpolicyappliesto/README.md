
## `github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryfeaturerolloutpolicyappliesto` Documentation

The `directoryfeaturerolloutpolicyappliesto` SDK allows for interaction with the Azure Resource Manager Service `directory` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/directory/beta/directoryfeaturerolloutpolicyappliesto"
```


### Client Initialization

```go
client := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyAppliesToClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.AddDirectoryFeatureRolloutPolicyByIdAppliesToRef`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := directoryfeaturerolloutpolicyappliesto.ReferenceCreate{
	// ...
}


read, err := client.AddDirectoryFeatureRolloutPolicyByIdAppliesToRef(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.CreateDirectoryFeatureRolloutPolicyByIdAppliesTo`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := directoryfeaturerolloutpolicyappliesto.DirectoryObject{
	// ...
}


read, err := client.CreateDirectoryFeatureRolloutPolicyByIdAppliesTo(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.GetDirectoryFeatureRolloutPolicyByIdAppliesToByIds`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.GetDirectoryFeatureRolloutPolicyByIdAppliesToByIds(ctx, id)` can be used to do batched pagination
items, err := client.GetDirectoryFeatureRolloutPolicyByIdAppliesToByIdsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.GetDirectoryFeatureRolloutPolicyByIdAppliesToCount`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

read, err := client.GetDirectoryFeatureRolloutPolicyByIdAppliesToCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.GetDirectoryFeatureRolloutPolicyByIdAppliesToUserOwnedObject`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := directoryfeaturerolloutpolicyappliesto.GetDirectoryFeatureRolloutPolicyByIdAppliesToUserOwnedObjectRequest{
	// ...
}


read, err := client.GetDirectoryFeatureRolloutPolicyByIdAppliesToUserOwnedObject(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.ListDirectoryFeatureRolloutPolicyByIdAppliesToRefs`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.ListDirectoryFeatureRolloutPolicyByIdAppliesToRefs(ctx, id)` can be used to do batched pagination
items, err := client.ListDirectoryFeatureRolloutPolicyByIdAppliesToRefsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.ListDirectoryFeatureRolloutPolicyByIdAppliesTos`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

// alternatively `client.ListDirectoryFeatureRolloutPolicyByIdAppliesTos(ctx, id)` can be used to do batched pagination
items, err := client.ListDirectoryFeatureRolloutPolicyByIdAppliesTosComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.RemoveDirectoryFeatureRolloutPolicyByIdAppliesToByIdRef`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyAppliesToID("featureRolloutPolicyIdValue", "directoryObjectIdValue")

read, err := client.RemoveDirectoryFeatureRolloutPolicyByIdAppliesToByIdRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryFeatureRolloutPolicyAppliesToClient.ValidateDirectoryFeatureRolloutPolicyByIdAppliesToProperty`

```go
ctx := context.TODO()
id := directoryfeaturerolloutpolicyappliesto.NewDirectoryFeatureRolloutPolicyID("featureRolloutPolicyIdValue")

payload := directoryfeaturerolloutpolicyappliesto.ValidateDirectoryFeatureRolloutPolicyByIdAppliesToPropertyRequest{
	// ...
}


read, err := client.ValidateDirectoryFeatureRolloutPolicyByIdAppliesToProperty(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
