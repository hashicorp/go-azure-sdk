
## `github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowapiconnectorconfigurationpostfederationsignup` Documentation

The `identityb2xuserflowapiconnectorconfigurationpostfederationsignup` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowapiconnectorconfigurationpostfederationsignup"
```


### Client Initialization

```go
client := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupUploadClientCertificate`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupUploadClientCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.ReferenceUpdate{
	// ...
}


read, err := client.CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.DeleteIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.DeleteIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.RemoveIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.RemoveIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient.UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostfederationsignup.IdentityApiConnector{
	// ...
}


read, err := client.UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
