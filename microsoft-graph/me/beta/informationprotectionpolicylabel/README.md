
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionpolicylabel` Documentation

The `informationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionpolicylabel"
```


### Client Initialization

```go
client := informationprotectionpolicylabel.NewInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabel`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateApplication(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabelEvaluateClassificationResult`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabelEvaluateRemoval`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateRemoval(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabelExtractLabel`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelExtractLabel(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.DeleteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

read, err := client.DeleteInformationProtectionPolicyLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.GetInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

read, err := client.GetInformationProtectionPolicyLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.GetInformationProtectionPolicyLabelCount`

```go
ctx := context.TODO()


read, err := client.GetInformationProtectionPolicyLabelCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ListInformationProtectionPolicyLabels`

```go
ctx := context.TODO()


// alternatively `client.ListInformationProtectionPolicyLabels(ctx)` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.UpdateInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
