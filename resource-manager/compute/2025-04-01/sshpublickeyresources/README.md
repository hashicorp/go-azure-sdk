
## `github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/sshpublickeyresources` Documentation

The `sshpublickeyresources` SDK allows for interaction with Azure Resource Manager `compute` (API Version `2025-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/compute/2025-04-01/sshpublickeyresources"
```


### Client Initialization

```go
client := sshpublickeyresources.NewSshPublicKeyResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysCreate`

```go
ctx := context.TODO()
id := sshpublickeyresources.NewSshPublicKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sshPublicKeyName")

payload := sshpublickeyresources.SshPublicKeyResource{
	// ...
}


read, err := client.SshPublicKeysCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysDelete`

```go
ctx := context.TODO()
id := sshpublickeyresources.NewSshPublicKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sshPublicKeyName")

read, err := client.SshPublicKeysDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysGenerateKeyPair`

```go
ctx := context.TODO()
id := sshpublickeyresources.NewSshPublicKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sshPublicKeyName")

payload := sshpublickeyresources.SshGenerateKeyPairInputParameters{
	// ...
}


read, err := client.SshPublicKeysGenerateKeyPair(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysGet`

```go
ctx := context.TODO()
id := sshpublickeyresources.NewSshPublicKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sshPublicKeyName")

read, err := client.SshPublicKeysGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.SshPublicKeysListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.SshPublicKeysListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.SshPublicKeysListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.SshPublicKeysListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SshPublicKeyResourcesClient.SshPublicKeysUpdate`

```go
ctx := context.TODO()
id := sshpublickeyresources.NewSshPublicKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "sshPublicKeyName")

payload := sshpublickeyresources.SshPublicKeyUpdateResource{
	// ...
}


read, err := client.SshPublicKeysUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
