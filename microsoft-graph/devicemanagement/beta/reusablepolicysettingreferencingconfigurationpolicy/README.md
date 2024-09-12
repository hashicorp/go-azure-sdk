
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


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.AssignReusablePolicySettingReferencingConfigurationPolicies`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.AssignReusablePolicySettingReferencingConfigurationPoliciesRequest{
	// ...
}


// alternatively `client.AssignReusablePolicySettingReferencingConfigurationPolicies(ctx, id, payload, reusablepolicysettingreferencingconfigurationpolicy.DefaultAssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.AssignReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx, id, payload, reusablepolicysettingreferencingconfigurationpolicy.DefaultAssignReusablePolicySettingReferencingConfigurationPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

read, err := client.ClearReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.CreateReusablePolicySettingReferencingConfigurationPolicyCopy`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.CreateReusablePolicySettingReferencingConfigurationPolicyCopyRequest{
	// ...
}


read, err := client.CreateReusablePolicySettingReferencingConfigurationPolicyCopy(ctx, id, payload)
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


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.CreateReusablePolicySettingReferencingConfigurationPolicyRetrieveEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

read, err := client.CreateReusablePolicySettingReferencingConfigurationPolicyRetrieveEnrollmentTimeDeviceMembershipTarget(ctx, id)
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

read, err := client.DeleteReusablePolicySettingReferencingConfigurationPolicy(ctx, id, reusablepolicysettingreferencingconfigurationpolicy.DefaultDeleteReusablePolicySettingReferencingConfigurationPolicyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.GetReusablePolicySettingReferencingConfigurationPoliciesCount`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingID("deviceManagementReusablePolicySettingIdValue")

read, err := client.GetReusablePolicySettingReferencingConfigurationPoliciesCount(ctx, id, reusablepolicysettingreferencingconfigurationpolicy.DefaultGetReusablePolicySettingReferencingConfigurationPoliciesCountOperationOptions())
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

read, err := client.GetReusablePolicySettingReferencingConfigurationPolicy(ctx, id, reusablepolicysettingreferencingconfigurationpolicy.DefaultGetReusablePolicySettingReferencingConfigurationPolicyOperationOptions())
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

// alternatively `client.ListReusablePolicySettingReferencingConfigurationPolicies(ctx, id, reusablepolicysettingreferencingconfigurationpolicy.DefaultListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions())` can be used to do batched pagination
items, err := client.ListReusablePolicySettingReferencingConfigurationPoliciesComplete(ctx, id, reusablepolicysettingreferencingconfigurationpolicy.DefaultListReusablePolicySettingReferencingConfigurationPoliciesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ReusablePolicySettingReferencingConfigurationPolicyClient.SetReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget`

```go
ctx := context.TODO()
id := reusablepolicysettingreferencingconfigurationpolicy.NewDeviceManagementReusablePolicySettingIdReferencingConfigurationPolicyID("deviceManagementReusablePolicySettingIdValue", "deviceManagementConfigurationPolicyIdValue")

payload := reusablepolicysettingreferencingconfigurationpolicy.SetReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTargetRequest{
	// ...
}


read, err := client.SetReusablePolicySettingReferencingConfigurationPolicyEnrollmentTimeDeviceMembershipTarget(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
