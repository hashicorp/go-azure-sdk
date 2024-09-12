
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronization` Documentation

The `synchronization` SDK allows for interaction with the Azure Resource Manager Service `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/synchronization"
```


### Client Initialization

```go
client := synchronization.NewSynchronizationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationClient.AcquireSynchronizationAccessToken`

```go
ctx := context.TODO()
id := synchronization.NewServicePrincipalID("servicePrincipalIdValue")

payload := synchronization.AcquireSynchronizationAccessTokenRequest{
	// ...
}


read, err := client.AcquireSynchronizationAccessToken(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationClient.DeleteSynchronization`

```go
ctx := context.TODO()
id := synchronization.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.DeleteSynchronization(ctx, id, synchronization.DefaultDeleteSynchronizationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationClient.GetSynchronization`

```go
ctx := context.TODO()
id := synchronization.NewServicePrincipalID("servicePrincipalIdValue")

read, err := client.GetSynchronization(ctx, id, synchronization.DefaultGetSynchronizationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationClient.SetSynchronization`

```go
ctx := context.TODO()
id := synchronization.NewServicePrincipalID("servicePrincipalIdValue")

payload := synchronization.Synchronization{
	// ...
}


read, err := client.SetSynchronization(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
