
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/hosts` Documentation

The `hosts` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2024-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2024-09-01/hosts"
```


### Client Initialization

```go
client := hosts.NewHostsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HostsClient.Get`

```go
ctx := context.TODO()
id := hosts.NewHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "clusterName", "hostId")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostsClient.List`

```go
ctx := context.TODO()
id := hosts.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudName", "clusterName")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
