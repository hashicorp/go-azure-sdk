
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/managedinstancetdecertificates` Documentation

The `managedinstancetdecertificates` SDK allows for interaction with Azure Resource Manager `sql` (API Version `2023-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2023-08-01-preview/managedinstancetdecertificates"
```


### Client Initialization

```go
client := managedinstancetdecertificates.NewManagedInstanceTdeCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagedInstanceTdeCertificatesClient.Create`

```go
ctx := context.TODO()
id := commonids.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

payload := managedinstancetdecertificates.TdeCertificate{
	// ...
}


if err := client.CreateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
