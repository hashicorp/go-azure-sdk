
## `github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/confluents` Documentation

The `confluents` SDK allows for interaction with Azure Resource Manager `confluent` (API Version `2024-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/confluent/2024-07-01/confluents"
```


### Client Initialization

```go
client := confluents.NewConfluentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ConfluentsClient.AccessDeleteRoleBinding`

```go
ctx := context.TODO()
id := confluents.NewDeleteRoleBindingID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "roleBindingId")

read, err := client.AccessDeleteRoleBinding(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.MarketplaceAgreementsCreate`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := confluents.ConfluentAgreementResource{
	// ...
}


read, err := client.MarketplaceAgreementsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.MarketplaceAgreementsList`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.MarketplaceAgreementsList(ctx, id)` can be used to do batched pagination
items, err := client.MarketplaceAgreementsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ConfluentsClient.OrganizationDeleteClusterAPIKey`

```go
ctx := context.TODO()
id := confluents.NewApiKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "apiKeyId")

read, err := client.OrganizationDeleteClusterAPIKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.OrganizationGetClusterAPIKey`

```go
ctx := context.TODO()
id := confluents.NewApiKeyID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "apiKeyId")

read, err := client.OrganizationGetClusterAPIKey(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.OrganizationGetSchemaRegistryClusterById`

```go
ctx := context.TODO()
id := confluents.NewSchemaRegistryClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "organizationName", "environmentId", "clusterId")

read, err := client.OrganizationGetSchemaRegistryClusterById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.ValidationsValidateOrganization`

```go
ctx := context.TODO()
id := confluents.NewValidationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "validationName")

payload := confluents.OrganizationResource{
	// ...
}


read, err := client.ValidationsValidateOrganization(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ConfluentsClient.ValidationsValidateOrganizationV2`

```go
ctx := context.TODO()
id := confluents.NewValidationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "validationName")

payload := confluents.OrganizationResource{
	// ...
}


read, err := client.ValidationsValidateOrganizationV2(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
