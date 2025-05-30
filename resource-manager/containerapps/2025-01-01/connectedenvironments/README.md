
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-01-01/connectedenvironments` Documentation

The `connectedenvironments` SDK allows for interaction with Azure Resource Manager `containerapps` (API Version `2025-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-01-01/connectedenvironments"
```


### Client Initialization

```go
client := connectedenvironments.NewConnectedEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesCreateOrUpdate`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "certificateName")

payload := connectedenvironments.Certificate{
	// ...
}


read, err := client.CertificatesCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesDelete`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "certificateName")

read, err := client.CertificatesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesGet`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "certificateName")

read, err := client.CertificatesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesList`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

// alternatively `client.CertificatesList(ctx, id)` can be used to do batched pagination
items, err := client.CertificatesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesUpdate`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName", "certificateName")

payload := connectedenvironments.CertificatePatch{
	// ...
}


read, err := client.CertificatesUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.CheckNameAvailability`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

payload := connectedenvironments.CheckNameAvailabilityRequest{
	// ...
}


read, err := client.CheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

payload := connectedenvironments.ConnectedEnvironment{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectedEnvironmentsClient.Delete`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectedEnvironmentsClient.Get`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConnectedEnvironmentsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConnectedEnvironmentsClient.ListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConnectedEnvironmentsClient.Update`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentName")

read, err := client.Update(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
