
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mesecurityinformationprotectionsensitivitylabel` Documentation

The `mesecurityinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/mesecurityinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := mesecurityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.CreateMeSecurityInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()

payload := mesecurityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.CreateMeSecurityInformationProtectionSensitivityLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication`

```go
ctx := context.TODO()

payload := mesecurityinformationprotectionsensitivitylabel.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateApplication(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult`

```go
ctx := context.TODO()

payload := mesecurityinformationprotectionsensitivitylabel.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateClassificationResult(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval`

```go
ctx := context.TODO()

payload := mesecurityinformationprotectionsensitivitylabel.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateMeSecurityInformationProtectionSensitivityLabelSecurityEvaluateRemoval(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.CreateMeSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel`

```go
ctx := context.TODO()

payload := mesecurityinformationprotectionsensitivitylabel.CreateMeSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabelRequest{
	// ...
}


read, err := client.CreateMeSecurityInformationProtectionSensitivityLabelSecurityExtractContentLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.DeleteMeSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := mesecurityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.DeleteMeSecurityInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.GetMeSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := mesecurityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.GetMeSecurityInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.GetMeSecurityInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()


read, err := client.GetMeSecurityInformationProtectionSensitivityLabelCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.ListMeSecurityInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()


// alternatively `client.ListMeSecurityInformationProtectionSensitivityLabels(ctx)` can be used to do batched pagination
items, err := client.ListMeSecurityInformationProtectionSensitivityLabelsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeSecurityInformationProtectionSensitivityLabelClient.UpdateMeSecurityInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := mesecurityinformationprotectionsensitivitylabel.NewMeSecurityInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := mesecurityinformationprotectionsensitivitylabel.SecuritySensitivityLabel{
	// ...
}


read, err := client.UpdateMeSecurityInformationProtectionSensitivityLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
