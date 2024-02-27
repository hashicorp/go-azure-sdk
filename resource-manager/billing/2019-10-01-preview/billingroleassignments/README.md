
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2019-10-01-preview/billingroleassignments` Documentation

The `billingroleassignments` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2019-10-01-preview`).

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
id := billingroleassignments.NewBillingAccountID("billingAccountValue")

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
id := billingroleassignments.NewBillingProfileID("billingAccountValue", "billingProfileValue")

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
id := billingroleassignments.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

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
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewInvoiceSectionBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewInvoiceSectionBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleAssignmentValue")

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
id := billingroleassignments.NewBillingAccountID("billingAccountValue")

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
id := billingroleassignments.NewBillingProfileID("billingAccountValue", "billingProfileValue")

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
id := billingroleassignments.NewDepartmentID("billingAccountValue", "departmentValue")

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
id := billingroleassignments.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

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
id := billingroleassignments.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

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
id := billingroleassignments.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

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
