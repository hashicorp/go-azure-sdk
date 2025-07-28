
## `github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/openapis` Documentation

The `openapis` SDK allows for interaction with Azure Resource Manager `recoveryservices` (API Version `2025-02-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/recoveryservices/2025-02-01/openapis"
```


### Client Initialization

```go
client := openapis.NewOpenapisClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OpenapisClient.RecoveryServicesCapabilities`

```go
ctx := context.TODO()
id := openapis.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

payload := openapis.ResourceCapabilities{
	// ...
}


read, err := client.RecoveryServicesCapabilities(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RecoveryServicesCheckNameAvailability`

```go
ctx := context.TODO()
id := openapis.NewProviderLocationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "locationName")

payload := openapis.CheckNameAvailabilityParameters{
	// ...
}


read, err := client.RecoveryServicesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.RegisteredIdentitiesDelete`

```go
ctx := context.TODO()
id := openapis.NewRegisteredIdentityID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "registeredIdentityName")

read, err := client.RegisteredIdentitiesDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.ReplicationUsagesList`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.ReplicationUsagesList(ctx, id)` can be used to do batched pagination
items, err := client.ReplicationUsagesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.UsagesListByVaults`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

// alternatively `client.UsagesListByVaults(ctx, id)` can be used to do batched pagination
items, err := client.UsagesListByVaultsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.VaultCertificatesCreate`

```go
ctx := context.TODO()
id := openapis.NewCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName", "certificateName")

payload := openapis.CertificateRequest{
	// ...
}


read, err := client.VaultCertificatesCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.VaultsCreateOrUpdate`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := openapis.Vault{
	// ...
}


if err := client.VaultsCreateOrUpdateThenPoll(ctx, id, payload, openapis.DefaultVaultsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.VaultsDelete`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

if err := client.VaultsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `OpenapisClient.VaultsGet`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

read, err := client.VaultsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `OpenapisClient.VaultsListBySubscriptionId`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.VaultsListBySubscriptionId(ctx, id)` can be used to do batched pagination
items, err := client.VaultsListBySubscriptionIdComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `OpenapisClient.VaultsUpdate`

```go
ctx := context.TODO()
id := openapis.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultName")

payload := openapis.PatchVault{
	// ...
}


if err := client.VaultsUpdateThenPoll(ctx, id, payload, openapis.DefaultVaultsUpdateOperationOptions()); err != nil {
	// handle the error
}
```
