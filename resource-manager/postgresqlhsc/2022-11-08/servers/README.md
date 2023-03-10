
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/servers` Documentation

The `servers` SDK allows for interaction with the Azure Resource Manager Service `postgresqlhsc` (API Version `2022-11-08`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2022-11-08/servers"
```


### Client Initialization

```go
client := servers.NewServersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServersClient.Get`

```go
ctx := context.TODO()
id := servers.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value", "serverValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServersClient.ListByCluster`

```go
ctx := context.TODO()
id := servers.NewServerGroupsv2ID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverGroupsv2Value")

read, err := client.ListByCluster(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
