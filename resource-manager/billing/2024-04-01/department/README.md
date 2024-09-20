
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/department` Documentation

The `department` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/department"
```


### Client Initialization

```go
client := department.NewDepartmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DepartmentClient.Get`

```go
ctx := context.TODO()
id := department.NewDepartmentID("billingAccountName", "departmentName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DepartmentClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := department.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, department.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, department.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
