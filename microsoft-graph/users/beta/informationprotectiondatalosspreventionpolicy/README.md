
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectiondatalosspreventionpolicy` Documentation

The `informationprotectiondatalosspreventionpolicy` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := informationprotectiondatalosspreventionpolicy.NewInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.CreateInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

payload := informationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.CreateInformationProtectionDataLossPreventionPolicyEvaluate`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

payload := informationprotectiondatalosspreventionpolicy.CreateInformationProtectionDataLossPreventionPolicyEvaluateRequest{
	// ...
}


read, err := client.CreateInformationProtectionDataLossPreventionPolicyEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.DeleteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserIdInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.DeleteInformationProtectionDataLossPreventionPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.GetInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserIdInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.GetInformationProtectionDataLossPreventionPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.GetInformationProtectionDataLossPreventionPolicyCount`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

read, err := client.GetInformationProtectionDataLossPreventionPolicyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.ListInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

// alternatively `client.ListInformationProtectionDataLossPreventionPolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListInformationProtectionDataLossPreventionPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.UpdateInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := informationprotectiondatalosspreventionpolicy.NewUserIdInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

payload := informationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
