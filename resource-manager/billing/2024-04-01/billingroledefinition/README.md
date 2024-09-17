
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
id := billingroledefinition.NewBillingRoleDefinitionID("billingAccountValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewBillingProfileBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewCustomerBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "customerValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewDepartmentBillingRoleDefinitionID("billingAccountValue", "departmentValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewEnrollmentAccountBillingRoleDefinitionID("billingAccountValue", "enrollmentAccountValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewInvoiceSectionBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleDefinitionValue")

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
id := billingroledefinition.NewBillingAccountID("billingAccountValue")

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
id := billingroledefinition.NewBillingProfileID("billingAccountValue", "billingProfileValue")

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
id := billingroledefinition.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

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
id := billingroledefinition.NewDepartmentID("billingAccountValue", "departmentValue")

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
id := billingroledefinition.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

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
id := billingroledefinition.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
