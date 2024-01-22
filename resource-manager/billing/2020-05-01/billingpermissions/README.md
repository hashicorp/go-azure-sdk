
## `github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingpermissions` Documentation

The `billingpermissions` SDK allows for interaction with the Azure Resource Manager Service `billing` (API Version `2020-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/billing/2020-05-01/billingpermissions"
```


### Client Initialization

```go
client := billingpermissions.NewBillingPermissionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `BillingPermissionsClient.ListByBillingAccount`

```go
ctx := context.TODO()
id := billingpermissions.NewBillingAccountID("billingAccountValue")

// alternatively `client.ListByBillingAccount(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingAccountComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionsClient.ListByBillingProfile`

```go
ctx := context.TODO()
id := billingpermissions.NewBillingProfileID("billingAccountValue", "billingProfileValue")

// alternatively `client.ListByBillingProfile(ctx, id)` can be used to do batched pagination
items, err := client.ListByBillingProfileComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionsClient.ListByCustomer`

```go
ctx := context.TODO()
id := billingpermissions.NewCustomerID("billingAccountValue", "customerValue")

// alternatively `client.ListByCustomer(ctx, id)` can be used to do batched pagination
items, err := client.ListByCustomerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `BillingPermissionsClient.ListByInvoiceSections`

```go
ctx := context.TODO()
id := billingpermissions.NewInvoiceSectionID("billingAccountValue", "billingProfileValue", "invoiceSectionValue")

// alternatively `client.ListByInvoiceSections(ctx, id)` can be used to do batched pagination
items, err := client.ListByInvoiceSectionsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
