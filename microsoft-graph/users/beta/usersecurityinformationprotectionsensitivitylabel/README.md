
## `github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usersecurityinformationprotectionsensitivitylabel` Documentation

The `usersecurityinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/users/beta/usersecurityinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := usersecurityinformationprotectionsensitivitylabel.NewUserSecurityInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.CreateUserByIdSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.CreateUserByIdSecurityInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest{
	// ...
}


read, err := client.CreateUserByIdSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.DeleteUserByIdSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteUserByIdSecurityInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.GetUserByIdSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.GetUserByIdSecurityInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.GetUserByIdSecurityInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

read, err := client.GetUserByIdSecurityInformationProtectionSensitivityLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.ListUserByIdSecurityInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

// alternatively `client.ListUserByIdSecurityInformationProtectionSensitivityLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListUserByIdSecurityInformationProtectionSensitivityLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserSecurityInformationProtectionSensitivityLabelClient.UpdateUserByIdSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := usersecurityinformationprotectionsensitivitylabel.NewUserSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := usersecurityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.UpdateUserByIdSecurityInformationProtectionSensitivityLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
