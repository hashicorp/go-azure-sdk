
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/hcxenterprisesites` Documentation

The `hcxenterprisesites` SDK allows for interaction with Azure Resource Manager `vmware` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2023-09-01/hcxenterprisesites"
```


### Client Initialization

```go
client := hcxenterprisesites.NewHcxEnterpriseSitesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `HcxEnterpriseSitesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := hcxenterprisesites.NewHcxEnterpriseSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "hcxEnterpriseSiteValue")

payload := hcxenterprisesites.HcxEnterpriseSite{
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


### Example Usage: `HcxEnterpriseSitesClient.Delete`

```go
ctx := context.TODO()
id := hcxenterprisesites.NewHcxEnterpriseSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "hcxEnterpriseSiteValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HcxEnterpriseSitesClient.Get`

```go
ctx := context.TODO()
id := hcxenterprisesites.NewHcxEnterpriseSiteID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "hcxEnterpriseSiteValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `HcxEnterpriseSitesClient.List`

```go
ctx := context.TODO()
id := hcxenterprisesites.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
