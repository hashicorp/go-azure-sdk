
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernancesensitivitylabel` Documentation

The `datasecurityandgovernancesensitivitylabel` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernancesensitivitylabel"
```


### Client Initialization

```go
client := datasecurityandgovernancesensitivitylabel.NewDataSecurityAndGovernanceSensitivityLabelClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.CreateDataSecurityAndGovernanceSensitivityLabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserID("userId")

payload := datasecurityandgovernancesensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.CreateDataSecurityAndGovernanceSensitivityLabel(ctx, id, payload, datasecurityandgovernancesensitivitylabel.DefaultCreateDataSecurityAndGovernanceSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritance`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserID("userId")

payload := datasecurityandgovernancesensitivitylabel.CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceRequest{
	// ...
}


read, err := client.CreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritance(ctx, id, payload, datasecurityandgovernancesensitivitylabel.DefaultCreateDataSecurityAndGovernanceSensitivityLabelComputeRightsAndInheritanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.DeleteDataSecurityAndGovernanceSensitivityLabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

read, err := client.DeleteDataSecurityAndGovernanceSensitivityLabel(ctx, id, datasecurityandgovernancesensitivitylabel.DefaultDeleteDataSecurityAndGovernanceSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.EvaluateDataSecurityAndGovernanceSensitivityLabels`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserID("userId")

payload := datasecurityandgovernancesensitivitylabel.EvaluateDataSecurityAndGovernanceSensitivityLabelsRequest{
	// ...
}


read, err := client.EvaluateDataSecurityAndGovernanceSensitivityLabels(ctx, id, payload, datasecurityandgovernancesensitivitylabel.DefaultEvaluateDataSecurityAndGovernanceSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.GetDataSecurityAndGovernanceSensitivityLabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

read, err := client.GetDataSecurityAndGovernanceSensitivityLabel(ctx, id, datasecurityandgovernancesensitivitylabel.DefaultGetDataSecurityAndGovernanceSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.GetDataSecurityAndGovernanceSensitivityLabelsCount`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserID("userId")

read, err := client.GetDataSecurityAndGovernanceSensitivityLabelsCount(ctx, id, datasecurityandgovernancesensitivitylabel.DefaultGetDataSecurityAndGovernanceSensitivityLabelsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.ListDataSecurityAndGovernanceSensitivityLabels`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserID("userId")

// alternatively `client.ListDataSecurityAndGovernanceSensitivityLabels(ctx, id, datasecurityandgovernancesensitivitylabel.DefaultListDataSecurityAndGovernanceSensitivityLabelsOperationOptions())` can be used to do batched pagination
items, err := client.ListDataSecurityAndGovernanceSensitivityLabelsComplete(ctx, id, datasecurityandgovernancesensitivitylabel.DefaultListDataSecurityAndGovernanceSensitivityLabelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DataSecurityAndGovernanceSensitivityLabelClient.UpdateDataSecurityAndGovernanceSensitivityLabel`

```go
ctx := context.TODO()
id := datasecurityandgovernancesensitivitylabel.NewUserIdDataSecurityAndGovernanceSensitivityLabelID("userId", "sensitivityLabelId")

payload := datasecurityandgovernancesensitivitylabel.SensitivityLabel{
	// ...
}


read, err := client.UpdateDataSecurityAndGovernanceSensitivityLabel(ctx, id, payload, datasecurityandgovernancesensitivitylabel.DefaultUpdateDataSecurityAndGovernanceSensitivityLabelOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
