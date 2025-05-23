
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-01-01/containerappsrevisionreplicas` Documentation

The `containerappsrevisionreplicas` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-01-01/containerappsrevisionreplicas"
```


### Client Initialization

```go
client := containerappsrevisionreplicas.NewContainerAppsRevisionReplicasClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ContainerAppsRevisionReplicasClient.GetReplica`

```go
ctx := context.TODO()
id := containerappsrevisionreplicas.NewReplicaID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "revisionName", "replicaName")

read, err := client.GetReplica(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ContainerAppsRevisionReplicasClient.ListReplicas`

```go
ctx := context.TODO()
id := containerappsrevisionreplicas.NewRevisionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "containerAppName", "revisionName")

read, err := client.ListReplicas(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
