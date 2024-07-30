
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostattributecollection` Documentation

The `b2xuserflowapiconnectorconfigurationpostattributecollection` SDK allows for interaction with the Azure Resource Manager Service `identity` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/b2xuserflowapiconnectorconfigurationpostattributecollection"
```


### Client Initialization

```go
client := b2xuserflowapiconnectorconfigurationpostattributecollection.NewB2xUserFlowApiConnectorConfigurationPostAttributeCollectionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := b2xuserflowapiconnectorconfigurationpostattributecollection.CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificateRequest{
	// ...
}


read, err := client.CreateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionUploadClientCertificate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.CreateUpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := b2xuserflowapiconnectorconfigurationpostattributecollection.ReferenceUpdate{
	// ...
}


read, err := client.CreateUpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.GetB2xUserFlowApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

read, err := client.RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient.UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollection`

```go
ctx := context.TODO()
id := b2xuserflowapiconnectorconfigurationpostattributecollection.NewIdentityB2xUserFlowID("b2xIdentityUserFlowIdValue")

payload := b2xuserflowapiconnectorconfigurationpostattributecollection.IdentityApiConnector{
	// ...
}


read, err := client.UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
