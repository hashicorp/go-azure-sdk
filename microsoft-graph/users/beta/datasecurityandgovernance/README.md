
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernance` Documentation

The `datasecurityandgovernance` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/datasecurityandgovernance"
```


### Client Initialization

```go
client := datasecurityandgovernance.NewDataSecurityAndGovernanceClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataSecurityAndGovernanceClient.DeleteDataSecurityAndGovernance`

```go
ctx := context.TODO()
id := datasecurityandgovernance.NewUserID("userId")

read, err := client.DeleteDataSecurityAndGovernance(ctx, id, datasecurityandgovernance.DefaultDeleteDataSecurityAndGovernanceOperationOptions())
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
id := datasecurityandgovernance.NewUserID("userId")

read, err := client.GetDataSecurityAndGovernance(ctx, id, datasecurityandgovernance.DefaultGetDataSecurityAndGovernanceOperationOptions())
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
id := datasecurityandgovernance.NewUserID("userId")

payload := datasecurityandgovernance.ProcessDataSecurityAndGovernanceContentRequest{
	// ...
}


read, err := client.ProcessDataSecurityAndGovernanceContent(ctx, id, payload, datasecurityandgovernance.DefaultProcessDataSecurityAndGovernanceContentOperationOptions())
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
id := datasecurityandgovernance.NewUserID("userId")

payload := datasecurityandgovernance.UserDataSecurityAndGovernance{
	// ...
}


read, err := client.UpdateDataSecurityAndGovernance(ctx, id, payload, datasecurityandgovernance.DefaultUpdateDataSecurityAndGovernanceOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
