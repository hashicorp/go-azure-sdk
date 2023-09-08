
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionsensitivitylabelsublabel` Documentation

The `meinformationprotectionsensitivitylabelsublabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.CreateMeInformationProtectionSensitivityLabelByIdSublabel`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := meinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateMeInformationProtectionSensitivityLabelByIdSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.CreateMeInformationProtectionSensitivityLabelByIdSublabelEvaluate`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := meinformationprotectionsensitivitylabelsublabel.CreateMeInformationProtectionSensitivityLabelByIdSublabelEvaluateRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionSensitivityLabelByIdSublabelEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.DeleteMeInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.DeleteMeInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.GetMeInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.GetMeInformationProtectionSensitivityLabelByIdSublabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.GetMeInformationProtectionSensitivityLabelByIdSublabelCount`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.GetMeInformationProtectionSensitivityLabelByIdSublabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.ListMeInformationProtectionSensitivityLabelByIdSublabels`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

// alternatively `client.ListMeInformationProtectionSensitivityLabelByIdSublabels(ctx, id)` can be used to do batched pagination
items, err := client.ListMeInformationProtectionSensitivityLabelByIdSublabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelSublabelClient.UpdateMeInformationProtectionSensitivityLabelByIdSublabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

payload := meinformationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateMeInformationProtectionSensitivityLabelByIdSublabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
