
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotaconnector` Documentation

The `zebrafotaconnector` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/zebrafotaconnector"
```


### Client Initialization

```go
client := zebrafotaconnector.NewZebraFotaConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ZebraFotaConnectorClient.CreateZebraFotaConnectorApproveFotaApp`

```go
ctx := context.TODO()


read, err := client.CreateZebraFotaConnectorApproveFotaApp(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.CreateZebraFotaConnectorConnect`

```go
ctx := context.TODO()


read, err := client.CreateZebraFotaConnectorConnect(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.CreateZebraFotaConnectorDisconnect`

```go
ctx := context.TODO()


read, err := client.CreateZebraFotaConnectorDisconnect(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.CreateZebraFotaConnectorHasActiveDeployment`

```go
ctx := context.TODO()


read, err := client.CreateZebraFotaConnectorHasActiveDeployment(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.DeleteZebraFotaConnector`

```go
ctx := context.TODO()


read, err := client.DeleteZebraFotaConnector(ctx, zebrafotaconnector.DefaultDeleteZebraFotaConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.GetZebraFotaConnector`

```go
ctx := context.TODO()


read, err := client.GetZebraFotaConnector(ctx, zebrafotaconnector.DefaultGetZebraFotaConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ZebraFotaConnectorClient.UpdateZebraFotaConnector`

```go
ctx := context.TODO()

payload := zebrafotaconnector.ZebraFotaConnector{
	// ...
}


read, err := client.UpdateZebraFotaConnector(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
