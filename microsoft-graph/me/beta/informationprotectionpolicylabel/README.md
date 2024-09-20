
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionpolicylabel` Documentation

The `informationprotectionpolicylabel` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/informationprotectionpolicylabel"
```


### Client Initialization

```go
client := informationprotectionpolicylabel.NewInformationProtectionPolicyLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabel`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabel(ctx, payload, informationprotectionpolicylabel.DefaultCreateInformationProtectionPolicyLabelOperationOptions())
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
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelId")

read, err := client.DeleteInformationProtectionPolicyLabel(ctx, id, informationprotectionpolicylabel.DefaultDeleteInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.EvaluateInformationProtectionPolicyLabelsApplications`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsApplicationsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsApplications(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsApplicationsComplete(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.EvaluateInformationProtectionPolicyLabelsClassificationResults`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsClassificationResultsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsClassificationResults(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsClassificationResultsComplete(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsClassificationResultsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.EvaluateInformationProtectionPolicyLabelsRemovals`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsRemovalsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsRemovals(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsRemovalsComplete(ctx, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsRemovalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `InformationProtectionPolicyLabelClient.ExtractInformationProtectionPolicyLabelsLabel`

```go
ctx := context.TODO()

payload := informationprotectionpolicylabel.ExtractInformationProtectionPolicyLabelsLabelRequest{
	// ...
}


read, err := client.ExtractInformationProtectionPolicyLabelsLabel(ctx, payload, informationprotectionpolicylabel.DefaultExtractInformationProtectionPolicyLabelsLabelOperationOptions())
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
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelId")

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
id := informationprotectionpolicylabel.NewMeInformationProtectionPolicyLabelID("informationProtectionLabelId")

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateInformationProtectionPolicyLabel(ctx, id, payload, informationprotectionpolicylabel.DefaultUpdateInformationProtectionPolicyLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
