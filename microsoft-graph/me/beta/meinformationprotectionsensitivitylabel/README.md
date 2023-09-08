
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionsensitivitylabel` Documentation

The `meinformationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := meinformationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.CreateMeInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()

payload := meinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateMeInformationProtectionSensitivityLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.CreateMeInformationProtectionSensitivityLabelEvaluate`

```go
ctx := context.TODO()

payload := meinformationprotectionsensitivitylabel.CreateMeInformationProtectionSensitivityLabelEvaluateRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionSensitivityLabelEvaluate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.DeleteMeInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.DeleteMeInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.GetMeInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

read, err := client.GetMeInformationProtectionSensitivityLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.GetMeInformationProtectionSensitivityLabelCount`

```go
ctx := context.TODO()


read, err := client.GetMeInformationProtectionSensitivityLabelCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.ListMeInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()


// alternatively `client.ListMeInformationProtectionSensitivityLabels(ctx)` can be used to do batched pagination
items, err := client.ListMeInformationProtectionSensitivityLabelsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeInformationProtectionSensitivityLabelClient.UpdateMeInformationProtectionSensitivityLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelIdValue")

payload := meinformationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateMeInformationProtectionSensitivityLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
