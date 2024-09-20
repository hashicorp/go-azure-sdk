
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingroleassignments` Documentation

The `billingroleassignments` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingroleassignments"
```


### Client Initialization

```go
client := billingroleassignments.NewBillingRoleAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRoleAssignmentsClient.DeleteByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountName", "billingRoleAssignmentName")

read, err := client.DeleteByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.DeleteByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingProfileBillingRoleAssignmentID("billingAccountName", "billingProfileName", "billingRoleAssignmentName")

read, err := client.DeleteByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.DeleteByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignments.NewInvoiceSectionBillingRoleAssignmentID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingRoleAssignmentName")

read, err := client.DeleteByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountName", "billingRoleAssignmentName")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingProfileBillingRoleAssignmentID("billingAccountName", "billingProfileName", "billingRoleAssignmentName")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.GetByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignments.NewInvoiceSectionBillingRoleAssignmentID("billingAccountName", "billingProfileName", "invoiceSectionName", "billingRoleAssignmentName")

read, err := client.GetByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingAccountID("billingAccountName")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingProfileID("billingAccountName", "billingProfileName")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentsClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignments.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
