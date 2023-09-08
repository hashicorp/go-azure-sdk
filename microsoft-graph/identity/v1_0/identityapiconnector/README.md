
## `github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityapiconnector` Documentation

The `identityapiconnector` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityapiconnector"
```


### Client Initialization

```go
client := identityapiconnector.NewIdentityApiConnectorClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityApiConnectorClient.CreateIdentityApiConnector`

```go
ctx := context.TODO()

payload := identityapiconnector.IdentityApiConnector{
	// ...
}


read, err := client.CreateIdentityApiConnector(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityApiConnectorClient.CreateIdentityApiConnectorByIdUploadClientCertificate`

```go
ctx := context.TODO()
id := identityapiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

payload := identityapiconnector.CreateIdentityApiConnectorByIdUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateIdentityApiConnectorByIdUploadClientCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityApiConnectorClient.DeleteIdentityApiConnectorById`

```go
ctx := context.TODO()
id := identityapiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

read, err := client.DeleteIdentityApiConnectorById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityApiConnectorClient.GetIdentityApiConnectorById`

```go
ctx := context.TODO()
id := identityapiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

read, err := client.GetIdentityApiConnectorById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityApiConnectorClient.GetIdentityApiConnectorCount`

```go
ctx := context.TODO()


read, err := client.GetIdentityApiConnectorCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityApiConnectorClient.ListIdentityApiConnectors`

```go
ctx := context.TODO()


// alternatively `client.ListIdentityApiConnectors(ctx)` can be used to do batched pagination
items, err := client.ListIdentityApiConnectorsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IdentityApiConnectorClient.UpdateIdentityApiConnectorById`

```go
ctx := context.TODO()
id := identityapiconnector.NewIdentityApiConnectorID("identityApiConnectorIdValue")

payload := identityapiconnector.IdentityApiConnector{
	// ...
}


read, err := client.UpdateIdentityApiConnectorById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
