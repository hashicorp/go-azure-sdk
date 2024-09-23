
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreport` Documentation

The `grouppolicymigrationreport` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/grouppolicymigrationreport"
```


### Client Initialization

```go
client := grouppolicymigrationreport.NewGroupPolicyMigrationReportClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupPolicyMigrationReportClient.CreateGroupPolicyMigrationReport`

```go
ctx := context.TODO()

payload := grouppolicymigrationreport.GroupPolicyMigrationReport{
	// ...
}


read, err := client.CreateGroupPolicyMigrationReport(ctx, payload, grouppolicymigrationreport.DefaultCreateGroupPolicyMigrationReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.CreateGroupPolicyMigrationReportsMigrationReport`

```go
ctx := context.TODO()

payload := grouppolicymigrationreport.CreateGroupPolicyMigrationReportsMigrationReportRequest{
	// ...
}


read, err := client.CreateGroupPolicyMigrationReportsMigrationReport(ctx, payload, grouppolicymigrationreport.DefaultCreateGroupPolicyMigrationReportsMigrationReportOperationOptions())
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
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportId")

read, err := client.DeleteGroupPolicyMigrationReport(ctx, id, grouppolicymigrationreport.DefaultDeleteGroupPolicyMigrationReportOperationOptions())
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
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportId")

read, err := client.GetGroupPolicyMigrationReport(ctx, id, grouppolicymigrationreport.DefaultGetGroupPolicyMigrationReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.GetGroupPolicyMigrationReportsCount`

```go
ctx := context.TODO()


read, err := client.GetGroupPolicyMigrationReportsCount(ctx, grouppolicymigrationreport.DefaultGetGroupPolicyMigrationReportsCountOperationOptions())
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


// alternatively `client.ListGroupPolicyMigrationReports(ctx, grouppolicymigrationreport.DefaultListGroupPolicyMigrationReportsOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupPolicyMigrationReportsComplete(ctx, grouppolicymigrationreport.DefaultListGroupPolicyMigrationReportsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupPolicyMigrationReportClient.UpdateGroupPolicyMigrationReport`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportId")

payload := grouppolicymigrationreport.GroupPolicyMigrationReport{
	// ...
}


read, err := client.UpdateGroupPolicyMigrationReport(ctx, id, payload, grouppolicymigrationreport.DefaultUpdateGroupPolicyMigrationReportOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupPolicyMigrationReportClient.UpdateGroupPolicyMigrationReportScopeTags`

```go
ctx := context.TODO()
id := grouppolicymigrationreport.NewDeviceManagementGroupPolicyMigrationReportID("groupPolicyMigrationReportId")

payload := grouppolicymigrationreport.UpdateGroupPolicyMigrationReportScopeTagsRequest{
	// ...
}


read, err := client.UpdateGroupPolicyMigrationReportScopeTags(ctx, id, payload, grouppolicymigrationreport.DefaultUpdateGroupPolicyMigrationReportScopeTagsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
