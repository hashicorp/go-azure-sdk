
## `github.com/hashicorp/go-azure-sdk/resource-manager/attestation/2021-06-01/attestations` Documentation

The `attestations` SDK allows for interaction with Azure Resource Manager `attestation` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/attestation/2021-06-01/attestations"
```


### Client Initialization

```go
client := attestations.NewAttestationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AttestationsClient.AttestationProvidersGetDefaultByLocation`

```go
ctx := context.TODO()
id := attestations.NewLocationID("12345678-1234-9876-4563-123456789012", "locationName")

read, err := client.AttestationProvidersGetDefaultByLocation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AttestationsClient.AttestationProvidersListDefault`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.AttestationProvidersListDefault(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
