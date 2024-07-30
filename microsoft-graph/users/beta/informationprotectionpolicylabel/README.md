
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


### Example Usage: `InformationProtectionPolicyLabelClient.CreateInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateApplication(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelEvaluateRemoval(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

payload := informationprotectionpolicylabel.CreateInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateInformationProtectionPolicyLabelExtractLabel(ctx, id, payload)
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
id := informationprotectionpolicylabel.NewUserIdInformationProtectionPolicyLabelID("userIdValue", "informationProtectionLabelIdValue")

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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

read, err := client.GetInformationProtectionPolicyLabelCount(ctx, id)
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
id := informationprotectionpolicylabel.NewUserID("userIdValue")

// alternatively `client.ListInformationProtectionPolicyLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListInformationProtectionPolicyLabelsComplete(ctx, id)
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
