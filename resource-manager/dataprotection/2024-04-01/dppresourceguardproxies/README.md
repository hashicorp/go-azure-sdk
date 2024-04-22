
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/dppresourceguardproxies` Documentation

The `dppresourceguardproxies` SDK allows for interaction with the Azure Resource Manager Service `dataprotection` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2024-04-01/dppresourceguardproxies"
```


### Client Initialization

```go
client := dppresourceguardproxies.NewDppResourceGuardProxiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DppResourceGuardProxiesClient.DppResourceGuardProxyCreateOrUpdate`

```go
ctx := context.TODO()
id := dppresourceguardproxies.NewBackupResourceGuardProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "backupResourceGuardProxyValue")

payload := dppresourceguardproxies.ResourceGuardProxyBaseResource{
	// ...
}


read, err := client.DppResourceGuardProxyCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppResourceGuardProxiesClient.DppResourceGuardProxyDelete`

```go
ctx := context.TODO()
id := dppresourceguardproxies.NewBackupResourceGuardProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "backupResourceGuardProxyValue")

read, err := client.DppResourceGuardProxyDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppResourceGuardProxiesClient.DppResourceGuardProxyGet`

```go
ctx := context.TODO()
id := dppresourceguardproxies.NewBackupResourceGuardProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "backupResourceGuardProxyValue")

read, err := client.DppResourceGuardProxyGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppResourceGuardProxiesClient.DppResourceGuardProxyList`

```go
ctx := context.TODO()
id := dppresourceguardproxies.NewBackupVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue")

// alternatively `client.DppResourceGuardProxyList(ctx, id)` can be used to do batched pagination
items, err := client.DppResourceGuardProxyListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DppResourceGuardProxiesClient.DppResourceGuardProxyUnlockDelete`

```go
ctx := context.TODO()
id := dppresourceguardproxies.NewBackupResourceGuardProxyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "backupVaultValue", "backupResourceGuardProxyValue")

payload := dppresourceguardproxies.UnlockDeleteRequest{
	// ...
}


read, err := client.DppResourceGuardProxyUnlockDelete(ctx, id, payload, dppresourceguardproxies.DefaultDppResourceGuardProxyUnlockDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
