
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreport` Documentation

The `grouppolicymigrationreport` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreport"
```


### Client Initialization

```go
client := grouppolicymigrationreport.NewGroupPolicyMigrationReportClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPolicyMigrationReportClient.CreateGroupPolicyMigrationReport`

```go
ctx := context.TODO()

payload := grouppolicymigrationreport.GroupPolicyMigrationReport{
	// ...
}


read, err := client.CreateGroupPolicyMigrationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.CreateGroupPolicyMigrationReportCreateMigrationReport`

```go
ctx := context.TODO()

payload := grouppolicymigrationreport.CreateGroupPolicyMigrationReportCreateMigrationReportRequest{
	// ...
}


read, err := client.CreateGroupPolicyMigrationReportCreateMigrationReport(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.DeleteGroupPolicyMigrationReport`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue")

read, err := client.DeleteGroupPolicyMigrationReport(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.GetGroupPolicyMigrationReport`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue")

read, err := client.GetGroupPolicyMigrationReport(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.GetGroupPolicyMigrationReportCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyMigrationReportCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.ListGroupPolicyMigrationReports`

```go
ctx := context.TODO()


// alternatively `client.ListGroupPolicyMigrationReports(ctx)` can be used to do batched pagination
items, err := client.ListGroupPolicyMigrationReportsComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyMigrationReportClient.UpdateDeviceManagementGroupPolicyMigrationReportScopeTag`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue")

payload := grouppolicymigrationreport.UpdateDeviceManagementGroupPolicyMigrationReportScopeTagRequest{
	// ...
}


read, err := client.UpdateDeviceManagementGroupPolicyMigrationReportScopeTag(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.UpdateGroupPolicyMigrationReport`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportIdValue")

payload := grouppolicymigrationreport.GroupPolicyMigrationReport{
	// ...
}


read, err := client.UpdateGroupPolicyMigrationReport(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
