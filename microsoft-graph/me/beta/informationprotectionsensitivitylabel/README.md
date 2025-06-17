
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabel` Documentation

The `informationprotectionsensitivitylabel` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := informationprotectionsensitivitylabel.NewInformationProtectionSensitivityLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionSensitivityLabelClient.CreateInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()

payload := informationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabel(ctx, payload, informationprotectionsensitivitylabel.DefaultCreateInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.CreateInformationProtectionSensitivityLabelComputeRightsAndInheritance`

```go
ctx := context.TODO()

payload := informationprotectionsensitivitylabel.CreateInformationProtectionSensitivityLabelComputeRightsAndInheritanceRequest{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelComputeRightsAndInheritance(ctx, payload, informationprotectionsensitivitylabel.DefaultCreateInformationProtectionSensitivityLabelComputeRightsAndInheritanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.DeleteInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

read, err := client.DeleteInformationProtectionSensitivityLabel(ctx, id, informationprotectionsensitivitylabel.DefaultDeleteInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.EvaluateInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()

payload := informationprotectionsensitivitylabel.EvaluateInformationProtectionSensitivityLabelsRequest{
	// ...
}


read, err := client.EvaluateInformationProtectionSensitivityLabels(ctx, payload, informationprotectionsensitivitylabel.DefaultEvaluateInformationProtectionSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.GetInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

read, err := client.GetInformationProtectionSensitivityLabel(ctx, id, informationprotectionsensitivitylabel.DefaultGetInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.GetInformationProtectionSensitivityLabelsCount`

```go
ctx := context.TODO()


read, err := client.GetInformationProtectionSensitivityLabelsCount(ctx, informationprotectionsensitivitylabel.DefaultGetInformationProtectionSensitivityLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.ListInformationProtectionSensitivityLabels`

```go
ctx := context.TODO()


// alternatively `client.ListInformationProtectionSensitivityLabels(ctx, informationprotectionsensitivitylabel.DefaultListInformationProtectionSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionSensitivityLabelsComplete(ctx, informationprotectionsensitivitylabel.DefaultListInformationProtectionSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.UpdateInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabel.NewMeInformationProtectionSensitivityLabelID("sensitivityLabelId")

payload := informationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateInformationProtectionSensitivityLabel(ctx, id, payload, informationprotectionsensitivitylabel.DefaultUpdateInformationProtectionSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
