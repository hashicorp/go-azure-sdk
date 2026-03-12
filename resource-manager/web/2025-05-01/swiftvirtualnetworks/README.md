
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworks` Documentation

The `swiftvirtualnetworks` SDK allows for interaction with Azure Resource Manager `web` (API Version `2025-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2025-05-01/swiftvirtualnetworks"
```


### Client Initialization

```go
client := swiftvirtualnetworks.NewSwiftVirtualNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SwiftVirtualNetworksClient.WebAppsCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := swiftvirtualnetworks.SwiftVirtualNetwork{
	// ...
}


read, err := client.WebAppsCreateOrUpdateSwiftVirtualNetworkConnectionWithCheck(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworksClient.WebAppsDeleteSwiftVirtualNetwork`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsDeleteSwiftVirtualNetwork(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworksClient.WebAppsGetSwiftVirtualNetworkConnection`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

read, err := client.WebAppsGetSwiftVirtualNetworkConnection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SwiftVirtualNetworksClient.WebAppsUpdateSwiftVirtualNetworkConnectionWithCheck`

```go
ctx := context.TODO()
id := commonids.NewAppServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

payload := swiftvirtualnetworks.SwiftVirtualNetwork{
	// ...
}


read, err := client.WebAppsUpdateSwiftVirtualNetworkConnectionWithCheck(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
