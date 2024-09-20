
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabel` Documentation

The `securityinformationprotectionsensitivitylabel` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/securityinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := securityinformationprotectionsensitivitylabel.NewSecurityInformationProtectionSensitivityLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecurityInformationProtectionSensitivityLabelClient.CreateSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

payload := securityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabel(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultCreateSecurityInformationProtectionSensitivityLabelOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

payload := securityinformationprotectionsensitivitylabel.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest{
	// ...
}


read, err := client.CreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultCreateSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userId", "sensitivityLabelId")

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
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userId", "sensitivityLabelId")

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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

read, err := client.GetSecurityInformationProtectionSensitivityLabelsCount(ctx, id, securityinformationprotectionsensitivitylabel.DefaultGetSecurityInformationProtectionSensitivityLabelsCountOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplications(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsComplete(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationsOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResults(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsComplete(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultsOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

payload := securityinformationprotectionsensitivitylabel.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsRequest{
	// ...
}


// alternatively `client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovals(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsComplete(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalsOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserID("userId")

// alternatively `client.ListSecurityInformationProtectionSensitivityLabels(ctx, id, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListSecurityInformationProtectionSensitivityLabelsComplete(ctx, id, securityinformationprotectionsensitivitylabel.DefaultListSecurityInformationProtectionSensitivityLabelsOperationOptions())
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
id := securityinformationprotectionsensitivitylabel.NewUserIdSecurityInformationProtectionSensitivityLabelID("userId", "sensitivityLabelId")

payload := securityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.UpdateSecurityInformationProtectionSensitivityLabel(ctx, id, payload, securityinformationprotectionsensitivitylabel.DefaultUpdateSecurityInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
