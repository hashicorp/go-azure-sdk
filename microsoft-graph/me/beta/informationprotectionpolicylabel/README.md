
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


### Example Usage: `InformationProtectionPolicyLabelClient.DeleteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelIdValue")

read, err := client.DeleteInformationProtectionPolicyLabel(ctx, id, informationprotectionpolicylabel.DefaultDeleteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ExtractInformationProtectionPolicyLabel`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.ExtractInformationProtectionPolicyLabelRequest{
	// ...
}


read, err := client.ExtractInformationProtectionPolicyLabel(ctx, payload)
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

read, err := client.GetInformationProtectionPolicyLabel(ctx, id, informationprotectionpolicylabel.DefaultGetInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.GetInformationProtectionPolicyLabelsCount`

```go
ctx := context.TODO()


read, err := client.GetInformationProtectionPolicyLabelsCount(ctx, informationprotectionpolicylabel.DefaultGetInformationProtectionPolicyLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ListInformationProtectionPolicyLabelEvaluateApplications`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateApplicationsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateApplications(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateApplicationsComplete(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ListInformationProtectionPolicyLabelEvaluateClassificationResults`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateClassificationResultsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateClassificationResults(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateClassificationResultsComplete(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ListInformationProtectionPolicyLabelEvaluateRemovals`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateRemovalsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateRemovals(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateRemovalsComplete(ctx, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ListInformationProtectionPolicyLabels`

```go
ctx := context.TODO()


// alternatively `client.ListInformationProtectionPolicyLabels(ctx, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelsComplete(ctx, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelsOperationOptions())
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
