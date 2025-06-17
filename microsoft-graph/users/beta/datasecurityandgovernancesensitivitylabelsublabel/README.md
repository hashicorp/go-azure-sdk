
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernancesensitivitylabelsublabel` Documentation

The `datasecurityandgovernancesensitivitylabelsublabel` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernancesensitivitylabelsublabel"
```


### Client Initialization

```go
client := datasecurityandgovernancesensitivitylabelsublabel.NewDataSecurityAndGovernanceSensitivityLabelSublabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.CreateDataSecurityAndGovernanceSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

payload := datasecurityandgovernancesensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.CreateDataSecurityAndGovernanceSensitivityLabelSublabel(ctx, id, payload, datasecurityandgovernancesensitivitylabelsublabel.DefaultCreateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.CreateDataSecurityAndGovernanceSensitivityLabelSublabelComputeRightsAndInheritance`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

payload := datasecurityandgovernancesensitivitylabelsublabel.CreateDataSecurityAndGovernanceSensitivityLabelSublabelComputeRightsAndInheritanceRequest{
	// ...
}


read, err := client.CreateDataSecurityAndGovernanceSensitivityLabelSublabelComputeRightsAndInheritance(ctx, id, payload, datasecurityandgovernancesensitivitylabelsublabel.DefaultCreateDataSecurityAndGovernanceSensitivityLabelSublabelComputeRightsAndInheritanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.DeleteDataSecurityAndGovernanceSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1")

read, err := client.DeleteDataSecurityAndGovernanceSensitivityLabelSublabel(ctx, id, datasecurityandgovernancesensitivitylabelsublabel.DefaultDeleteDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.EvaluateDataSecurityAndGovernanceSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

payload := datasecurityandgovernancesensitivitylabelsublabel.EvaluateDataSecurityAndGovernanceSensitivityLabelSublabelsRequest{
	// ...
}


read, err := client.EvaluateDataSecurityAndGovernanceSensitivityLabelSublabels(ctx, id, payload, datasecurityandgovernancesensitivitylabelsublabel.DefaultEvaluateDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.GetDataSecurityAndGovernanceSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1")

read, err := client.GetDataSecurityAndGovernanceSensitivityLabelSublabel(ctx, id, datasecurityandgovernancesensitivitylabelsublabel.DefaultGetDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.GetDataSecurityAndGovernanceSensitivityLabelSublabelsCount`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

read, err := client.GetDataSecurityAndGovernanceSensitivityLabelSublabelsCount(ctx, id, datasecurityandgovernancesensitivitylabelsublabel.DefaultGetDataSecurityAndGovernanceSensitivityLabelSublabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.ListDataSecurityAndGovernanceSensitivityLabelSublabels`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

// alternatively `client.ListDataSecurityAndGovernanceSensitivityLabelSublabels(ctx, id, datasecurityandgovernancesensitivitylabelsublabel.DefaultListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListDataSecurityAndGovernanceSensitivityLabelSublabelsComplete(ctx, id, datasecurityandgovernancesensitivitylabelsublabel.DefaultListDataSecurityAndGovernanceSensitivityLabelSublabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelSublabelClient.UpdateDataSecurityAndGovernanceSensitivityLabelSublabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabelsublabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelIdSublabelID("userId", "sensitivityLabelId", "sensitivityLabelId1")

payload := datasecurityandgovernancesensitivitylabelsublabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateDataSecurityAndGovernanceSensitivityLabelSublabel(ctx, id, payload, datasecurityandgovernancesensitivitylabelsublabel.DefaultUpdateDataSecurityAndGovernanceSensitivityLabelSublabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
