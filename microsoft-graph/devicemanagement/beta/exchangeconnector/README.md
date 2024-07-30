
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeconnector` Documentation

The `exchangeconnector` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/exchangeconnector"
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


### Example Usage: `ExchangeConnectorClient.CreateExchangeConnectorSync`

```go
ctx := context.TODO()
id := exchangeconnector.NewDeviceManagementExchangeConnectorID("deviceManagementExchangeConnectorIdValue")

payload := exchangeconnector.CreateExchangeConnectorSyncRequest{
	// ...
}


read, err := client.CreateExchangeConnectorSync(ctx, id, payload)
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

read, err := client.DeleteExchangeConnector(ctx, id)
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

read, err := client.GetExchangeConnector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ExchangeConnectorClient.GetExchangeConnectorCount`

```go
ctx := context.TODO()


read, err := client.GetExchangeConnectorCount(ctx)
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


// alternatively `client.ListExchangeConnectors(ctx)` can be used to do batched pagination
items, err := client.ListExchangeConnectorsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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
