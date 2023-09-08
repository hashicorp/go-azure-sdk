
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectiondatalosspreventionpolicy` Documentation

The `meinformationprotectiondatalosspreventionpolicy` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := meinformationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.CreateMeInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()

payload := meinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateMeInformationProtectionDataLossPreventionPolicy(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.CreateMeInformationProtectionDataLossPreventionPolicyEvaluate`

```go
ctx := context.TODO()

payload := meinformationprotectiondatalosspreventionpolicy.CreateMeInformationProtectionDataLossPreventionPolicyEvaluateRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionDataLossPreventionPolicyEvaluate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.DeleteMeInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := meinformationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyIdValue")

read, err := client.DeleteMeInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.GetMeInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := meinformationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyIdValue")

read, err := client.GetMeInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.GetMeInformationProtectionDataLossPreventionPolicyCount`

```go
ctx := context.TODO()


read, err := client.GetMeInformationProtectionDataLossPreventionPolicyCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.ListMeInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()


// alternatively `client.ListMeInformationProtectionDataLossPreventionPolicies(ctx)` can be used to do batched pagination
items, err := client.ListMeInformationProtectionDataLossPreventionPoliciesComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeInformationProtectionDataLossPreventionPolicyClient.UpdateMeInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := meinformationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyIdValue")

payload := meinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateMeInformationProtectionDataLossPreventionPolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
