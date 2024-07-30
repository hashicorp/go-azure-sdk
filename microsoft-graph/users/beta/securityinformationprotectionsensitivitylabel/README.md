
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabel` Documentation

The `securityinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := securityinformationprotectionsensitivitylabel.NewSecurityInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := securityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.DeleteSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteSecurityInformationProtectionSensitivityLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.GetSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.GetSecurityInformationProtectionSensitivityLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.GetSecurityInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

read, err := client.GetSecurityInformationProtectionSensitivityLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.ListSecurityInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userIdValue")

// alternatively `client.ListSecurityInformationProtectionSensitivityLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.UpdateSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := securityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.UpdateSecurityInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
