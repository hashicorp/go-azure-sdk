
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/exchangeconnector` Documentation

The `exchangeconnector` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/stable/exchangeconnector"
```


### Client Initialization

```go
client := exchangeconnector.NewExchangeConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ExchangeConnectorClient.CreateExchangeConnector`

```go
ctx := context.TODO()

payload := exchangeconnector.DeviceManagementExchangeConnector{
	// ...
}


read, err := client.CreateExchangeConnector(ctx, payload)
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
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

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
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

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
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

payload := exchangeconnector.SyncExchangeConnectorRequest{
	// ...
}


read, err := client.SyncExchangeConnector(ctx, id, payload)
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
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

payload := exchangeconnector.DeviceManagementExchangeConnector{
	// ...
}


read, err := client.UpdateExchangeConnector(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
