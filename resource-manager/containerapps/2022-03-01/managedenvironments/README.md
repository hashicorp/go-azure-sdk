
## `github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/managedenvironments` Documentation

The `managedenvironments` SDK allows for interaction with the Azure Resource Manager Service `containerapps` (API Version `2022-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/managedenvironments"
```


### Client Initialization

```go
client := managedenvironments.NewManagedEnvironmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `ManagedEnvironmentsClient.CertificatesCreateOrUpdate`

```go
ctx := context.TODO()
id := managedenvironments.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue", "certificateValue")

payload := managedenvironments.Certificate{
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


### Example Usage: `ManagedEnvironmentsClient.CertificatesDelete`

```go
ctx := context.TODO()
id := managedenvironments.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue", "certificateValue")
read, err := client.CertificatesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedEnvironmentsClient.CertificatesGet`

```go
ctx := context.TODO()
id := managedenvironments.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue", "certificateValue")
read, err := client.CertificatesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedEnvironmentsClient.CertificatesList`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")
// alternatively `client.CertificatesList(ctx, id)` can be used to do batched pagination
items, err := client.CertificatesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedEnvironmentsClient.CertificatesUpdate`

```go
ctx := context.TODO()
id := managedenvironments.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue", "certificateValue")

payload := managedenvironments.CertificatePatch{
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


### Example Usage: `ManagedEnvironmentsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")

payload := managedenvironments.ManagedEnvironment{
	// ...
}

future, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedEnvironmentsClient.Delete`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")
future, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `ManagedEnvironmentsClient.Get`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")
read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedEnvironmentsClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := managedenvironments.NewResourceGroupID()
// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedEnvironmentsClient.ListBySubscription`

```go
ctx := context.TODO()
id := managedenvironments.NewSubscriptionID()
// alternatively `client.ListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ManagedEnvironmentsClient.NamespacesCheckNameAvailability`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")

payload := managedenvironments.CheckNameAvailabilityRequest{
	// ...
}

read, err := client.NamespacesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagedEnvironmentsClient.Update`

```go
ctx := context.TODO()
id := managedenvironments.NewManagedEnvironmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "environmentValue")

payload := managedenvironments.ManagedEnvironment{
	// ...
}

future, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```
