
## `github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/vcenter` Documentation

The `vcenter` SDK allows for interaction with Azure Resource Manager `migrate` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/vcenter"
```


### Client Initialization

```go
client := vcenter.NewVCenterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VCenterClient.DeleteVCenter`

```go
ctx := context.TODO()
id := vcenter.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "vcenterName")

read, err := client.DeleteVCenter(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VCenterClient.GetAllVCentersInSite`

```go
ctx := context.TODO()
id := vcenter.NewVMwareSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName")

// alternatively `client.GetAllVCentersInSite(ctx, id, vcenter.DefaultGetAllVCentersInSiteOperationOptions())` can be used to do batched pagination
items, err := client.GetAllVCentersInSiteComplete(ctx, id, vcenter.DefaultGetAllVCentersInSiteOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `VCenterClient.GetVCenter`

```go
ctx := context.TODO()
id := vcenter.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "vcenterName")

read, err := client.GetVCenter(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VCenterClient.PutVCenter`

```go
ctx := context.TODO()
id := vcenter.NewVCenterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteName", "vcenterName")

payload := vcenter.VCenter{
	// ...
}


read, err := client.PutVCenter(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
