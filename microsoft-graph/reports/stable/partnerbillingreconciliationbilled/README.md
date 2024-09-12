
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingreconciliationbilled` Documentation

The `partnerbillingreconciliationbilled` SDK allows for interaction with the Azure Resource Manager Service `reports` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/stable/partnerbillingreconciliationbilled"
```


### Client Initialization

```go
client := partnerbillingreconciliationbilled.NewPartnerBillingReconciliationBilledClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerBillingReconciliationBilledClient.CreatePartnerBillingReconciliationBilledExport`

```go
ctx := context.TODO()

payload := partnerbillingreconciliationbilled.CreatePartnerBillingReconciliationBilledExportRequest{
	// ...
}


read, err := client.CreatePartnerBillingReconciliationBilledExport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationBilledClient.DeletePartnerBillingReconciliationBilled`

```go
ctx := context.TODO()


read, err := client.DeletePartnerBillingReconciliationBilled(ctx, partnerbillingreconciliationbilled.DefaultDeletePartnerBillingReconciliationBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationBilledClient.GetPartnerBillingReconciliationBilled`

```go
ctx := context.TODO()


read, err := client.GetPartnerBillingReconciliationBilled(ctx, partnerbillingreconciliationbilled.DefaultGetPartnerBillingReconciliationBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationBilledClient.UpdatePartnerBillingReconciliationBilled`

```go
ctx := context.TODO()

payload := partnerbillingreconciliationbilled.PartnersBillingBilledReconciliation{
	// ...
}


read, err := client.UpdatePartnerBillingReconciliationBilled(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
