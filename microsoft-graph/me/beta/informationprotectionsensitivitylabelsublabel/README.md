
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabelsublabel` Documentation

The `informationprotectionsensitivitylabelsublabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabelsublabel"
```


### Client Initialization

```go
client := informationprotectionsensitivitylabelsublabel.NewInformationProtectionSensitivityLabelSublabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.CreateInformationProtectionSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := informationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.CreateInformationProtectionSensitivityLabelSublabelEvaluate`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := informationprotectionsensitivitylabelsublabel.CreateInformationProtectionSensitivityLabelSublabelEvaluateRequest{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelSublabelEvaluate(ctx, id, payload)
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
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.DeleteInformationProtectionSensitivityLabelSublabel(ctx, id)
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
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

read, err := client.GetInformationProtectionSensitivityLabelSublabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelSublabelClient.GetInformationProtectionSensitivityLabelSublabelCount`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.GetInformationProtectionSensitivityLabelSublabelCount(ctx, id)
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
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

// alternatively `client.ListInformationProtectionSensitivityLabelSublabels(ctx, id)` can be used to do batched pagination
items, err := client.ListInformationProtectionSensitivityLabelSublabelsComplete(ctx, id)
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
id := informationprotectionsensitivitylabelsublabel.NewMeInformationProtectionSensitivityLabelIdSublabelID("sensitivityLabelIdValue", "sensitivityLabelId1Value")

payload := informationprotectionsensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateInformationProtectionSensitivityLabelSublabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
