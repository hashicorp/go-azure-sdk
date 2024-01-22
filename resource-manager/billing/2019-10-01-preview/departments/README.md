
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/departments` Documentation

The `departments` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/departments"
```


### Client Initialization

```go
client := departments.NewDepartmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DepartmentsClient.Get`

```go
ctx := context.TODO()
id := departments.NewDepartmentID("billingAccountValue", "departmentValue")

read, err := client.Get(ctx, id, departments.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepartmentsClient.ListByBillingAccountName`

```go
ctx := context.TODO()
id := departments.NewBillingAccountID("billingAccountValue")

read, err := client.ListByBillingAccountName(ctx, id, departments.DefaultListByBillingAccountNameOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
