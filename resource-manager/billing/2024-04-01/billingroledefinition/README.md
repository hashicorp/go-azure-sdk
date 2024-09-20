
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroledefinition` Documentation

The `billingroledefinition` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroledefinition"
```


### Client Initialization

```go
client := billingroledefinition.NewBillingRoleDefinitionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRoleDefinitionClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := billingroledefinition.NewBillingRoleDefinitionID("billingAccountName", "roleDefinitionName")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := billingroledefinition.NewBillingProfileBillingRoleDefinitionID("billingAccountName", "billingProfileName", "roleDefinitionName")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.GetByCustomer`

```go
ctx := context.TODO()
id := billingroledefinition.NewCustomerBillingRoleDefinitionID("billingAccountName", "billingProfileName", "customerName", "roleDefinitionName")

read, err := client.GetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.GetByDepartment`

```go
ctx := context.TODO()
id := billingroledefinition.NewDepartmentBillingRoleDefinitionID("billingAccountName", "departmentName", "roleDefinitionName")

read, err := client.GetByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.GetByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroledefinition.NewEnrollmentAccountBillingRoleDefinitionID("billingAccountName", "enrollmentAccountName", "roleDefinitionName")

read, err := client.GetByEnrollmentAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.GetByInvoiceSection`

```go
ctx := context.TODO()
id := billingroledefinition.NewInvoiceSectionBillingRoleDefinitionID("billingAccountName", "billingProfileName", "invoiceSectionName", "roleDefinitionName")

read, err := client.GetByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingroledefinition.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingroledefinition.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingroledefinition.NewBillingProfileCustomerID("billingAccountName", "billingProfileName", "customerName")

// alternatively `client.ListByCustomer(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByDepartment`

```go
ctx := context.TODO()
id := billingroledefinition.NewDepartmentID("billingAccountName", "departmentName")

// alternatively `client.ListByDepartment(ctx, id)` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroledefinition.NewEnrollmentAccountID("billingAccountName", "enrollmentAccountName")

// alternatively `client.ListByEnrollmentAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingroledefinition.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
