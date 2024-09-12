
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionsensitivitylabel` Documentation

The `informationprotectionsensitivitylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionsensitivitylabel"
```


### Client Initialization

```go
client := informationprotectionsensitivitylabel.NewInformationProtectionSensitivityLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionSensitivityLabelClient.CreateInformationProtectionSensitivityLabel`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := informationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionSensitivityLabelClient.CreateInformationProtectionSensitivityLabelEvaluate`

```go
ctx := context.TODO()
id := informationprotectionsensitivitylabel.NewUserID("userIdValue")

payload := informationprotectionsensitivitylabel.CreateInformationProtectionSensitivityLabelEvaluateRequest{
	// ...
}


read, err := client.CreateInformationProtectionSensitivityLabelEvaluate(ctx, id, payload)
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
id := informationprotectionsensitivitylabel.NewUserIdInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

read, err := client.DeleteInformationProtectionSensitivityLabel(ctx, id, informationprotectionsensitivitylabel.DefaultDeleteInformationProtectionSensitivityLabelOperationOptions())
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
id := informationprotectionsensitivitylabel.NewUserIdInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

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
id := informationprotectionsensitivitylabel.NewUserID("userIdValue")

read, err := client.GetInformationProtectionSensitivityLabelsCount(ctx, id, informationprotectionsensitivitylabel.DefaultGetInformationProtectionSensitivityLabelsCountOperationOptions())
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
id := informationprotectionsensitivitylabel.NewUserID("userIdValue")

// alternatively `client.ListInformationProtectionSensitivityLabels(ctx, id, informationprotectionsensitivitylabel.DefaultListInformationProtectionSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionSensitivityLabelsComplete(ctx, id, informationprotectionsensitivitylabel.DefaultListInformationProtectionSensitivityLabelsOperationOptions())
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
id := informationprotectionsensitivitylabel.NewUserIdInformationProtectionSensitivityLabelID("userIdValue", "sensitivityLabelIdValue")

payload := informationprotectionsensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateInformationProtectionSensitivityLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
