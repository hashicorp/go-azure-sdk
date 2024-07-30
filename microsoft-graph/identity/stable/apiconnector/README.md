
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/apiconnector` Documentation

The `apiconnector` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/apiconnector"
```


### Client Initialization

```go
client := apiconnector.NewApiConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApiConnectorClient.CreateApiConnector`

```go
ctx := context.TODO()

payload := apiconnector.IdentityApiConnector{
	// ...
}


read, err := client.CreateApiConnector(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiConnectorClient.CreateApiConnectorUploadClientCertificate`

```go
ctx := context.TODO()
id := apiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

payload := apiconnector.CreateApiConnectorUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateApiConnectorUploadClientCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiConnectorClient.DeleteApiConnector`

```go
ctx := context.TODO()
id := apiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

read, err := client.DeleteApiConnector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiConnectorClient.GetApiConnector`

```go
ctx := context.TODO()
id := apiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

read, err := client.GetApiConnector(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiConnectorClient.GetApiConnectorCount`

```go
ctx := context.TODO()


read, err := client.GetApiConnectorCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApiConnectorClient.ListApiConnectors`

```go
ctx := context.TODO()


// alternatively `client.ListApiConnectors(ctx)` can be used to do batched pagination
items, err := client.ListApiConnectorsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApiConnectorClient.UpdateApiConnector`

```go
ctx := context.TODO()
id := apiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

payload := apiconnector.IdentityApiConnector{
	// ...
}


read, err := client.UpdateApiConnector(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
