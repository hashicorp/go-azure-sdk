
## `github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutionsreferencedata` Documentation

The `securitysolutionsreferencedata` SDK allows for interaction with Azure Resource Manager `security` (API Version `2020-01-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/security/2020-01-01/securitysolutionsreferencedata"
```


### Client Initialization

```go
client := securitysolutionsreferencedata.NewSecuritySolutionsReferenceDataClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SecuritySolutionsReferenceDataClient.List`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SecuritySolutionsReferenceDataClient.ListByHomeRegion`

```go
ctx := context.TODO()
id := securitysolutionsreferencedata.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.ListByHomeRegion(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
