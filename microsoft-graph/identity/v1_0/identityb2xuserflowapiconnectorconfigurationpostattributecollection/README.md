
## `github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowapiconnectorconfigurationpostattributecollection` Documentation

The `identityb2xuserflowapiconnectorconfigurationpostattributecollection` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/identity/v1.0/identityb2xuserflowapiconnectorconfigurationpostattributecollection"
```


### Client Initialization

```go
client := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostattributecollection.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostattributecollection.ReferenceUpdate{
	// ...
}


read, err := client.CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.DeleteIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.DeleteIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.RemoveIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.RemoveIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollectionRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IdentityB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := identityb2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := identityb2xuserflowapiconnectorconfigurationpostattributecollection.IdentityApiConnector{
	// ...
}


read, err := client.UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostAttributeCollection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
