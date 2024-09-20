
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicylabel` Documentation

The `informationprotectionpolicylabel` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicylabel"
```


### Client Initialization

```go
client := informationprotectionpolicylabel.NewInformationProtectionPolicyLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewUserID("userId")

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabel(ctx, id, payload, informationprotectionpolicylabel.DefaultCreateInformationProtectionPolicyLabelOperationOptions())
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userId", "informationProtectionLabelId")

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
id := informationprotectionpolicylabel.NewUserID("userId")

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsApplicationsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsApplications(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsApplicationsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsApplicationsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userId")

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsClassificationResultsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsClassificationResults(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsClassificationResultsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsClassificationResultsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userId")

payload := informationprotectionpolicylabel.EvaluateInformationProtectionPolicyLabelsRemovalsRequest{
	// ...
}


// alternatively `client.EvaluateInformationProtectionPolicyLabelsRemovals(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.EvaluateInformationProtectionPolicyLabelsRemovalsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultEvaluateInformationProtectionPolicyLabelsRemovalsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userId")

payload := informationprotectionpolicylabel.ExtractInformationProtectionPolicyLabelsLabelRequest{
	// ...
}


read, err := client.ExtractInformationProtectionPolicyLabelsLabel(ctx, id, payload, informationprotectionpolicylabel.DefaultExtractInformationProtectionPolicyLabelsLabelOperationOptions())
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userId", "informationProtectionLabelId")

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
id := informationprotectionpolicylabel.NewUserID("userId")

read, err := client.GetInformationProtectionPolicyLabelsCount(ctx, id, informationprotectionpolicylabel.DefaultGetInformationProtectionPolicyLabelsCountOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userId")

// alternatively `client.ListInformationProtectionPolicyLabels(ctx, id, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelsComplete(ctx, id, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userId", "informationProtectionLabelId")

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
