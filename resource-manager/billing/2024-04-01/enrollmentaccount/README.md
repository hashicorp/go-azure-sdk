
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/enrollmentaccount` Documentation

The `enrollmentaccount` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/enrollmentaccount"
```


### Client Initialization

```go
client := enrollmentaccount.NewEnrollmentAccountClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnrollmentAccountClient.Get`

```go
ctx := context.TODO()
id := enrollmentaccount.NewEnrollmentAccountID("billingAccountName", "enrollmentAccountName")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnrollmentAccountClient.GetByDepartment`

```go
ctx := context.TODO()
id := enrollmentaccount.NewDepartmentEnrollmentAccountID("billingAccountName", "departmentName", "enrollmentAccountName")

read, err := client.GetByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnrollmentAccountClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := enrollmentaccount.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id, enrollmentaccount.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, enrollmentaccount.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EnrollmentAccountClient.ListByDepartment`

```go
ctx := context.TODO()
id := enrollmentaccount.NewDepartmentID("billingAccountName", "departmentName")

// alternatively `client.ListByDepartment(ctx, id, enrollmentaccount.DefaultListByDepartmentOperationOptions())` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id, enrollmentaccount.DefaultListByDepartmentOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
