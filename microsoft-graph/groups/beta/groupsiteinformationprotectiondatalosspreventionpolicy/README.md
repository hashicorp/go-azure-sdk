
## `github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectiondatalosspreventionpolicy` Documentation

The `groupsiteinformationprotectiondatalosspreventionpolicy` SDK allows for interaction with the Azure Resource Manager Service `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/groups/beta/groupsiteinformationprotectiondatalosspreventionpolicy"
```


### Client Initialization

```go
client := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteInformationProtectionDataLossPreventionPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.CreateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicy`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.CreateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyEvaluate`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteID("groupIdValue", "siteIdValue")

payload := groupsiteinformationprotectiondatalosspreventionpolicy.CreateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyEvaluateRequest{
	// ...
}


read, err := client.CreateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyEvaluate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.DeleteGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.DeleteGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.GetGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.GetGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyCount`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteID("groupIdValue", "siteIdValue")

read, err := client.GetGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicies`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteID("groupIdValue", "siteIdValue")

// alternatively `client.ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListGroupByIdSiteByIdInformationProtectionDataLossPreventionPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupSiteInformationProtectionDataLossPreventionPolicyClient.UpdateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById`

```go
ctx := context.TODO()
id := groupsiteinformationprotectiondatalosspreventionpolicy.NewGroupSiteInformationProtectionDataLossPreventionPolicyID("groupIdValue", "siteIdValue", "dataLossPreventionPolicyIdValue")

payload := groupsiteinformationprotectiondatalosspreventionpolicy.DataLossPreventionPolicy{
	// ...
}


read, err := client.UpdateGroupByIdSiteByIdInformationProtectionDataLossPreventionPolicyById(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
