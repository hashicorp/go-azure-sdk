
## `github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/v2workspaceconnectionresource` Documentation

The `v2workspaceconnectionresource` SDK allows for interaction with Azure Resource Manager `machinelearningservices` (API Version `2024-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/v2workspaceconnectionresource"
```


### Client Initialization

```go
client := v2workspaceconnectionresource.NewV2WorkspaceConnectionResourceClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionCreateOrUpdateDeployment`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "deploymentName")

payload := v2workspaceconnectionresource.EndpointDeploymentResourcePropertiesBasicResource{
	// ...
}


if err := client.ConnectionCreateOrUpdateDeploymentThenPoll(ctx, id, payload, v2workspaceconnectionresource.DefaultConnectionCreateOrUpdateDeploymentOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionDeleteDeployment`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "deploymentName")

if err := client.ConnectionDeleteDeploymentThenPoll(ctx, id, v2workspaceconnectionresource.DefaultConnectionDeleteDeploymentOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionGetAllModels`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.ConnectionGetAllModels(ctx, id)` can be used to do batched pagination
items, err := client.ConnectionGetAllModelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionGetDeployment`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionDeploymentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "deploymentName")

read, err := client.ConnectionGetDeployment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionGetModels`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

// alternatively `client.ConnectionGetModels(ctx, id, v2workspaceconnectionresource.DefaultConnectionGetModelsOperationOptions())` can be used to do batched pagination
items, err := client.ConnectionGetModelsComplete(ctx, id, v2workspaceconnectionresource.DefaultConnectionGetModelsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionListDeployments`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

// alternatively `client.ConnectionListDeployments(ctx, id, v2workspaceconnectionresource.DefaultConnectionListDeploymentsOperationOptions())` can be used to do batched pagination
items, err := client.ConnectionListDeploymentsComplete(ctx, id, v2workspaceconnectionresource.DefaultConnectionListDeploymentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistCreate`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")

payload := v2workspaceconnectionresource.RaiBlocklistPropertiesBasicResource{
	// ...
}


if err := client.ConnectionRaiBlocklistCreateThenPoll(ctx, id, payload, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistDelete`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")

if err := client.ConnectionRaiBlocklistDeleteThenPoll(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistGet`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")

read, err := client.ConnectionRaiBlocklistGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemAddBulk`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")
var payload []RaiBlocklistItemBulkRequest

if err := client.ConnectionRaiBlocklistItemAddBulkThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemCreate`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName", "raiBlocklistItemName")

payload := v2workspaceconnectionresource.RaiBlocklistItemPropertiesBasicResource{
	// ...
}


if err := client.ConnectionRaiBlocklistItemCreateThenPoll(ctx, id, payload, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistItemCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemDelete`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName", "raiBlocklistItemName")

if err := client.ConnectionRaiBlocklistItemDeleteThenPoll(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistItemDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemDeleteBulk`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")
var payload interface{}

if err := client.ConnectionRaiBlocklistItemDeleteBulkThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemGet`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistItemID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName", "raiBlocklistItemName")

read, err := client.ConnectionRaiBlocklistItemGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistItemsList`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewRaiBlocklistID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiBlocklistName")

// alternatively `client.ConnectionRaiBlocklistItemsList(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistItemsListOperationOptions())` can be used to do batched pagination
items, err := client.ConnectionRaiBlocklistItemsListComplete(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistItemsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiBlocklistsList`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

// alternatively `client.ConnectionRaiBlocklistsList(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistsListOperationOptions())` can be used to do batched pagination
items, err := client.ConnectionRaiBlocklistsListComplete(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiBlocklistsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiPoliciesList`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

// alternatively `client.ConnectionRaiPoliciesList(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiPoliciesListOperationOptions())` can be used to do batched pagination
items, err := client.ConnectionRaiPoliciesListComplete(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiPoliciesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiPolicyCreate`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiPolicyName")

payload := v2workspaceconnectionresource.RaiPolicyPropertiesBasicResource{
	// ...
}


if err := client.ConnectionRaiPolicyCreateThenPoll(ctx, id, payload, v2workspaceconnectionresource.DefaultConnectionRaiPolicyCreateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiPolicyDelete`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiPolicyName")

if err := client.ConnectionRaiPolicyDeleteThenPoll(ctx, id, v2workspaceconnectionresource.DefaultConnectionRaiPolicyDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.ConnectionRaiPolicyGet`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionRaiPolicyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName", "raiPolicyName")

read, err := client.ConnectionRaiPolicyGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsCreate`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

payload := v2workspaceconnectionresource.WorkspaceConnectionPropertiesV2BasicResource{
	// ...
}


read, err := client.WorkspaceConnectionsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsDelete`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

read, err := client.WorkspaceConnectionsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsGet`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

read, err := client.WorkspaceConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsList`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName")

// alternatively `client.WorkspaceConnectionsList(ctx, id, v2workspaceconnectionresource.DefaultWorkspaceConnectionsListOperationOptions())` can be used to do batched pagination
items, err := client.WorkspaceConnectionsListComplete(ctx, id, v2workspaceconnectionresource.DefaultWorkspaceConnectionsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsListSecrets`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

read, err := client.WorkspaceConnectionsListSecrets(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsTestConnection`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

payload := v2workspaceconnectionresource.WorkspaceConnectionPropertiesV2BasicResource{
	// ...
}


if err := client.WorkspaceConnectionsTestConnectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `V2WorkspaceConnectionResourceClient.WorkspaceConnectionsUpdate`

```go
ctx := context.TODO()
id := v2workspaceconnectionresource.NewConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceName", "connectionName")

payload := v2workspaceconnectionresource.WorkspaceConnectionUpdateParameter{
	// ...
}


read, err := client.WorkspaceConnectionsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
