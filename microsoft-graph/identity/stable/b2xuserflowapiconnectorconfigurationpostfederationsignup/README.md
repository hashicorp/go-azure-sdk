
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostfederationsignup` Documentation

The `b2xuserflowapiconnectorconfigurationpostfederationsignup` SDK allows for interaction with Microsoft Graph `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostfederationsignup"
```


### Client Initialization

```go
client := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewB2xUserFlowApiConnectorConfigurationPostFederationSignupClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificate`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

payload := b2xuserflowapiconnectorconfigurationpostfederationsignup.CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificate(ctx, id, payload, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultCreateB2xUserFlowApiConnectorConfigurationPostFederationSignupUploadClientCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

read, err := client.DeleteB2xUserFlowApiConnectorConfigurationPostFederationSignup(ctx, id, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultDeleteB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.GetB2xUserFlowApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

read, err := client.GetB2xUserFlowApiConnectorConfigurationPostFederationSignup(ctx, id, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultGetB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

read, err := client.GetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx, id, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultGetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

read, err := client.RemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx, id, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultRemoveB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

payload := b2xuserflowapiconnectorconfigurationpostfederationsignup.ReferenceUpdate{
	// ...
}


read, err := client.SetB2xUserFlowApiConnectorConfigurationPostFederationSignupRef(ctx, id, payload, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultSetB2xUserFlowApiConnectorConfigurationPostFederationSignupRefOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostFederationSignupClient.UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowId")

payload := b2xuserflowapiconnectorconfigurationpostfederationsignup.IdentityApiConnector{
	// ...
}


read, err := client.UpdateB2xUserFlowApiConnectorConfigurationPostFederationSignup(ctx, id, payload, b2xuserflowapiconnectorconfigurationpostfederationsignup.DefaultUpdateB2xUserFlowApiConnectorConfigurationPostFederationSignupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
