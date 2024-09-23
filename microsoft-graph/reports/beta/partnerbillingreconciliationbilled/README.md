
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationbilled` Documentation

The `partnerbillingreconciliationbilled` SDK allows for interaction with Microsoft Graph `reports` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationbilled"
```


### Client Initialization

```go
client := partnerbillingreconciliationbilled.NewPartnerBillingReconciliationBilledClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerBillingReconciliationBilledClient.CreatePartnerBillingReconciliationBilledExport`

```go
ctx := context.TODO()

payload := partnerbillingreconciliationbilled.CreatePartnerBillingReconciliationBilledExportRequest{
	// ...
}


read, err := client.CreatePartnerBillingReconciliationBilledExport(ctx, payload, partnerbillingreconciliationbilled.DefaultCreatePartnerBillingReconciliationBilledExportOperationOptions())
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


read, err := client.UpdatePartnerBillingReconciliationBilled(ctx, payload, partnerbillingreconciliationbilled.DefaultUpdatePartnerBillingReconciliationBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
