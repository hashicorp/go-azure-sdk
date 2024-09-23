
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjobschema` Documentation

The `synchronizationjobschema` SDK allows for interaction with Microsoft Graph `applications` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/stable/synchronizationjobschema"
```


### Client Initialization

```go
client := synchronizationjobschema.NewSynchronizationJobSchemaClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationJobSchemaClient.DeleteSynchronizationJobSchema`

```go
ctx := context.TODO()
id := synchronizationjobschema.NewApplicationIdSynchronizationJobID("applicationId", "synchronizationJobId")

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
id := synchronizationjobschema.NewApplicationIdSynchronizationJobID("applicationId", "synchronizationJobId")

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
id := synchronizationjobschema.NewApplicationIdSynchronizationJobID("applicationId", "synchronizationJobId")

payload := synchronizationjobschema.ParseSynchronizationJobSchemaExpressionRequest{
	// ...
}


read, err := client.ParseSynchronizationJobSchemaExpression(ctx, id, payload, synchronizationjobschema.DefaultParseSynchronizationJobSchemaExpressionOperationOptions())
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
id := synchronizationjobschema.NewApplicationIdSynchronizationJobID("applicationId", "synchronizationJobId")

payload := synchronizationjobschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateSynchronizationJobSchema(ctx, id, payload, synchronizationjobschema.DefaultUpdateSynchronizationJobSchemaOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
