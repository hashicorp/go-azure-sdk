
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroledefinitions` Documentation

The `billingroledefinitions` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

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
id := billingroledefinitions.NewBillingRoleDefinitionID("billingAccountName", "billingRoleDefinitionName")

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
id := billingroledefinitions.NewBillingProfileBillingRoleDefinitionID("billingAccountName", "billingProfileName", "billingRoleDefinitionName")

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
id := billingroledefinitions.NewDepartmentBillingRoleDefinitionID("billingAccountName", "departmentName", "billingRoleDefinitionName")

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
id := billingroledefinitions.NewEnrollmentAccountBillingRoleDefinitionID("billingAccountName", "enrollmentAccountName", "billingRoleDefinitionName")

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
id := billingroledefinitions.NewInvoiceSectionBillingRoleDefinitionID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingRoleDefinitionName")

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
id := billingroledefinitions.NewBillingAccountID("billingAccountName")

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
id := billingroledefinitions.NewBillingProfileID("billingAccountName", "billingProfileName")

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
id := billingroledefinitions.NewDepartmentID("billingAccountName", "departmentName")

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
id := billingroledefinitions.NewEnrollmentAccountID("billingAccountName", "enrollmentAccountName")

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
id := billingroledefinitions.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
