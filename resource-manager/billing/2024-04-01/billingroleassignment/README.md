
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroleassignment` Documentation

The `billingroleassignment` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingroleassignment"
```


### Client Initialization

```go
client := billingroleassignment.NewBillingRoleAssignmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingRoleAssignmentClient.CreateByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingAccountID("billingAccountValue")

payload := billingroleassignment.BillingRoleAssignmentProperties{
	// ...
}


if err := client.CreateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingroleassignment.BillingRoleAssignmentProperties{
	// ...
}


if err := client.CreateByBillingProfileThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateByCustomer`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

payload := billingroleassignment.BillingRoleAssignmentProperties{
	// ...
}


if err := client.CreateByCustomerThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignment.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

payload := billingroleassignment.BillingRoleAssignmentProperties{
	// ...
}


if err := client.CreateByInvoiceSectionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateOrUpdateByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

payload := billingroleassignment.BillingRoleAssignment{
	// ...
}


if err := client.CreateOrUpdateByBillingAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateOrUpdateByDepartment`

```go
ctx := context.TODO()
id := billingroleassignment.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

payload := billingroleassignment.BillingRoleAssignment{
	// ...
}


if err := client.CreateOrUpdateByDepartmentThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.CreateOrUpdateByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

payload := billingroleassignment.BillingRoleAssignment{
	// ...
}


if err := client.CreateOrUpdateByEnrollmentAccountThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

read, err := client.DeleteByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue")

read, err := client.DeleteByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByCustomer`

```go
ctx := context.TODO()
id := billingroleassignment.NewCustomerBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "customerValue", "billingRoleAssignmentValue")

read, err := client.DeleteByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByDepartment`

```go
ctx := context.TODO()
id := billingroleassignment.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

read, err := client.DeleteByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

read, err := client.DeleteByEnrollmentAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.DeleteByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignment.NewInvoiceSectionBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleAssignmentValue")

read, err := client.DeleteByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingRoleAssignmentID("billingAccountValue", "billingRoleAssignmentValue")

read, err := client.GetByBillingAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "billingRoleAssignmentValue")

read, err := client.GetByBillingProfile(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByCustomer`

```go
ctx := context.TODO()
id := billingroleassignment.NewCustomerBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "customerValue", "billingRoleAssignmentValue")

read, err := client.GetByCustomer(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByDepartment`

```go
ctx := context.TODO()
id := billingroleassignment.NewDepartmentBillingRoleAssignmentID("billingAccountValue", "departmentValue", "billingRoleAssignmentValue")

read, err := client.GetByDepartment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewEnrollmentAccountBillingRoleAssignmentID("billingAccountValue", "enrollmentAccountValue", "billingRoleAssignmentValue")

read, err := client.GetByEnrollmentAccount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.GetByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignment.NewInvoiceSectionBillingRoleAssignmentID("billingAccountValue", "billingProfileValue", "invoiceSectionValue", "billingRoleAssignmentValue")

read, err := client.GetByInvoiceSection(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id, billingroleassignment.DefaultListByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id, billingroleassignment.DefaultListByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id, billingroleassignment.DefaultListByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id, billingroleassignment.DefaultListByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id, billingroleassignment.DefaultListByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id, billingroleassignment.DefaultListByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByDepartment`

```go
ctx := context.TODO()
id := billingroleassignment.NewDepartmentID("billingAccountValue", "departmentValue")

// alternatively `client.ListByDepartment(ctx, id)` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

// alternatively `client.ListByEnrollmentAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignment.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id, billingroleassignment.DefaultListByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id, billingroleassignment.DefaultListByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ResolveByBillingAccount`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingAccountID("billingAccountValue")

// alternatively `client.ResolveByBillingAccount(ctx, id, billingroleassignment.DefaultResolveByBillingAccountOperationOptions())` can be used to do batched pagination
items, err := client.ResolveByBillingAccountComplete(ctx, id, billingroleassignment.DefaultResolveByBillingAccountOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ResolveByBillingProfile`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ResolveByBillingProfile(ctx, id, billingroleassignment.DefaultResolveByBillingProfileOperationOptions())` can be used to do batched pagination
items, err := client.ResolveByBillingProfileComplete(ctx, id, billingroleassignment.DefaultResolveByBillingProfileOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ResolveByCustomer`

```go
ctx := context.TODO()
id := billingroleassignment.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ResolveByCustomer(ctx, id, billingroleassignment.DefaultResolveByCustomerOperationOptions())` can be used to do batched pagination
items, err := client.ResolveByCustomerComplete(ctx, id, billingroleassignment.DefaultResolveByCustomerOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingRoleAssignmentClient.ResolveByInvoiceSection`

```go
ctx := context.TODO()
id := billingroleassignment.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ResolveByInvoiceSection(ctx, id, billingroleassignment.DefaultResolveByInvoiceSectionOperationOptions())` can be used to do batched pagination
items, err := client.ResolveByInvoiceSectionComplete(ctx, id, billingroleassignment.DefaultResolveByInvoiceSectionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
