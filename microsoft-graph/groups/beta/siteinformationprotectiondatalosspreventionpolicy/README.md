
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectiondatalosspreventionpolicy` Documentation

The `siteinformationprotectiondatalosspreventionpolicy` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := siteinformationprotectiondatalosspreventionpolicy.NewSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.CreateSiteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateSiteInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.CreateSiteInformationProtectionDataLossPreventionPolicyEvaluate`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupIdValue", "siteIdValue")

payload := siteinformationprotectiondatalosspreventionpolicy.CreateSiteInformationProtectionDataLossPreventionPolicyEvaluateRequest{
	// ...
}


read, err := client.CreateSiteInformationProtectionDataLossPreventionPolicyEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.DeleteSiteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.DeleteSiteInformationProtectionDataLossPreventionPolicy(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultDeleteSiteInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.GetSiteInformationProtectionDataLossPreventionPoliciesCount`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupIdValue", "siteIdValue")

read, err := client.GetSiteInformationProtectionDataLossPreventionPoliciesCount(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultGetSiteInformationProtectionDataLossPreventionPoliciesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.GetSiteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.GetSiteInformationProtectionDataLossPreventionPolicy(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultGetSiteInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.ListSiteInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListSiteInformationProtectionDataLossPreventionPolicies(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultListSiteInformationProtectionDataLossPreventionPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListSiteInformationProtectionDataLossPreventionPoliciesComplete(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultListSiteInformationProtectionDataLossPreventionPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.UpdateSiteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

payload := siteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateSiteInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
