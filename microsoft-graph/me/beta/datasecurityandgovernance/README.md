
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/datasecurityandgovernance` Documentation

The `datasecurityandgovernance` SDK allows for interaction with Microsoft Graph `me` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/beta/datasecurityandgovernance"
```


### Client Initialization

```go
client := datasecurityandgovernance.NewDataSecurityAndGovernanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataSecurityAndGovernanceClient.DeleteDataSecurityAndGovernance`

```go
ctx := context.TODO()


read, err := client.DeleteDataSecurityAndGovernance(ctx, datasecurityandgovernance.DefaultDeleteDataSecurityAndGovernanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceClient.GetDataSecurityAndGovernance`

```go
ctx := context.TODO()


read, err := client.GetDataSecurityAndGovernance(ctx, datasecurityandgovernance.DefaultGetDataSecurityAndGovernanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceClient.ProcessDataSecurityAndGovernanceContent`

```go
ctx := context.TODO()

payload := datasecurityandgovernance.ProcessDataSecurityAndGovernanceContentRequest{
	// ...
}


read, err := client.ProcessDataSecurityAndGovernanceContent(ctx, payload, datasecurityandgovernance.DefaultProcessDataSecurityAndGovernanceContentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DataSecurityAndGovernanceClient.UpdateDataSecurityAndGovernance`

```go
ctx := context.TODO()

payload := datasecurityandgovernance.UserDataSecurityAndGovernance{
	// ...
}


read, err := client.UpdateDataSecurityAndGovernance(ctx, payload, datasecurityandgovernance.DefaultUpdateDataSecurityAndGovernanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
