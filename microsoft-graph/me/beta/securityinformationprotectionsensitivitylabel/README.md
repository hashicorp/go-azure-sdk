
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

read, err := client.DeleteSecurityInformationProtectionSensitivityLabel(ctx, id, securityinformationprotectionsensitivitylabel.DefaultDeleteSecurityInformationProtectionSensitivityLabelOperationOptions())
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

read, err := client.GetSecurityInformationProtectionSensitivityLabel(ctx, id, securityinformationprotectionsensitivitylabel.DefaultGetSecurityInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.GetSecurityInformationProtectionSensitivityLabelsCount`

```go
ctx := context.TODO()


read, err := client.GetSecurityInformationProtectionSensitivityLabelsCount(ctx, securityinformationprotectionsensitivitylabel.DefaultGetSecurityInformationProtectionSensitivityLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications`

```go
ctx := context.TODO()

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsComplete(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults`

```go
ctx := context.TODO()

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsComplete(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals`

```go
ctx := context.TODO()

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsComplete(ctx, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.ListSecurityInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()


// alternatively `client.ListSecurityInformationProtectionSensitivityLabels(ctx, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelsComplete(ctx, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelsOperationOptions())
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
