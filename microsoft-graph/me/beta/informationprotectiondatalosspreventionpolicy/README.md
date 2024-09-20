
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectiondatalosspreventionpolicy` Documentation

The `informationprotectiondatalosspreventionpolicy` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := informationprotectiondatalosspreventionpolicy.NewInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.CreateInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()

payload := informationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateInformationProtectionDataLossPreventionPolicy(ctx, payload, informationprotectiondatalosspreventionpolicy.DefaultCreateInformationProtectionDataLossPreventionPolicyOperationOptions())
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
id := informationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyId")

read, err := client.DeleteInformationProtectionDataLossPreventionPolicy(ctx, id, informationprotectiondatalosspreventionpolicy.DefaultDeleteInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.EvaluateInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()

payload := informationprotectiondatalosspreventionpolicy.EvaluateInformationProtectionDataLossPreventionPoliciesRequest{
	// ...
}


read, err := client.EvaluateInformationProtectionDataLossPreventionPolicies(ctx, payload, informationprotectiondatalosspreventionpolicy.DefaultEvaluateInformationProtectionDataLossPreventionPoliciesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionDataLossPreventionPolicyClient.GetInformationProtectionDataLossPreventionPoliciesCount`

```go
ctx := context.TODO()


read, err := client.GetInformationProtectionDataLossPreventionPoliciesCount(ctx, informationprotectiondatalosspreventionpolicy.DefaultGetInformationProtectionDataLossPreventionPoliciesCountOperationOptions())
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
id := informationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyId")

read, err := client.GetInformationProtectionDataLossPreventionPolicy(ctx, id, informationprotectiondatalosspreventionpolicy.DefaultGetInformationProtectionDataLossPreventionPolicyOperationOptions())
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


// alternatively `client.ListInformationProtectionDataLossPreventionPolicies(ctx, informationprotectiondatalosspreventionpolicy.DefaultListInformationProtectionDataLossPreventionPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionDataLossPreventionPoliciesComplete(ctx, informationprotectiondatalosspreventionpolicy.DefaultListInformationProtectionDataLossPreventionPoliciesOperationOptions())
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
id := informationprotectiondatalosspreventionpolicy.NewMeInformationProtectionDataLossPreventionPolicyID("dataLossPreventionPolicyId")

payload := informationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateInformationProtectionDataLossPreventionPolicy(ctx, id, payload, informationprotectiondatalosspreventionpolicy.DefaultUpdateInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
