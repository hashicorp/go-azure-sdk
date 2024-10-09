
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervhost` Documentation

The `hypervhost` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervhost"
```


### Client Initialization

```go
client := hypervhost.NewHyperVHostClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HyperVHostClient.GetAllHostsInSite`

```go
ctx := context.TODO()
id := hypervhost.NewHyperVSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName")

// alternatively `client.GetAllHostsInSite(ctx, id, hypervhost.DefaultGetAllHostsInSiteOperationOptions())` can be used to do batched pagination
items, err := client.GetAllHostsInSiteComplete(ctx, id, hypervhost.DefaultGetAllHostsInSiteOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `HyperVHostClient.GetHost`

```go
ctx := context.TODO()
id := hypervhost.NewHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName", "hostName")

read, err := client.GetHost(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HyperVHostClient.PutHost`

```go
ctx := context.TODO()
id := hypervhost.NewHostID("12345678-1234-9876-4563-123456789012", "example-resource-group", "hyperVSiteName", "hostName")

payload := hypervhost.HyperVHost{
	// ...
}


read, err := client.PutHost(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
