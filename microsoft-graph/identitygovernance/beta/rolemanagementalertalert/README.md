
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/rolemanagementalertalert` Documentation

The `rolemanagementalertalert` SDK allows for interaction with the Azure Resource Manager Service `identitygovernance` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/rolemanagementalertalert"
```


### Client Initialization

```go
client := rolemanagementalertalert.NewRoleManagementAlertAlertClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RoleManagementAlertAlertClient.CreateRoleManagementAlert`

```go
ctx := context.TODO()

payload := rolemanagementalertalert.UnifiedRoleManagementAlert{
	// ...
}


read, err := client.CreateRoleManagementAlert(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementAlertAlertClient.CreateRoleManagementAlertRefresh`

```go
ctx := context.TODO()
id := rolemanagementalertalert.NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue")

read, err := client.CreateRoleManagementAlertRefresh(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementAlertAlertClient.DeleteRoleManagementAlert`

```go
ctx := context.TODO()
id := rolemanagementalertalert.NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue")

read, err := client.DeleteRoleManagementAlert(ctx, id, rolemanagementalertalert.DefaultDeleteRoleManagementAlertOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementAlertAlertClient.GetRoleManagementAlert`

```go
ctx := context.TODO()
id := rolemanagementalertalert.NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue")

read, err := client.GetRoleManagementAlert(ctx, id, rolemanagementalertalert.DefaultGetRoleManagementAlertOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementAlertAlertClient.GetRoleManagementAlertCount`

```go
ctx := context.TODO()


read, err := client.GetRoleManagementAlertCount(ctx, rolemanagementalertalert.DefaultGetRoleManagementAlertCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RoleManagementAlertAlertClient.ListRoleManagementAlerts`

```go
ctx := context.TODO()


// alternatively `client.ListRoleManagementAlerts(ctx, rolemanagementalertalert.DefaultListRoleManagementAlertsOperationOptions())` can be used to do batched pagination
items, err := client.ListRoleManagementAlertsComplete(ctx, rolemanagementalertalert.DefaultListRoleManagementAlertsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `RoleManagementAlertAlertClient.UpdateRoleManagementAlert`

```go
ctx := context.TODO()
id := rolemanagementalertalert.NewIdentityGovernanceRoleManagementAlertAlertID("unifiedRoleManagementAlertIdValue")

payload := rolemanagementalertalert.UnifiedRoleManagementAlert{
	// ...
}


read, err := client.UpdateRoleManagementAlert(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
