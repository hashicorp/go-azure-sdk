
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionsensitivitylabel` Documentation

The `userinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/userinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := userinformationprotectionsensitivitylabel.NewUserInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.CreateUserByIdInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := userinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.CreateUserByIdInformationProtectionSensitivityLabelEvaluate`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := userinformationprotectionsensitivitylabel.CreateUserByIdInformationProtectionSensitivityLabelEvaluateRequest{
	// ...
}


read, err := client.CreateUserByIdInformationProtectionSensitivityLabelEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.DeleteUserByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteUserByIdInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.GetUserByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.GetUserByIdInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.GetUserByIdInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserID("userIdValue")

read, err := client.GetUserByIdInformationProtectionSensitivityLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.ListUserByIdInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserID("userIdValue")

// alternatively `client.ListUserByIdInformationProtectionSensitivityLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdInformationProtectionSensitivityLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserInformationProtectionSensitivityLabelClient.UpdateUserByIdInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := userinformationprotectionsensitivitylabel.NewUserInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := userinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateUserByIdInformationProtectionSensitivityLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
