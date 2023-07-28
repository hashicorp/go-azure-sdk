
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/connectedenvironments` Documentation

The `connectedenvironments` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2022-11-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/connectedenvironments"
```


### Client Initialization

```go
client := connectedenvironments.NewConnectedEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConnectedEnvironmentsClient.CertificatesCreateOrUpdate`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "certificateValue")

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
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "certificateValue")

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
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "certificateValue")

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
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

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
id := connectedenvironments.NewConnectedEnvironmentCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue", "certificateValue")

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
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

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
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

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
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `ConnectedEnvironmentsClient.Get`

```go
ctx := context.TODO()
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

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
id := connectedenvironments.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

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
id := connectedenvironments.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

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
id := connectedenvironments.NewConnectedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "connectedEnvironmentValue")

read, err := client.Update(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
