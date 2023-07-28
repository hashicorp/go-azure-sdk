
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedcertificates` Documentation

The `managedcertificates` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedcertificates"
```


### Client Initialization

```go
client := managedcertificates.NewManagedCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedCertificatesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedcertificates.NewManagedCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "managedCertificateValue")

payload := managedcertificates.ManagedCertificate{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedCertificatesClient.Delete`

```go
ctx := context.TODO()
id := managedcertificates.NewManagedCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "managedCertificateValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedCertificatesClient.Get`

```go
ctx := context.TODO()
id := managedcertificates.NewManagedCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "managedCertificateValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedCertificatesClient.List`

```go
ctx := context.TODO()
id := managedcertificates.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedCertificatesClient.Update`

```go
ctx := context.TODO()
id := managedcertificates.NewManagedCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedEnvironmentValue", "managedCertificateValue")

payload := managedcertificates.ManagedCertificatePatch{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
