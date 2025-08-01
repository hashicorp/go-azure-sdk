
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/resourceguards` Documentation

The `resourceguards` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/resourceguards"
```


### Client Initialization

```go
client := resourceguards.NewResourceGuardsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ResourceGuardsClient.Delete`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.Get`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetBackupSecurityPINRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetBackupSecurityPINRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetBackupSecurityPINRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultBackupSecurityPINRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewGetBackupSecurityPINRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "getBackupSecurityPINRequestName")

read, err := client.GetDefaultBackupSecurityPINRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultDeleteProtectedItemRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewDeleteProtectedItemRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "deleteProtectedItemRequestName")

read, err := client.GetDefaultDeleteProtectedItemRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultDeleteResourceGuardProxyRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewDeleteResourceGuardProxyRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "deleteResourceGuardProxyRequestName")

read, err := client.GetDefaultDeleteResourceGuardProxyRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultDisableSoftDeleteRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewDisableSoftDeleteRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "disableSoftDeleteRequestName")

read, err := client.GetDefaultDisableSoftDeleteRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultUpdateProtectedItemRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewUpdateProtectedItemRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "updateProtectedItemRequestName")

read, err := client.GetDefaultUpdateProtectedItemRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDefaultUpdateProtectionPolicyRequestsObject`

```go
ctx := context.TODO()
id := resourceguards.NewUpdateProtectionPolicyRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "updateProtectionPolicyRequestName")

read, err := client.GetDefaultUpdateProtectionPolicyRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.GetDeleteProtectedItemRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetDeleteProtectedItemRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetDeleteProtectedItemRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetDeleteResourceGuardProxyRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetDeleteResourceGuardProxyRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetDeleteResourceGuardProxyRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetDisableSoftDeleteRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetDisableSoftDeleteRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetDisableSoftDeleteRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetResourcesInResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.GetResourcesInResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.GetResourcesInResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetResourcesInSubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.GetResourcesInSubscription(ctx, id)` can be used to do batched pagination
items, err := client.GetResourcesInSubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetUpdateProtectedItemRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetUpdateProtectedItemRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetUpdateProtectedItemRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.GetUpdateProtectionPolicyRequestsObjects`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.GetUpdateProtectionPolicyRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.GetUpdateProtectionPolicyRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ResourceGuardsClient.Patch`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

payload := resourceguards.PatchResourceGuardInput{
	// ...
}


read, err := client.Patch(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ResourceGuardsClient.Put`

```go
ctx := context.TODO()
id := resourceguards.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

payload := resourceguards.ResourceGuardResource{
	// ...
}


read, err := client.Put(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
