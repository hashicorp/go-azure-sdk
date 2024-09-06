
## `github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/certificate` Documentation

The `certificate` SDK allows for interaction with the Azure Resource Manager Service `apimanagement` (API Version `2024-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/apimanagement/2024-05-01/certificate"
```


### Client Initialization

```go
client := certificate.NewCertificateClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CertificateClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := certificate.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "certificateIdValue")

payload := certificate.CertificateCreateOrUpdateParameters{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload, certificate.DefaultCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.Delete`

```go
ctx := context.TODO()
id := certificate.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "certificateIdValue")

read, err := client.Delete(ctx, id, certificate.DefaultDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.Get`

```go
ctx := context.TODO()
id := certificate.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "certificateIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.GetEntityTag`

```go
ctx := context.TODO()
id := certificate.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "certificateIdValue")

read, err := client.GetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.ListByService`

```go
ctx := context.TODO()
id := certificate.NewServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue")

// alternatively `client.ListByService(ctx, id, certificate.DefaultListByServiceOperationOptions())` can be used to do batched pagination
items, err := client.ListByServiceComplete(ctx, id, certificate.DefaultListByServiceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateClient.RefreshSecret`

```go
ctx := context.TODO()
id := certificate.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "certificateIdValue")

read, err := client.RefreshSecret(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateCreateOrUpdate`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "certificateIdValue")

payload := certificate.CertificateCreateOrUpdateParameters{
	// ...
}


read, err := client.WorkspaceCertificateCreateOrUpdate(ctx, id, payload, certificate.DefaultWorkspaceCertificateCreateOrUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateDelete`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "certificateIdValue")

read, err := client.WorkspaceCertificateDelete(ctx, id, certificate.DefaultWorkspaceCertificateDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateGet`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "certificateIdValue")

read, err := client.WorkspaceCertificateGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateGetEntityTag`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "certificateIdValue")

read, err := client.WorkspaceCertificateGetEntityTag(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateListByWorkspace`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue")

// alternatively `client.WorkspaceCertificateListByWorkspace(ctx, id, certificate.DefaultWorkspaceCertificateListByWorkspaceOperationOptions())` can be used to do batched pagination
items, err := client.WorkspaceCertificateListByWorkspaceComplete(ctx, id, certificate.DefaultWorkspaceCertificateListByWorkspaceOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `CertificateClient.WorkspaceCertificateRefreshSecret`

```go
ctx := context.TODO()
id := certificate.NewWorkspaceCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serviceValue", "workspaceIdValue", "certificateIdValue")

read, err := client.WorkspaceCertificateRefreshSecret(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
