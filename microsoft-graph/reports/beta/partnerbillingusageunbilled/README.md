
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusageunbilled` Documentation

The `partnerbillingusageunbilled` SDK allows for interaction with the Azure Resource Manager Service `reports` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusageunbilled"
```


### Client Initialization

```go
client := partnerbillingusageunbilled.NewPartnerBillingUsageUnbilledClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerBillingUsageUnbilledClient.CreatePartnerBillingUsageUnbilledExport`

```go
ctx := context.TODO()

payload := partnerbillingusageunbilled.CreatePartnerBillingUsageUnbilledExportRequest{
	// ...
}


read, err := client.CreatePartnerBillingUsageUnbilledExport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageUnbilledClient.DeletePartnerBillingUsageUnbilled`

```go
ctx := context.TODO()


read, err := client.DeletePartnerBillingUsageUnbilled(ctx, partnerbillingusageunbilled.DefaultDeletePartnerBillingUsageUnbilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageUnbilledClient.GetPartnerBillingUsageUnbilled`

```go
ctx := context.TODO()


read, err := client.GetPartnerBillingUsageUnbilled(ctx, partnerbillingusageunbilled.DefaultGetPartnerBillingUsageUnbilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageUnbilledClient.UpdatePartnerBillingUsageUnbilled`

```go
ctx := context.TODO()

payload := partnerbillingusageunbilled.PartnersBillingUnbilledUsage{
	// ...
}


read, err := client.UpdatePartnerBillingUsageUnbilled(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
