
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationunbilled` Documentation

The `partnerbillingreconciliationunbilled` SDK allows for interaction with Microsoft Graph `reports` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingreconciliationunbilled"
```


### Client Initialization

```go
client := partnerbillingreconciliationunbilled.NewPartnerBillingReconciliationUnbilledClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerBillingReconciliationUnbilledClient.CreatePartnerBillingReconciliationUnbilledExport`

```go
ctx := context.TODO()

payload := partnerbillingreconciliationunbilled.CreatePartnerBillingReconciliationUnbilledExportRequest{
	// ...
}


read, err := client.CreatePartnerBillingReconciliationUnbilledExport(ctx, payload, partnerbillingreconciliationunbilled.DefaultCreatePartnerBillingReconciliationUnbilledExportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationUnbilledClient.DeletePartnerBillingReconciliationUnbilled`

```go
ctx := context.TODO()


read, err := client.DeletePartnerBillingReconciliationUnbilled(ctx, partnerbillingreconciliationunbilled.DefaultDeletePartnerBillingReconciliationUnbilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationUnbilledClient.GetPartnerBillingReconciliationUnbilled`

```go
ctx := context.TODO()


read, err := client.GetPartnerBillingReconciliationUnbilled(ctx, partnerbillingreconciliationunbilled.DefaultGetPartnerBillingReconciliationUnbilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingReconciliationUnbilledClient.UpdatePartnerBillingReconciliationUnbilled`

```go
ctx := context.TODO()

payload := partnerbillingreconciliationunbilled.PartnersBillingUnbilledReconciliation{
	// ...
}


read, err := client.UpdatePartnerBillingReconciliationUnbilled(ctx, payload, partnerbillingreconciliationunbilled.DefaultUpdatePartnerBillingReconciliationUnbilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
