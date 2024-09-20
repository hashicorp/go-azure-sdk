
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusagebilled` Documentation

The `partnerbillingusagebilled` SDK allows for interaction with Microsoft Graph `reports` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/reports/beta/partnerbillingusagebilled"
```


### Client Initialization

```go
client := partnerbillingusagebilled.NewPartnerBillingUsageBilledClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PartnerBillingUsageBilledClient.CreatePartnerBillingUsageBilledExport`

```go
ctx := context.TODO()

payload := partnerbillingusagebilled.CreatePartnerBillingUsageBilledExportRequest{
	// ...
}


read, err := client.CreatePartnerBillingUsageBilledExport(ctx, payload, partnerbillingusagebilled.DefaultCreatePartnerBillingUsageBilledExportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageBilledClient.DeletePartnerBillingUsageBilled`

```go
ctx := context.TODO()


read, err := client.DeletePartnerBillingUsageBilled(ctx, partnerbillingusagebilled.DefaultDeletePartnerBillingUsageBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageBilledClient.GetPartnerBillingUsageBilled`

```go
ctx := context.TODO()


read, err := client.GetPartnerBillingUsageBilled(ctx, partnerbillingusagebilled.DefaultGetPartnerBillingUsageBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PartnerBillingUsageBilledClient.UpdatePartnerBillingUsageBilled`

```go
ctx := context.TODO()

payload := partnerbillingusagebilled.PartnersBillingBilledUsage{
	// ...
}


read, err := client.UpdatePartnerBillingUsageBilled(ctx, payload, partnerbillingusagebilled.DefaultUpdatePartnerBillingUsageBilledOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
