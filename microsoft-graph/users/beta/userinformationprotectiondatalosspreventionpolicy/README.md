
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectiondatalosspreventionpolicy` Documentation

The `userinformationprotectiondatalosspreventionpolicy` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := userinformationprotectiondatalosspreventionpolicy.NewUserInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.CreateUserByIdInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

payload := userinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.CreateUserByIdInformationProtectionDataLossPreventionPolicyEvaluate`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

payload := userinformationprotectiondatalosspreventionpolicy.CreateUserByIdInformationProtectionDataLossPreventionPolicyEvaluateRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionDataLossPreventionPolicyEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.DeleteUserByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.DeleteUserByIdInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.GetUserByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.GetUserByIdInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.GetUserByIdInformationProtectionDataLossPreventionPolicyCount`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

read, err := client.GetUserByIdInformationProtectionDataLossPreventionPolicyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.ListUserByIdInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserID("userIdValue")

// alternatively `client.ListUserByIdInformationProtectionDataLossPreventionPolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdInformationProtectionDataLossPreventionPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserInformationProtectionDataLossPreventionPolicyClient.UpdateUserByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := userinformationprotectiondatalosspreventionpolicy.NewUserInformationProtectionDataLossPreventionPolicyID("userIdValue", "dataLossPreventionPolicyIdValue")

payload := userinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateUserByIdInformationProtectionDataLossPreventionPolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
