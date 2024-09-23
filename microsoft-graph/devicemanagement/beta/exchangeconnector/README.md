
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeconnector` Documentation

The `exchangeconnector` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeconnector"
```


### Client Initialization

```go
client := exchangeconnector.NewExchangeConnectorClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExchangeConnectorClient.CreateExchangeConnector`

```go
ctx := context.TODO()

payload := exchangeconnector.DeviceManagementExchangeConnector{
	// ...
}


read, err := client.CreateExchangeConnector(ctx, payload, exchangeconnector.DefaultCreateExchangeConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.DeleteExchangeConnector`

```go
ctx := context.TODO()
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorId")

read, err := client.DeleteExchangeConnector(ctx, id, exchangeconnector.DefaultDeleteExchangeConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.GetExchangeConnector`

```go
ctx := context.TODO()
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorId")

read, err := client.GetExchangeConnector(ctx, id, exchangeconnector.DefaultGetExchangeConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.GetExchangeConnectorsCount`

```go
ctx := context.TODO()


read, err := client.GetExchangeConnectorsCount(ctx, exchangeconnector.DefaultGetExchangeConnectorsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.ListExchangeConnectors`

```go
ctx := context.TODO()


// alternatively `client.ListExchangeConnectors(ctx, exchangeconnector.DefaultListExchangeConnectorsOperationOptions())` can be used to do batched pagination
items, err := client.ListExchangeConnectorsComplete(ctx, exchangeconnector.DefaultListExchangeConnectorsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ExchangeConnectorClient.SyncExchangeConnector`

```go
ctx := context.TODO()
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorId")

payload := exchangeconnector.SyncExchangeConnectorRequest{
	// ...
}


read, err := client.SyncExchangeConnector(ctx, id, payload, exchangeconnector.DefaultSyncExchangeConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.UpdateExchangeConnector`

```go
ctx := context.TODO()
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorId")

payload := exchangeconnector.DeviceManagementExchangeConnector{
	// ...
}


read, err := client.UpdateExchangeConnector(ctx, id, payload, exchangeconnector.DefaultUpdateExchangeConnectorOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
