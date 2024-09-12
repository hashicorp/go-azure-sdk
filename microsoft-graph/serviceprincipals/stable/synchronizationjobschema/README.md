
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjobschema` Documentation

The `synchronizationjobschema` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjobschema"
```


### Client Initialization

```go
client := synchronizationjobschema.NewSynchronizationJobSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationJobSchemaClient.DeleteSynchronizationJobSchema`

```go
ctx := context.TODO()
id := synchronizationjobschema.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.DeleteSynchronizationJobSchema(ctx, id, synchronizationjobschema.DefaultDeleteSynchronizationJobSchemaOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobSchemaClient.GetSynchronizationJobSchema`

```go
ctx := context.TODO()
id := synchronizationjobschema.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.GetSynchronizationJobSchema(ctx, id, synchronizationjobschema.DefaultGetSynchronizationJobSchemaOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobSchemaClient.ParseSynchronizationJobSchemaExpression`

```go
ctx := context.TODO()
id := synchronizationjobschema.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjobschema.ParseSynchronizationJobSchemaExpressionRequest{
	// ...
}


read, err := client.ParseSynchronizationJobSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobSchemaClient.UpdateSynchronizationJobSchema`

```go
ctx := context.TODO()
id := synchronizationjobschema.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjobschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateSynchronizationJobSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
