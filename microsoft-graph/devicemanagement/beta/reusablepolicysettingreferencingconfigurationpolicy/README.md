
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicy` Documentation

The `reusablepolicysettingreferencingconfigurationpolicy` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/reusablepolicysettingreferencingconfigurationpolicy"
```


### Client Initialization

```go
client := reusablepolicysettingreferencingconfigurationpolicy.NewReusablePolicySettingReferencingConfigurationPolicyClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPolicies`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesRequest{
	// ...
}


// alternatively `client.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPolicies(ctx, id, payload)` can be used to do batched pagination
items, err := client.AssignDeviceManagementReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx, id, payload)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.CreateReusablePolicySettingReferencingConfigurationPolicy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.CreateReusablePolicySettingReferencingConfigurationPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.CreateReusablePolicySettingReferencingConfigurationPolicyCreateCopy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.CreateReusablePolicySettingReferencingConfigurationPolicyCreateCopyRequest{
	// ...
}


read, err := client.CreateReusablePolicySettingReferencingConfigurationPolicyCreateCopy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.CreateReusablePolicySettingReferencingConfigurationPolicyReorder`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.CreateReusablePolicySettingReferencingConfigurationPolicyReorderRequest{
	// ...
}


read, err := client.CreateReusablePolicySettingReferencingConfigurationPolicyReorder(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.DeleteReusablePolicySettingReferencingConfigurationPolicy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

read, err := client.DeleteReusablePolicySettingReferencingConfigurationPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.GetReusablePolicySettingReferencingConfigurationPolicy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

read, err := client.GetReusablePolicySettingReferencingConfigurationPolicy(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.GetReusablePolicySettingReferencingConfigurationPolicyCount`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingIdValue")

read, err := client.GetReusablePolicySettingReferencingConfigurationPolicyCount(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.ListReusablePolicySettingReferencingConfigurationPolicies`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingIdValue")

// alternatively `client.ListReusablePolicySettingReferencingConfigurationPolicies(ctx, id)` can be used to do batched pagination
items, err := client.ListReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.UpdateReusablePolicySettingReferencingConfigurationPolicy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.DeviceManagementConfigurationPolicy{
	// ...
}


read, err := client.UpdateReusablePolicySettingReferencingConfigurationPolicy(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
