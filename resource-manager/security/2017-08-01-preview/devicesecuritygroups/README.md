
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/devicesecuritygroups` Documentation

The `devicesecuritygroups` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/devicesecuritygroups"
```


### Client Initialization

```go
client := devicesecuritygroups.NewDeviceSecurityGroupsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DeviceSecurityGroupsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := devicesecuritygroups.NewScopedDeviceSecurityGroupID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "deviceSecurityGroupValue")

payload := devicesecuritygroups.DeviceSecurityGroup{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceSecurityGroupsClient.Delete`

```go
ctx := context.TODO()
id := devicesecuritygroups.NewScopedDeviceSecurityGroupID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "deviceSecurityGroupValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceSecurityGroupsClient.Get`

```go
ctx := context.TODO()
id := devicesecuritygroups.NewScopedDeviceSecurityGroupID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group", "deviceSecurityGroupValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DeviceSecurityGroupsClient.List`

```go
ctx := context.TODO()
id := devicesecuritygroups.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
