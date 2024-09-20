
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroleassignments` Documentation

The `billingroleassignments` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2019-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroleassignments"
```


### Client Initialization

```go
client := billingroleassignments.NewBillingRoleAssignmentsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRoleAssignmentsClient.AddByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingAccountID("billingAccountName")

payload := billingroleassignments.BillingRoleAssignmentPayload{
	// ...
}


// alternatively `client.AddByBillingAccount(ctx, id, payload)` can be used to do batched pagination
items, err := client.AddByBillingAccountComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentsClient.AddByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingProfileID("billingAccountName", "billingProfileName")

payload := billingroleassignments.BillingRoleAssignmentPayload{
	// ...
}


// alternatively `client.AddByBillingProfile(ctx, id, payload)` can be used to do batched pagination
items, err := client.AddByBillingProfileComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentsClient.AddByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignments.NewInvoiceSectionID("billingAccountName", "billingProfileName", "invoiceSectionName")

payload := billingroleassignments.BillingRoleAssignmentPayload{
	// ...
}


// alternatively `client.AddByInvoiceSection(ctx, id, payload)` can be used to do batched pagination
items, err := client.AddByInvoiceSectionComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
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


### Example Usage: `BillingRoleAssignmentsClient.DeleteByDepartment`

```go
ctx := context.TODO()
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountName", "departmentName", "billingRoleAssignmentName")

read, err := client.DeleteByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.DeleteByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountName", "enrollmentAccountName", "billingRoleAssignmentName")

read, err := client.DeleteByEnrollmentAccount(ctx, id)
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


### Example Usage: `BillingRoleAssignmentsClient.EnrollmentAccountRoleAssignmentsPut`

```go
ctx := context.TODO()
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountName", "enrollmentAccountName", "billingRoleAssignmentName")

payload := billingroleassignments.BillingRoleAssignment{
	// ...
}


read, err := client.EnrollmentAccountRoleAssignmentsPut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.EnrollmentDepartmentRoleAssignmentsPut`

```go
ctx := context.TODO()
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountName", "departmentName", "billingRoleAssignmentName")

payload := billingroleassignments.BillingRoleAssignment{
	// ...
}


read, err := client.EnrollmentDepartmentRoleAssignmentsPut(ctx, id, payload)
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


### Example Usage: `BillingRoleAssignmentsClient.GetByDepartment`

```go
ctx := context.TODO()
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountName", "departmentName", "billingRoleAssignmentName")

read, err := client.GetByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentsClient.GetByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountName", "enrollmentAccountName", "billingRoleAssignmentName")

read, err := client.GetByEnrollmentAccount(ctx, id)
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


### Example Usage: `BillingRoleAssignmentsClient.ListByDepartment`

```go
ctx := context.TODO()
id := billingroleassignments.NewDepartmentID("billingAccountName", "departmentName")

// alternatively `client.ListByDepartment(ctx, id)` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentsClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignments.NewEnrollmentAccountID("billingAccountName", "enrollmentAccountName")

// alternatively `client.ListByEnrollmentAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id)
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


### Example Usage: `BillingRoleAssignmentsClient.RoleAssignmentsPut`

```go
ctx := context.TODO()
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountName", "billingRoleAssignmentName")

payload := billingroleassignments.BillingRoleAssignment{
	// ...
}


read, err := client.RoleAssignmentsPut(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
