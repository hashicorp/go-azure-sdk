
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionsensitivitylabelsublabel` Documentation

The `userinformationprotectionsensitivitylabelsublabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.CreateUserByIdInformationProtectionSensitivityLabelByIdSublabel`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := userinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionSensitivityLabelByIdSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.CreateUserByIdInformationProtectionSensitivityLabelByIdSublabelEvaluate`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := userinformationprotectionsensitivitylabelsublabel.CreateUserByIdInformationProtectionSensitivityLabelByIdSublabelEvaluateRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionSensitivityLabelByIdSublabelEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.DeleteUserByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelSublabelID("userIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.DeleteUserByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.GetUserByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelSublabelID("userIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.GetUserByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.GetUserByIdInformationProtectionSensitivityLabelByIdSublabelCount`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.GetUserByIdInformationProtectionSensitivityLabelByIdSublabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.ListUserByIdInformationProtectionSensitivityLabelByIdSublabels`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

// alternatively `client.ListUserByIdInformationProtectionSensitivityLabelByIdSublabels(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdInformationProtectionSensitivityLabelByIdSublabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelSublabelClient.UpdateUserByIdInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabelsublabel.NewUserInformationProtectionSensitivityLabelSublabelID("userIdValue", "sensitivityLabelIdValue", "sensitivityLabelId1Value")

payload := userinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateUserByIdInformationProtectionSensitivityLabelByIdSublabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
