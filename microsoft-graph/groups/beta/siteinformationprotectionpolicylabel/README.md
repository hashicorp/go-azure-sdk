
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel` Documentation

The `siteinformationprotectionpolicylabel` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectionpolicylabel"
```


### Client Initialization

```go
client := siteinformationprotectionpolicylabel.NewSiteInformationProtectionPolicyLabelClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabelEvaluateApplication`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.CreateSiteInformationProtectionPolicyLabelEvaluateApplicationRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabelEvaluateApplication(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabelEvaluateClassificationResult`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.CreateSiteInformationProtectionPolicyLabelEvaluateClassificationResultRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabelEvaluateClassificationResult(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabelEvaluateRemoval`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.CreateSiteInformationProtectionPolicyLabelEvaluateRemovalRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabelEvaluateRemoval(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.CreateSiteInformationProtectionPolicyLabelExtractLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectionpolicylabel.CreateSiteInformationProtectionPolicyLabelExtractLabelRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionPolicyLabelExtractLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.DeleteSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

read, err := client.DeleteSiteInformationProtectionPolicyLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.GetSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

read, err := client.GetSiteInformationProtectionPolicyLabel(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.GetSiteInformationProtectionPolicyLabelCount`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteInformationProtectionPolicyLabelCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.ListSiteInformationProtectionPolicyLabels`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListSiteInformationProtectionPolicyLabels(ctx, id)` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionPolicyLabelsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionPolicyLabelClient.UpdateSiteInformationProtectionPolicyLabel`

```go
ctx := context.TODO()
id := siteinformationprotectionpolicylabel.NewGroupIdSiteIdInformationProtectionPolicyLabelID("groupIdValue", "siteIdValue", "informationProtectionLabelIdValue")

payload := siteinformationprotectionpolicylabel.InformationProtectionLabel{
	// ...
}


read, err := client.UpdateSiteInformationProtectionPolicyLabel(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
