
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotadeployment` Documentation

The `zebrafotadeployment` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotadeployment"
```


### Client Initialization

```go
client := zebrafotadeployment.NewZebraFotaDeploymentClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ZebraFotaDeploymentClient.CancelZebraFotaDeployment`

```go
ctx := context.TODO()
id := zebrafotadeployment.NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentId")

read, err := client.CancelZebraFotaDeployment(ctx, id, zebrafotadeployment.DefaultCancelZebraFotaDeploymentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaDeploymentClient.CreateZebraFotaDeployment`

```go
ctx := context.TODO()

payload := zebrafotadeployment.ZebraFotaDeployment{
	// ...
}


read, err := client.CreateZebraFotaDeployment(ctx, payload, zebrafotadeployment.DefaultCreateZebraFotaDeploymentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaDeploymentClient.DeleteZebraFotaDeployment`

```go
ctx := context.TODO()
id := zebrafotadeployment.NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentId")

read, err := client.DeleteZebraFotaDeployment(ctx, id, zebrafotadeployment.DefaultDeleteZebraFotaDeploymentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaDeploymentClient.GetZebraFotaDeployment`

```go
ctx := context.TODO()
id := zebrafotadeployment.NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentId")

read, err := client.GetZebraFotaDeployment(ctx, id, zebrafotadeployment.DefaultGetZebraFotaDeploymentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaDeploymentClient.GetZebraFotaDeploymentsCount`

```go
ctx := context.TODO()


read, err := client.GetZebraFotaDeploymentsCount(ctx, zebrafotadeployment.DefaultGetZebraFotaDeploymentsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaDeploymentClient.ListZebraFotaDeployments`

```go
ctx := context.TODO()


// alternatively `client.ListZebraFotaDeployments(ctx, zebrafotadeployment.DefaultListZebraFotaDeploymentsOperationOptions())` can be used to do batched pagination
items, err := client.ListZebraFotaDeploymentsComplete(ctx, zebrafotadeployment.DefaultListZebraFotaDeploymentsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ZebraFotaDeploymentClient.UpdateZebraFotaDeployment`

```go
ctx := context.TODO()
id := zebrafotadeployment.NewDeviceManagementZebraFotaDeploymentID("zebraFotaDeploymentId")

payload := zebrafotadeployment.ZebraFotaDeployment{
	// ...
}


read, err := client.UpdateZebraFotaDeployment(ctx, id, payload, zebrafotadeployment.DefaultUpdateZebraFotaDeploymentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
