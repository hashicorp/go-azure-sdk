
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicylabel` Documentation

The `informationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/informationprotectionpolicylabel"
```


### Client Initialization

```go
client := informationprotectionpolicylabel.NewInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabel(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.ExtractInformationProtectionPolicyLabelRequest{
	// ...
}


read, err := client.ExtractInformationProtectionPolicyLabel(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

read, err := client.GetInformationProtectionPolicyLabelsCount(ctx, id, informationprotectionpolicylabel.DefaultGetInformationProtectionPolicyLabelsCountOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateApplicationsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateApplications(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateApplicationsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateApplicationsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateClassificationResultsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateClassificationResults(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateClassificationResultsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateClassificationResultsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.ListInformationProtectionPolicyLabelEvaluateRemovalsRequest{
	// ...
}


// alternatively `client.ListInformationProtectionPolicyLabelEvaluateRemovals(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelEvaluateRemovalsComplete(ctx, id, payload, informationprotectionpolicylabel.DefaultListInformationProtectionPolicyLabelEvaluateRemovalsOperationOptions())
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

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
