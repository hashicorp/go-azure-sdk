
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindings` Documentation

The `hostnamebindings` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/hostnamebindings"
```


### Client Initialization

```go
client := hostnamebindings.NewHostNameBindingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HostNameBindingsClient.WebAppsCreateOrUpdateHostNameBinding`

```go
ctx := context.TODO()
id := hostnamebindings.NewHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hostNameBindingName")

payload := hostnamebindings.HostNameBinding{
	// ...
}


read, err := client.WebAppsCreateOrUpdateHostNameBinding(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingsClient.WebAppsDeleteHostNameBinding`

```go
ctx := context.TODO()
id := hostnamebindings.NewHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hostNameBindingName")

read, err := client.WebAppsDeleteHostNameBinding(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingsClient.WebAppsGetHostNameBinding`

```go
ctx := context.TODO()
id := hostnamebindings.NewHostNameBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "hostNameBindingName")

read, err := client.WebAppsGetHostNameBinding(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HostNameBindingsClient.WebAppsListHostNameBindings`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.WebAppsListHostNameBindings(ctx, id)` can be used to do batched pagination
items, err := client.WebAppsListHostNameBindingsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
