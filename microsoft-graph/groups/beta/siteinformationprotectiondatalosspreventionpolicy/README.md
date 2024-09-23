
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectiondatalosspreventionpolicy` Documentation

The `siteinformationprotectiondatalosspreventionpolicy` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/siteinformationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := siteinformationprotectiondatalosspreventionpolicy.NewSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.CreateSiteInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateSiteInformationProtectionDataLossPreventionPolicy(ctx, id, payload, siteinformationprotectiondatalosspreventionpolicy.DefaultCreateSiteInformationProtectionDataLossPreventionPolicyOperationOptions())
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
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupId", "siteId", "dataLossPreventionPolicyId")

read, err := client.DeleteSiteInformationProtectionDataLossPreventionPolicy(ctx, id, siteinformationprotectiondatalosspreventionpolicy.DefaultDeleteSiteInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SiteInformationProtectionDataLossPreventionPolicyClient.EvaluateSiteInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupId", "siteId")

payload := siteinformationprotectiondatalosspreventionpolicy.EvaluateSiteInformationProtectionDataLossPreventionPoliciesRequest{
	// ...
}


read, err := client.EvaluateSiteInformationProtectionDataLossPreventionPolicies(ctx, id, payload, siteinformationprotectiondatalosspreventionpolicy.DefaultEvaluateSiteInformationProtectionDataLossPreventionPoliciesOperationOptions())
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
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupId", "siteId")

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
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupId", "siteId", "dataLossPreventionPolicyId")

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
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteID("groupId", "siteId")

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
id := siteinformationprotectiondatalosspreventionpolicy.NewGroupIdSiteIdInformationProtectionDataLossPreventionPolicyID("groupId", "siteId", "dataLossPreventionPolicyId")

payload := siteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateSiteInformationProtectionDataLossPreventionPolicy(ctx, id, payload, siteinformationprotectiondatalosspreventionpolicy.DefaultUpdateSiteInformationProtectionDataLossPreventionPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
