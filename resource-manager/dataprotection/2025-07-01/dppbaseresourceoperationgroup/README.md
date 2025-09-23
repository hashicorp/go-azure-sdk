
## `github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/dppbaseresourceoperationgroup` Documentation

The `dppbaseresourceoperationgroup` SDK allows for interaction with Azure Resource Manager `dataprotection` (API Version `2025-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/dataprotection/2025-07-01/dppbaseresourceoperationgroup"
```


### Client Initialization

```go
client := dppbaseresourceoperationgroup.NewDppBaseResourceOperationGroupClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetBackupSecurityPINRequestsObjects`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.ResourceGuardsGetBackupSecurityPINRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.ResourceGuardsGetBackupSecurityPINRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDefaultBackupSecurityPINRequestsObject`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewGetBackupSecurityPINRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "getBackupSecurityPINRequestName")

read, err := client.ResourceGuardsGetDefaultBackupSecurityPINRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDefaultDeleteProtectedItemRequestsObject`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewDeleteProtectedItemRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "deleteProtectedItemRequestName")

read, err := client.ResourceGuardsGetDefaultDeleteProtectedItemRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDefaultDisableSoftDeleteRequestsObject`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewDisableSoftDeleteRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "disableSoftDeleteRequestName")

read, err := client.ResourceGuardsGetDefaultDisableSoftDeleteRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDefaultUpdateProtectedItemRequestsObject`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewUpdateProtectedItemRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "updateProtectedItemRequestName")

read, err := client.ResourceGuardsGetDefaultUpdateProtectedItemRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDefaultUpdateProtectionPolicyRequestsObject`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewUpdateProtectionPolicyRequestID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName", "updateProtectionPolicyRequestName")

read, err := client.ResourceGuardsGetDefaultUpdateProtectionPolicyRequestsObject(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDeleteProtectedItemRequestsObjects`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.ResourceGuardsGetDeleteProtectedItemRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.ResourceGuardsGetDeleteProtectedItemRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetDisableSoftDeleteRequestsObjects`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.ResourceGuardsGetDisableSoftDeleteRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.ResourceGuardsGetDisableSoftDeleteRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetUpdateProtectedItemRequestsObjects`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.ResourceGuardsGetUpdateProtectedItemRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.ResourceGuardsGetUpdateProtectedItemRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DppBaseResourceOperationGroupClient.ResourceGuardsGetUpdateProtectionPolicyRequestsObjects`

```go
ctx := context.TODO()
id := dppbaseresourceoperationgroup.NewResourceGuardID("12345678-1234-9876-4563-123456789012", "example-resource-group", "resourceGuardName")

// alternatively `client.ResourceGuardsGetUpdateProtectionPolicyRequestsObjects(ctx, id)` can be used to do batched pagination
items, err := client.ResourceGuardsGetUpdateProtectionPolicyRequestsObjectsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
