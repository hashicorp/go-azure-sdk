
## `github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/beta/serviceprincipalsynchronizationjob` Documentation

The `serviceprincipalsynchronizationjob` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/serviceprincipals/beta/serviceprincipalsynchronizationjob"
```


### Client Initialization

```go
client := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.CreateServicePrincipalByIdSynchronizationJob`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipalsynchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.CreateServicePrincipalByIdSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.DeleteServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.DeleteServicePrincipalByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.GetServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.GetServicePrincipalByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.GetServicePrincipalByIdSynchronizationJobCount`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetServicePrincipalByIdSynchronizationJobCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.ListServicePrincipalByIdSynchronizationJobs`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

// alternatively `client.ListServicePrincipalByIdSynchronizationJobs(ctx, id)` can be used to do batched pagination
items, err := client.ListServicePrincipalByIdSynchronizationJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.PauseServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.PauseServicePrincipalByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.ProvisionServicePrincipalByIdSynchronizationJobByIdOnDemand`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjob.ProvisionServicePrincipalByIdSynchronizationJobByIdOnDemandRequest{
	// ...
}


read, err := client.ProvisionServicePrincipalByIdSynchronizationJobByIdOnDemand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.RestartServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjob.RestartServicePrincipalByIdSynchronizationJobByIdRequest{
	// ...
}


read, err := client.RestartServicePrincipalByIdSynchronizationJobById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.StartServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.StartServicePrincipalByIdSynchronizationJobById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.UpdateServicePrincipalByIdSynchronizationJobById`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.UpdateServicePrincipalByIdSynchronizationJobById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.ValidateServicePrincipalByIdSynchronizationJobByIdCredential`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := serviceprincipalsynchronizationjob.ValidateServicePrincipalByIdSynchronizationJobByIdCredentialRequest{
	// ...
}


read, err := client.ValidateServicePrincipalByIdSynchronizationJobByIdCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalSynchronizationJobClient.ValidateServicePrincipalByIdSynchronizationJobsCredential`

```go
ctx := context.TODO()
id := serviceprincipalsynchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

payload := serviceprincipalsynchronizationjob.ValidateServicePrincipalByIdSynchronizationJobsCredentialRequest{
	// ...
}


read, err := client.ValidateServicePrincipalByIdSynchronizationJobsCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
