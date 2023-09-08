
## `github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipalsynchronizationjobschema` Documentation

The `serviceprincipalsynchronizationjobschema` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `v1.0`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/v1.0/serviceprincipalsynchronizationjobschema"
```


### Client Initialization

```go
client := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobSchemaClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalSynchronizationJobSchemaClient.DeleteServicePrincipalByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.DeleteServicePrincipalByIdSynchronizationJobByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobSchemaClient.GetServicePrincipalByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.GetServicePrincipalByIdSynchronizationJobByIdSchema(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobSchemaClient.ParseServicePrincipalByIdSynchronizationJobByIdSchemaExpression`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjobschema.ParseServicePrincipalByIdSynchronizationJobByIdSchemaExpressionRequest{
	// ...
}


read, err := client.ParseServicePrincipalByIdSynchronizationJobByIdSchemaExpression(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobSchemaClient.UpdateServicePrincipalByIdSynchronizationJobByIdSchema`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjobschema.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjobschema.SynchronizationSchema{
	// ...
}


read, err := client.UpdateServicePrincipalByIdSynchronizationJobByIdSchema(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
