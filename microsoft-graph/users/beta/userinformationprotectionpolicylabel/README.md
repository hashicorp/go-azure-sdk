
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionpolicylabel` Documentation

The `userinformationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := userinformationprotectionpolicylabel.NewUserInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.CreateUserByIdInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

payload := userinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.CreateUserByIdInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

payload := userinformationprotectionpolicylabel.CreateUserByIdInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionPolicyLabelEvaluateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.CreateUserByIdInformationProtectionPolicyLabelEvaluateClassificationResult`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

payload := userinformationprotectionpolicylabel.CreateUserByIdInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.CreateUserByIdInformationProtectionPolicyLabelEvaluateRemoval`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

payload := userinformationprotectionpolicylabel.CreateUserByIdInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionPolicyLabelEvaluateRemoval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.CreateUserByIdInformationProtectionPolicyLabelExtractLabel`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

payload := userinformationprotectionpolicylabel.CreateUserByIdInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionPolicyLabelExtractLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.DeleteUserByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

read, err := client.DeleteUserByIdInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.GetUserByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

read, err := client.GetUserByIdInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.GetUserByIdInformationProtectionPolicyLabelCount`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

read, err := client.GetUserByIdInformationProtectionPolicyLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.ListUserByIdInformationProtectionPolicyLabels`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserID("userIdValue")

// alternatively `client.ListUserByIdInformationProtectionPolicyLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdInformationProtectionPolicyLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserInformationProtectionPolicyLabelClient.UpdateUserByIdInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionpolicylabel.NewUserInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

payload := userinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateUserByIdInformationProtectionPolicyLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
