
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/tenantattachrbac` Documentation

The `tenantattachrbac` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/tenantattachrbac"
```


### Client Initialization

```go
client := tenantattachrbac.NewTenantAttachRBACClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TenantAttachRBACClient.DeleteTenantAttachRBAC`

```go
ctx := context.TODO()


read, err := client.DeleteTenantAttachRBAC(ctx, tenantattachrbac.DefaultDeleteTenantAttachRBACOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TenantAttachRBACClient.EnableTenantAttachRBAC`

```go
ctx := context.TODO()

payload := tenantattachrbac.EnableTenantAttachRBACRequest{
	// ...
}


read, err := client.EnableTenantAttachRBAC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TenantAttachRBACClient.GetTenantAttachRBAC`

```go
ctx := context.TODO()


read, err := client.GetTenantAttachRBAC(ctx, tenantattachrbac.DefaultGetTenantAttachRBACOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TenantAttachRBACClient.UpdateTenantAttachRBAC`

```go
ctx := context.TODO()

payload := tenantattachrbac.TenantAttachRBAC{
	// ...
}


read, err := client.UpdateTenantAttachRBAC(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
