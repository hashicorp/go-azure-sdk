
## `github.com/hashicorp/go-azure-sdk/resource-manager/applications/v1.0/applicationsynchronizationjobschema` Documentation

The `applicationsynchronizationjobschema` SDK allows for interaction with the Azure Resource Manager Service `applications` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/applications/v1.0/applicationsynchronizationjobschema"
```


### Client Initialization

```go
client := applicationsynchronizationjobschema.NewApplicationSynchronizationJobSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationSynchronizationJobSchemaClient.DeleteApplicationByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationjobschema.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.DeleteApplicationByIdSynchronizationJobByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobSchemaClient.GetApplicationByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationjobschema.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

read, err := client.GetApplicationByIdSynchronizationJobByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobSchemaClient.ParseApplicationByIdSynchronizationJobByIdSchemaExpression`

```go
ctx := context.TODO()
id := applicationsynchronizationjobschema.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjobschema.ParseApplicationByIdSynchronizationJobByIdSchemaExpressionRequest{
	// ...
}


read, err := client.ParseApplicationByIdSynchronizationJobByIdSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationSynchronizationJobSchemaClient.UpdateApplicationByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := applicationsynchronizationjobschema.NewApplicationSynchronizationJobID("applicationIdValue", "synchronizationJobIdValue")

payload := applicationsynchronizationjobschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateApplicationByIdSynchronizationJobByIdSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
