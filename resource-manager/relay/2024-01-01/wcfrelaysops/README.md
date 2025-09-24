
## `github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelaysops` Documentation

The `wcfrelaysops` SDK allows for interaction with Azure Resource Manager `relay` (API Version `2024-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/relay/2024-01-01/wcfrelaysops"
```


### Client Initialization

```go
client := wcfrelaysops.NewWcfRelaysOpsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WcfRelaysOpsClient.WCFRelaysCreateOrUpdate`

```go
ctx := context.TODO()
id := wcfrelaysops.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName")

payload := wcfrelaysops.WcfRelay{
	// ...
}


read, err := client.WCFRelaysCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WcfRelaysOpsClient.WCFRelaysDelete`

```go
ctx := context.TODO()
id := wcfrelaysops.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName")

read, err := client.WCFRelaysDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WcfRelaysOpsClient.WCFRelaysGet`

```go
ctx := context.TODO()
id := wcfrelaysops.NewWcfRelayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName", "wcfRelayName")

read, err := client.WCFRelaysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WcfRelaysOpsClient.WCFRelaysListByNamespace`

```go
ctx := context.TODO()
id := wcfrelaysops.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceName")

// alternatively `client.WCFRelaysListByNamespace(ctx, id)` can be used to do batched pagination
items, err := client.WCFRelaysListByNamespaceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
