
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjob` Documentation

The `synchronizationjob` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronizationjob"
```


### Client Initialization

```go
client := synchronizationjob.NewSynchronizationJobClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationJobClient.CreateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.CreateSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.DeleteSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.DeleteSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.GetSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJobCount`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetSynchronizationJobCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ListSynchronizationJobs`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

// alternatively `client.ListSynchronizationJobs(ctx, id)` can be used to do batched pagination
items, err := client.ListSynchronizationJobsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SynchronizationJobClient.PauseServicePrincipalSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.PauseServicePrincipalSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ProvisionServicePrincipalSynchronizationJobOnDemand`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ProvisionServicePrincipalSynchronizationJobOnDemandRequest{
	// ...
}


read, err := client.ProvisionServicePrincipalSynchronizationJobOnDemand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.RestartServicePrincipalSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.RestartServicePrincipalSynchronizationJobRequest{
	// ...
}


read, err := client.RestartServicePrincipalSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.StartServicePrincipalSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

read, err := client.StartServicePrincipalSynchronizationJob(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.UpdateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.UpdateSynchronizationJob(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateServicePrincipalSynchronizationJobCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalIdValue", "synchronizationJobIdValue")

payload := synchronizationjob.ValidateServicePrincipalSynchronizationJobCredentialRequest{
	// ...
}


read, err := client.ValidateServicePrincipalSynchronizationJobCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateServicePrincipalSynchronizationJobsCredential`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalIdValue")

payload := synchronizationjob.ValidateServicePrincipalSynchronizationJobsCredentialRequest{
	// ...
}


read, err := client.ValidateServicePrincipalSynchronizationJobsCredential(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
