
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroledefinitions` Documentation

The `billingroledefinitions` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroledefinitions"
```


### Client Initialization

```go
client := billingroledefinitions.NewBillingRoleDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRoleDefinitionsClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := billingroledefinitions.NewBillingRoleDefinitionID("billingAccountValue", "billingRoleDefinitionValue")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionsClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := billingroledefinitions.NewBillingProfileBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "billingRoleDefinitionValue")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionsClient.GetByDepartment`

```go
ctx := context.TODO()
id := billingroledefinitions.NewDepartmentBillingRoleDefinitionID("billingAccountValue", "departmentValue", "billingRoleDefinitionValue")

read, err := client.GetByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionsClient.GetByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroledefinitions.NewEnrollmentAccountBillingRoleDefinitionID("billingAccountValue", "enrollmentAccountValue", "billingRoleDefinitionValue")

read, err := client.GetByEnrollmentAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionsClient.GetByInvoiceSection`

```go
ctx := context.TODO()
id := billingroledefinitions.NewInvoiceSectionBillingRoleDefinitionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleDefinitionValue")

read, err := client.GetByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleDefinitionsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingroledefinitions.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingroledefinitions.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionsClient.ListByDepartment`

```go
ctx := context.TODO()
id := billingroledefinitions.NewDepartmentID("billingAccountValue", "departmentValue")

// alternatively `client.ListByDepartment(ctx, id)` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionsClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroledefinitions.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

// alternatively `client.ListByEnrollmentAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleDefinitionsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingroledefinitions.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
