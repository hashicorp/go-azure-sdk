
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/portalconfig` Documentation

The `portalconfig` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2022-08-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2022-08-01/portalconfig"
```


### Client Initialization

```go
client := portalconfig.NewPortalConfigClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PortalConfigClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := portalconfig.NewPortalConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalConfigIdValue")

payload := portalconfig.PortalConfigContract{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, portalconfig.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalConfigClient.Get`

```go
ctx := context.TODO()
id := portalconfig.NewPortalConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalConfigIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalConfigClient.GetEntityTag`

```go
ctx := context.TODO()
id := portalconfig.NewPortalConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalConfigIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalConfigClient.ListByService`

```go
ctx := context.TODO()
id := portalconfig.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

read, err := client.ListByService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PortalConfigClient.Update`

```go
ctx := context.TODO()
id := portalconfig.NewPortalConfigID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "portalConfigIdValue")

payload := portalconfig.PortalConfigContract{
	// ...
}


read, err := client.Update(ctx, id, payload, portalconfig.DefaultUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
