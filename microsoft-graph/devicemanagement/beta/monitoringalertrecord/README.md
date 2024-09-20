
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/monitoringalertrecord` Documentation

The `monitoringalertrecord` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/monitoringalertrecord"
```


### Client Initialization

```go
client := monitoringalertrecord.NewMonitoringAlertRecordClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MonitoringAlertRecordClient.CreateMonitoringAlertRecord`

```go
ctx := context.TODO()

payload := monitoringalertrecord.DeviceManagementAlertRecord{
	// ...
}


read, err := client.CreateMonitoringAlertRecord(ctx, payload, monitoringalertrecord.DefaultCreateMonitoringAlertRecordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSent`

```go
ctx := context.TODO()

payload := monitoringalertrecord.CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSentRequest{
	// ...
}


read, err := client.CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSent(ctx, payload, monitoringalertrecord.DefaultCreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent`

```go
ctx := context.TODO()
id := monitoringalertrecord.NewDeviceManagementMonitoringAlertRecordID("alertRecordId")

read, err := client.CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent(ctx, id, monitoringalertrecord.DefaultCreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.DeleteMonitoringAlertRecord`

```go
ctx := context.TODO()
id := monitoringalertrecord.NewDeviceManagementMonitoringAlertRecordID("alertRecordId")

read, err := client.DeleteMonitoringAlertRecord(ctx, id, monitoringalertrecord.DefaultDeleteMonitoringAlertRecordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.GetMonitoringAlertRecord`

```go
ctx := context.TODO()
id := monitoringalertrecord.NewDeviceManagementMonitoringAlertRecordID("alertRecordId")

read, err := client.GetMonitoringAlertRecord(ctx, id, monitoringalertrecord.DefaultGetMonitoringAlertRecordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.GetMonitoringAlertRecordsCount`

```go
ctx := context.TODO()


read, err := client.GetMonitoringAlertRecordsCount(ctx, monitoringalertrecord.DefaultGetMonitoringAlertRecordsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MonitoringAlertRecordClient.ListMonitoringAlertRecords`

```go
ctx := context.TODO()


// alternatively `client.ListMonitoringAlertRecords(ctx, monitoringalertrecord.DefaultListMonitoringAlertRecordsOperationOptions())` can be used to do batched pagination
items, err := client.ListMonitoringAlertRecordsComplete(ctx, monitoringalertrecord.DefaultListMonitoringAlertRecordsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MonitoringAlertRecordClient.UpdateMonitoringAlertRecord`

```go
ctx := context.TODO()
id := monitoringalertrecord.NewDeviceManagementMonitoringAlertRecordID("alertRecordId")

payload := monitoringalertrecord.DeviceManagementAlertRecord{
	// ...
}


read, err := client.UpdateMonitoringAlertRecord(ctx, id, payload, monitoringalertrecord.DefaultUpdateMonitoringAlertRecordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
