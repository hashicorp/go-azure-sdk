
## `github.com/hashicorp/go-azure-sdk/data-plane/iotcentral/2022-10-31-preview/iotcentrals` Documentation

The `iotcentrals` SDK allows for interaction with <unknown source data type> `iotcentral` (API Version `2022-10-31-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/data-plane/iotcentral/2022-10-31-preview/iotcentrals"
```


### Client Initialization

```go
client := iotcentrals.NewIotcentralsClientWithBaseURI("")
client.Client.Authorizer = authorizer
```


### Example Usage: `IotcentralsClient.ApiTokensCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewApiTokenID("tokenId")

payload := iotcentrals.ApiToken{
	// ...
}


read, err := client.ApiTokensCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ApiTokensGet`

```go
ctx := context.TODO()
id := iotcentrals.NewApiTokenID("tokenId")

read, err := client.ApiTokensGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ApiTokensList`

```go
ctx := context.TODO()


// alternatively `client.ApiTokensList(ctx)` can be used to do batched pagination
items, err := client.ApiTokensListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ApiTokensRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewApiTokenID("tokenId")

read, err := client.ApiTokensRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DashboardsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDashboardID("dashboardId")

payload := iotcentrals.Dashboard{
	// ...
}


read, err := client.DashboardsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DashboardsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDashboardID("dashboardId")

read, err := client.DashboardsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DashboardsList`

```go
ctx := context.TODO()


// alternatively `client.DashboardsList(ctx, iotcentrals.DefaultDashboardsListOperationOptions())` can be used to do batched pagination
items, err := client.DashboardsListComplete(ctx, iotcentrals.DefaultDashboardsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DashboardsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDashboardID("dashboardId")

read, err := client.DashboardsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DashboardsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDashboardID("dashboardId")
var payload interface{}

read, err := client.DashboardsUpdate(ctx, id, payload, iotcentrals.DefaultDashboardsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeploymentManifestsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeploymentManifestID("deploymentManifestId")

payload := iotcentrals.DeploymentManifest{
	// ...
}


read, err := client.DeploymentManifestsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeploymentManifestsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDeploymentManifestID("deploymentManifestId")

read, err := client.DeploymentManifestsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeploymentManifestsList`

```go
ctx := context.TODO()


// alternatively `client.DeploymentManifestsList(ctx, iotcentrals.DefaultDeploymentManifestsListOperationOptions())` can be used to do batched pagination
items, err := client.DeploymentManifestsListComplete(ctx, iotcentrals.DefaultDeploymentManifestsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DeploymentManifestsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDeploymentManifestID("deploymentManifestId")

read, err := client.DeploymentManifestsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeploymentManifestsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeploymentManifestID("deploymentManifestId")
var payload interface{}

read, err := client.DeploymentManifestsUpdate(ctx, id, payload, iotcentrals.DefaultDeploymentManifestsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DestinationsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDestinationIdID("https://endpoint-url.example.com")

payload := iotcentrals.Destination{
	// ...
}


read, err := client.DestinationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DestinationsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDestinationIdID("https://endpoint-url.example.com")

read, err := client.DestinationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DestinationsList`

```go
ctx := context.TODO()


// alternatively `client.DestinationsList(ctx)` can be used to do batched pagination
items, err := client.DestinationsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DestinationsListExports`

```go
ctx := context.TODO()
id := iotcentrals.NewDestinationIdID("https://endpoint-url.example.com")

// alternatively `client.DestinationsListExports(ctx, id)` can be used to do batched pagination
items, err := client.DestinationsListExportsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DestinationsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDestinationIdID("https://endpoint-url.example.com")

read, err := client.DestinationsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DestinationsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDestinationIdID("https://endpoint-url.example.com")
var payload interface{}

read, err := client.DestinationsUpdate(ctx, id, payload, iotcentrals.DefaultDestinationsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceGroupID("deviceGroupId")

payload := iotcentrals.DeviceGroup{
	// ...
}


read, err := client.DeviceGroupsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceGroupID("deviceGroupId")

read, err := client.DeviceGroupsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsGetDevices`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceGroupID("deviceGroupId")

// alternatively `client.DeviceGroupsGetDevices(ctx, id, iotcentrals.DefaultDeviceGroupsGetDevicesOperationOptions())` can be used to do batched pagination
items, err := client.DeviceGroupsGetDevicesComplete(ctx, id, iotcentrals.DefaultDeviceGroupsGetDevicesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsList`

```go
ctx := context.TODO()


// alternatively `client.DeviceGroupsList(ctx, iotcentrals.DefaultDeviceGroupsListOperationOptions())` can be used to do batched pagination
items, err := client.DeviceGroupsListComplete(ctx, iotcentrals.DefaultDeviceGroupsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceGroupID("deviceGroupId")

read, err := client.DeviceGroupsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceGroupsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceGroupID("deviceGroupId")
var payload interface{}

read, err := client.DeviceGroupsUpdate(ctx, id, payload, iotcentrals.DefaultDeviceGroupsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceTemplatesCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceTemplateID("deviceTemplateId")

payload := iotcentrals.DeviceTemplate{
	// ...
}


read, err := client.DeviceTemplatesCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceTemplatesGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceTemplateID("deviceTemplateId")

read, err := client.DeviceTemplatesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceTemplatesList`

```go
ctx := context.TODO()


// alternatively `client.DeviceTemplatesList(ctx, iotcentrals.DefaultDeviceTemplatesListOperationOptions())` can be used to do batched pagination
items, err := client.DeviceTemplatesListComplete(ctx, iotcentrals.DefaultDeviceTemplatesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DeviceTemplatesRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceTemplateID("deviceTemplateId")

read, err := client.DeviceTemplatesRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DeviceTemplatesUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceTemplateID("deviceTemplateId")
var payload interface{}

read, err := client.DeviceTemplatesUpdate(ctx, id, payload, iotcentrals.DefaultDeviceTemplatesUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesApplyManifest`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

payload := iotcentrals.DeploymentManifest{
	// ...
}


read, err := client.DevicesApplyManifest(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

payload := iotcentrals.Device{
	// ...
}


read, err := client.DevicesCreate(ctx, id, payload, iotcentrals.DefaultDevicesCreateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesCreateAttestation`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

payload := iotcentrals.Attestation{
	// ...
}


read, err := client.DevicesCreateAttestation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesCreateRelationship`

```go
ctx := context.TODO()
id := iotcentrals.NewRelationshipID("deviceId", "relationshipId")

payload := iotcentrals.DeviceRelationship{
	// ...
}


read, err := client.DevicesCreateRelationship(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGet`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesGet(ctx, id, iotcentrals.DefaultDevicesGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetAttestation`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesGetAttestation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetCommandHistory`

```go
ctx := context.TODO()
id := iotcentrals.NewCommandID("deviceId", "commandName")

// alternatively `client.DevicesGetCommandHistory(ctx, id)` can be used to do batched pagination
items, err := client.DevicesGetCommandHistoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesGetComponentCommandHistory`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentCommandID("https://endpoint-url.example.com", "deviceId", "componentName")

// alternatively `client.DevicesGetComponentCommandHistory(ctx, id)` can be used to do batched pagination
items, err := client.DevicesGetComponentCommandHistoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesGetComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentID("deviceId", "componentName")

read, err := client.DevicesGetComponentProperties(ctx, id, iotcentrals.DefaultDevicesGetComponentPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetComponentTelemetryValue`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentTelemetryID("https://endpoint-url.example.com", "deviceId", "componentName")

read, err := client.DevicesGetComponentTelemetryValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetCredentials`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesGetCredentials(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleCommandHistory`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleCommandID("https://endpoint-url.example.com", "deviceId", "moduleName")

// alternatively `client.DevicesGetModuleCommandHistory(ctx, id)` can be used to do batched pagination
items, err := client.DevicesGetModuleCommandHistoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleComponentCommandHistory`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentCommandID("deviceId", "moduleName", "componentName", "commandName")

// alternatively `client.DevicesGetModuleComponentCommandHistory(ctx, id)` can be used to do batched pagination
items, err := client.DevicesGetModuleComponentCommandHistoryComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentID("https://endpoint-url.example.com", "deviceId", "moduleName")

read, err := client.DevicesGetModuleComponentProperties(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleComponentTelemetryValue`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentTelemetryID("deviceId", "moduleName", "componentName", "telemetryName")

read, err := client.DevicesGetModuleComponentTelemetryValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleID("deviceId", "moduleName")

read, err := client.DevicesGetModuleProperties(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetModuleTelemetryValue`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleTelemetryID("https://endpoint-url.example.com", "deviceId", "moduleName")

read, err := client.DevicesGetModuleTelemetryValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesGetProperties(ctx, id, iotcentrals.DefaultDevicesGetPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetRelationship`

```go
ctx := context.TODO()
id := iotcentrals.NewRelationshipID("deviceId", "relationshipId")

read, err := client.DevicesGetRelationship(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesGetTelemetryValue`

```go
ctx := context.TODO()
id := iotcentrals.NewTelemetryID("deviceId", "telemetryName")

read, err := client.DevicesGetTelemetryValue(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesList`

```go
ctx := context.TODO()


// alternatively `client.DevicesList(ctx, iotcentrals.DefaultDevicesListOperationOptions())` can be used to do batched pagination
items, err := client.DevicesListComplete(ctx, iotcentrals.DefaultDevicesListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesListComponents`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

// alternatively `client.DevicesListComponents(ctx, id)` can be used to do batched pagination
items, err := client.DevicesListComponentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesListModuleComponents`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleID("deviceId", "moduleName")

// alternatively `client.DevicesListModuleComponents(ctx, id)` can be used to do batched pagination
items, err := client.DevicesListModuleComponentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesListModules`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

// alternatively `client.DevicesListModules(ctx, id)` can be used to do batched pagination
items, err := client.DevicesListModulesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesListRelationships`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

// alternatively `client.DevicesListRelationships(ctx, id, iotcentrals.DefaultDevicesListRelationshipsOperationOptions())` can be used to do batched pagination
items, err := client.DevicesListRelationshipsComplete(ctx, id, iotcentrals.DefaultDevicesListRelationshipsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.DevicesRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesRemove(ctx, id, iotcentrals.DefaultDevicesRemoveOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRemoveAttestation`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")

read, err := client.DevicesRemoveAttestation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRemoveRelationship`

```go
ctx := context.TODO()
id := iotcentrals.NewRelationshipID("deviceId", "relationshipId")

read, err := client.DevicesRemoveRelationship(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesReplaceComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentID("deviceId", "componentName")
var payload interface{}

read, err := client.DevicesReplaceComponentProperties(ctx, id, payload, iotcentrals.DefaultDevicesReplaceComponentPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesReplaceModuleComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentID("https://endpoint-url.example.com", "deviceId", "moduleName")
var payload interface{}

read, err := client.DevicesReplaceModuleComponentProperties(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesReplaceModuleProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleID("deviceId", "moduleName")
var payload interface{}

read, err := client.DevicesReplaceModuleProperties(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesReplaceProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")
var payload interface{}

read, err := client.DevicesReplaceProperties(ctx, id, payload, iotcentrals.DefaultDevicesReplacePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRunCommand`

```go
ctx := context.TODO()
id := iotcentrals.NewCommandID("deviceId", "commandName")

payload := iotcentrals.DeviceCommand{
	// ...
}


read, err := client.DevicesRunCommand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRunComponentCommand`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentCommandID("https://endpoint-url.example.com", "deviceId", "componentName")

payload := iotcentrals.DeviceCommand{
	// ...
}


read, err := client.DevicesRunComponentCommand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRunModuleCommand`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleCommandID("https://endpoint-url.example.com", "deviceId", "moduleName")

payload := iotcentrals.DeviceCommand{
	// ...
}


read, err := client.DevicesRunModuleCommand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesRunModuleComponentCommand`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentCommandID("deviceId", "moduleName", "componentName", "commandName")

payload := iotcentrals.DeviceCommand{
	// ...
}


read, err := client.DevicesRunModuleComponentCommand(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")
var payload interface{}

read, err := client.DevicesUpdate(ctx, id, payload, iotcentrals.DefaultDevicesUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateAttestation`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")
var payload interface{}

read, err := client.DevicesUpdateAttestation(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewComponentID("deviceId", "componentName")
var payload interface{}

read, err := client.DevicesUpdateComponentProperties(ctx, id, payload, iotcentrals.DefaultDevicesUpdateComponentPropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateModuleComponentProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleComponentID("https://endpoint-url.example.com", "deviceId", "moduleName")
var payload interface{}

read, err := client.DevicesUpdateModuleComponentProperties(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateModuleProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewModuleID("deviceId", "moduleName")
var payload interface{}

read, err := client.DevicesUpdateModuleProperties(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateProperties`

```go
ctx := context.TODO()
id := iotcentrals.NewDeviceID("deviceId")
var payload interface{}

read, err := client.DevicesUpdateProperties(ctx, id, payload, iotcentrals.DefaultDevicesUpdatePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.DevicesUpdateRelationship`

```go
ctx := context.TODO()
id := iotcentrals.NewRelationshipID("deviceId", "relationshipId")
var payload interface{}

read, err := client.DevicesUpdateRelationship(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewEnrollmentGroupID("enrollmentGroupId")

payload := iotcentrals.EnrollmentGroup{
	// ...
}


read, err := client.EnrollmentGroupsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsCreateX509`

```go
ctx := context.TODO()
id := iotcentrals.NewEntryID("enrollmentGroupId", "primary")

payload := iotcentrals.SigningX509Certificate{
	// ...
}


read, err := client.EnrollmentGroupsCreateX509(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsGenerateVerificationCodeX509`

```go
ctx := context.TODO()
id := iotcentrals.NewEntryID("enrollmentGroupId", "primary")

read, err := client.EnrollmentGroupsGenerateVerificationCodeX509(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewEnrollmentGroupID("enrollmentGroupId")

read, err := client.EnrollmentGroupsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsGetX509`

```go
ctx := context.TODO()
id := iotcentrals.NewEntryID("enrollmentGroupId", "primary")

read, err := client.EnrollmentGroupsGetX509(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsList`

```go
ctx := context.TODO()


// alternatively `client.EnrollmentGroupsList(ctx)` can be used to do batched pagination
items, err := client.EnrollmentGroupsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewEnrollmentGroupID("enrollmentGroupId")

read, err := client.EnrollmentGroupsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsRemoveX509`

```go
ctx := context.TODO()
id := iotcentrals.NewEntryID("enrollmentGroupId", "primary")

read, err := client.EnrollmentGroupsRemoveX509(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewEnrollmentGroupID("enrollmentGroupId")
var payload interface{}

read, err := client.EnrollmentGroupsUpdate(ctx, id, payload, iotcentrals.DefaultEnrollmentGroupsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.EnrollmentGroupsVerifyX509`

```go
ctx := context.TODO()
id := iotcentrals.NewEntryID("enrollmentGroupId", "primary")

payload := iotcentrals.X509VerificationCertificate{
	// ...
}


read, err := client.EnrollmentGroupsVerifyX509(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ExportsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewExportIdID("https://endpoint-url.example.com")

payload := iotcentrals.Export{
	// ...
}


read, err := client.ExportsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ExportsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewExportIdID("https://endpoint-url.example.com")

read, err := client.ExportsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ExportsList`

```go
ctx := context.TODO()


// alternatively `client.ExportsList(ctx)` can be used to do batched pagination
items, err := client.ExportsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ExportsListDestinations`

```go
ctx := context.TODO()
id := iotcentrals.NewExportIdID("https://endpoint-url.example.com")

// alternatively `client.ExportsListDestinations(ctx, id)` can be used to do batched pagination
items, err := client.ExportsListDestinationsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ExportsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewExportIdID("https://endpoint-url.example.com")

read, err := client.ExportsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ExportsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewExportIdID("https://endpoint-url.example.com")
var payload interface{}

read, err := client.ExportsUpdate(ctx, id, payload, iotcentrals.DefaultExportsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.FileUploadsCreate`

```go
ctx := context.TODO()

payload := iotcentrals.FileUpload{
	// ...
}


read, err := client.FileUploadsCreate(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.FileUploadsGet`

```go
ctx := context.TODO()


read, err := client.FileUploadsGet(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.FileUploadsRemove`

```go
ctx := context.TODO()


read, err := client.FileUploadsRemove(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.FileUploadsUpdate`

```go
ctx := context.TODO()
var payload interface{}

read, err := client.FileUploadsUpdate(ctx, payload, iotcentrals.DefaultFileUploadsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.JobsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewJobID("jobId")

payload := iotcentrals.Job{
	// ...
}


read, err := client.JobsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.JobsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewJobID("jobId")

read, err := client.JobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.JobsGetDevices`

```go
ctx := context.TODO()
id := iotcentrals.NewJobID("jobId")

// alternatively `client.JobsGetDevices(ctx, id)` can be used to do batched pagination
items, err := client.JobsGetDevicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.JobsList`

```go
ctx := context.TODO()


// alternatively `client.JobsList(ctx, iotcentrals.DefaultJobsListOperationOptions())` can be used to do batched pagination
items, err := client.JobsListComplete(ctx, iotcentrals.DefaultJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.JobsRerun`

```go
ctx := context.TODO()
id := iotcentrals.NewRerunID("jobId", "rerunId")

read, err := client.JobsRerun(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.JobsResume`

```go
ctx := context.TODO()
id := iotcentrals.NewJobID("jobId")

read, err := client.JobsResume(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.JobsStop`

```go
ctx := context.TODO()
id := iotcentrals.NewJobID("jobId")

read, err := client.JobsStop(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.OrganizationsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewOrganizationID("organizationId")

payload := iotcentrals.Organization{
	// ...
}


read, err := client.OrganizationsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.OrganizationsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewOrganizationID("organizationId")

read, err := client.OrganizationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.OrganizationsList`

```go
ctx := context.TODO()


// alternatively `client.OrganizationsList(ctx)` can be used to do batched pagination
items, err := client.OrganizationsListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.OrganizationsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewOrganizationID("organizationId")

read, err := client.OrganizationsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.OrganizationsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewOrganizationID("organizationId")
var payload interface{}

read, err := client.OrganizationsUpdate(ctx, id, payload, iotcentrals.DefaultOrganizationsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.QueryRun`

```go
ctx := context.TODO()

payload := iotcentrals.QueryRequest{
	// ...
}


read, err := client.QueryRun(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.RolesGet`

```go
ctx := context.TODO()
id := iotcentrals.NewRoleID("roleId")

read, err := client.RolesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.RolesList`

```go
ctx := context.TODO()


// alternatively `client.RolesList(ctx)` can be used to do batched pagination
items, err := client.RolesListComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewScheduledJobID("scheduledJobId")

payload := iotcentrals.ScheduledJob{
	// ...
}


read, err := client.ScheduledJobsCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsGet`

```go
ctx := context.TODO()
id := iotcentrals.NewScheduledJobID("scheduledJobId")

read, err := client.ScheduledJobsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsList`

```go
ctx := context.TODO()


// alternatively `client.ScheduledJobsList(ctx, iotcentrals.DefaultScheduledJobsListOperationOptions())` can be used to do batched pagination
items, err := client.ScheduledJobsListComplete(ctx, iotcentrals.DefaultScheduledJobsListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsListJobs`

```go
ctx := context.TODO()
id := iotcentrals.NewScheduledJobID("scheduledJobId")

// alternatively `client.ScheduledJobsListJobs(ctx, id, iotcentrals.DefaultScheduledJobsListJobsOperationOptions())` can be used to do batched pagination
items, err := client.ScheduledJobsListJobsComplete(ctx, id, iotcentrals.DefaultScheduledJobsListJobsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewScheduledJobID("scheduledJobId")

read, err := client.ScheduledJobsRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.ScheduledJobsUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewScheduledJobID("scheduledJobId")
var payload interface{}

read, err := client.ScheduledJobsUpdate(ctx, id, payload, iotcentrals.DefaultScheduledJobsUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.UsersCreate`

```go
ctx := context.TODO()
id := iotcentrals.NewUserID("userId")

payload := iotcentrals.User{
	// ...
}


read, err := client.UsersCreate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.UsersGet`

```go
ctx := context.TODO()
id := iotcentrals.NewUserID("userId")

read, err := client.UsersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.UsersList`

```go
ctx := context.TODO()


// alternatively `client.UsersList(ctx, iotcentrals.DefaultUsersListOperationOptions())` can be used to do batched pagination
items, err := client.UsersListComplete(ctx, iotcentrals.DefaultUsersListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `IotcentralsClient.UsersRemove`

```go
ctx := context.TODO()
id := iotcentrals.NewUserID("userId")

read, err := client.UsersRemove(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `IotcentralsClient.UsersUpdate`

```go
ctx := context.TODO()
id := iotcentrals.NewUserID("userId")
var payload interface{}

read, err := client.UsersUpdate(ctx, id, payload, iotcentrals.DefaultUsersUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
