
## `github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionpolicylabel` Documentation

The `meinformationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/me/beta/meinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := meinformationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.CreateMeInformationProtectionPolicyLabel`

```go
ctx := context.TODO()

payload := meinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateMeInformationProtectionPolicyLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.CreateMeInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()

payload := meinformationprotectionpolicylabel.CreateMeInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionPolicyLabelEvaluateApplication(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.CreateMeInformationProtectionPolicyLabelEvaluateClassificationResult`

```go
ctx := context.TODO()

payload := meinformationprotectionpolicylabel.CreateMeInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.CreateMeInformationProtectionPolicyLabelEvaluateRemoval`

```go
ctx := context.TODO()

payload := meinformationprotectionpolicylabel.CreateMeInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionPolicyLabelEvaluateRemoval(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.CreateMeInformationProtectionPolicyLabelExtractLabel`

```go
ctx := context.TODO()

payload := meinformationprotectionpolicylabel.CreateMeInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateMeInformationProtectionPolicyLabelExtractLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.DeleteMeInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

read, err := client.DeleteMeInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.GetMeInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

read, err := client.GetMeInformationProtectionPolicyLabelById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.GetMeInformationProtectionPolicyLabelCount`

```go
ctx := context.TODO()


read, err := client.GetMeInformationProtectionPolicyLabelCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.ListMeInformationProtectionPolicyLabels`

```go
ctx := context.TODO()


// alternatively `client.ListMeInformationProtectionPolicyLabels(ctx)` can be used to do batched pagination
items, err := client.ListMeInformationProtectionPolicyLabelsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeInformationProtectionPolicyLabelClient.UpdateMeInformationProtectionPolicyLabelById`

```go
ctx := context.TODO()
id := meinformationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

payload := meinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateMeInformationProtectionPolicyLabelById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
