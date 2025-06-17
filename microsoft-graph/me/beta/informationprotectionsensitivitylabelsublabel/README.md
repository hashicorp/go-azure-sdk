
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabelsublabel` Documentation

The `informationprotectionsensitivitylabelsublabel` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := informationprotectionsensitivitylabelsublabel.NewInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.CreateInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

payload := informationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelSublabel(ctx, id, payload, informationprotectionsensitivitylabelsublabel.DefaultCreateInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.CreateInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritance`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

payload := informationprotectionsensitivitylabelsublabel.CreateInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceRequest{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritance(ctx, id, payload, informationprotectionsensitivitylabelsublabel.DefaultCreateInformationProtectionSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.DeleteInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1")

read, err := client.DeleteInformationProtectionSensitivityLabelSublabel(ctx, id, informationprotectionsensitivitylabelsublabel.DefaultDeleteInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.EvaluateInformationProtectionSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

payload := informationprotectionsensitivitylabelsublabel.EvaluateInformationProtectionSensitivityLabelSublabelsRequest{
	// ...
}


read, err := client.EvaluateInformationProtectionSensitivityLabelSublabels(ctx, id, payload, informationprotectionsensitivitylabelsublabel.DefaultEvaluateInformationProtectionSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.GetInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1")

read, err := client.GetInformationProtectionSensitivityLabelSublabel(ctx, id, informationprotectionsensitivitylabelsublabel.DefaultGetInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.GetInformationProtectionSensitivityLabelSublabelsCount`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

read, err := client.GetInformationProtectionSensitivityLabelSublabelsCount(ctx, id, informationprotectionsensitivitylabelsublabel.DefaultGetInformationProtectionSensitivityLabelSublabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.ListInformationProtectionSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

// alternatively `client.ListInformationProtectionSensitivityLabelSublabels(ctx, id, informationprotectionsensitivitylabelsublabel.DefaultListInformationProtectionSensitivityLabelSublabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionSensitivityLabelSublabelsComplete(ctx, id, informationprotectionsensitivitylabelsublabel.DefaultListInformationProtectionSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.UpdateInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelId", "sensitivityLabelId1")

payload := informationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateInformationProtectionSensitivityLabelSublabel(ctx, id, payload, informationprotectionsensitivitylabelsublabel.DefaultUpdateInformationProtectionSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
