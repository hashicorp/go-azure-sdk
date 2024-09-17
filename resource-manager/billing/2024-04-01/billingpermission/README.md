
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingpermission` Documentation

The `billingpermission` SDK allows for interaction with Azure Resource Manager `billing` (API Version `2024-04-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2024-04-01/billingpermission"
```


### Client Initialization

```go
client := billingpermission.NewBillingPermissionClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingPermissionClient.CheckAccessByBillingAccount`

```go
ctx := context.TODO()
id := billingpermission.NewBillingAccountID("billingAccountValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByBillingAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.CheckAccessByBillingProfile`

```go
ctx := context.TODO()
id := billingpermission.NewBillingProfileID("billingAccountValue", "billingProfileValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByBillingProfile(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.CheckAccessByCustomer`

```go
ctx := context.TODO()
id := billingpermission.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByCustomer(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.CheckAccessByDepartment`

```go
ctx := context.TODO()
id := billingpermission.NewDepartmentID("billingAccountValue", "departmentValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByDepartment(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.CheckAccessByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingpermission.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByEnrollmentAccount(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.CheckAccessByInvoiceSection`

```go
ctx := context.TODO()
id := billingpermission.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

payload := billingpermission.CheckAccessRequest{
	// ...
}


read, err := client.CheckAccessByInvoiceSection(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `BillingPermissionClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingpermission.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingpermission.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingpermission.NewBillingProfileCustomerID("billingAccountValue", "billingProfileValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByCustomerAtBillingAccount`

```go
ctx := context.TODO()
id := billingpermission.NewCustomerID("billingAccountValue", "customerValue")

// alternatively `client.ListByCustomerAtBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerAtBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByDepartment`

```go
ctx := context.TODO()
id := billingpermission.NewDepartmentID("billingAccountValue", "departmentValue")

// alternatively `client.ListByDepartment(ctx, id)` can be used to do batched pagination
items, err := client.ListByDepartmentComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByEnrollmentAccount`

```go
ctx := context.TODO()
id := billingpermission.NewEnrollmentAccountID("billingAccountValue", "enrollmentAccountValue")

// alternatively `client.ListByEnrollmentAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByEnrollmentAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionClient.ListByInvoiceSection`

```go
ctx := context.TODO()
id := billingpermission.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSection(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
