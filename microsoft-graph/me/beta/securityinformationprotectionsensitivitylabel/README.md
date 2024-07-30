
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/securityinformationprotectionsensitivitylabel` Documentation

The `securityinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/securityinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := securityinformationprotectionsensitivitylabel.NewSecurityInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()

payload := securityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabel(ctx, payload)
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

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication(ctx, payload)
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

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult(ctx, payload)
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

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval(ctx, payload)
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

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx, payload)
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
id := securityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

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
id := securityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

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


read, err := client.GetSecurityInformationProtectionSensitivityLabelCount(ctx)
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


// alternatively `client.ListSecurityInformationProtectionSensitivityLabels(ctx)` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelsComplete(ctx)
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
id := securityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

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
