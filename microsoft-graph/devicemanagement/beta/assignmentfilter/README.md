
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter` Documentation

The `assignmentfilter` SDK allows for interaction with the Azure Resource Manager Service `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/assignmentfilter"
```


### Client Initialization

```go
client := assignmentfilter.NewAssignmentFilterClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AssignmentFilterClient.CreateAssignmentFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.CreateAssignmentFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.CreateAssignmentFilterEnable`

```go
ctx := context.TODO()

payload := assignmentfilter.CreateAssignmentFilterEnableRequest{
	// ...
}


read, err := client.CreateAssignmentFilterEnable(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.DeleteAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

read, err := client.DeleteAssignmentFilter(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.GetAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

read, err := client.GetAssignmentFilter(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.GetAssignmentFilterCount`

```go
ctx := context.TODO()


read, err := client.GetAssignmentFilterCount(ctx)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.ListAssignmentFilters`

```go
ctx := context.TODO()


// alternatively `client.ListAssignmentFilters(ctx)` can be used to do batched pagination
items, err := client.ListAssignmentFiltersComplete(ctx)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AssignmentFilterClient.UpdateAssignmentFilter`

```go
ctx := context.TODO()
id := assignmentfilter.NewDeviceManagementAssignmentFilterID("deviceAndAppManagementAssignmentFilterIdValue")

payload := assignmentfilter.DeviceAndAppManagementAssignmentFilter{
	// ...
}


read, err := client.UpdateAssignmentFilter(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AssignmentFilterClient.ValidateDeviceManagementAssignmentFiltersFilter`

```go
ctx := context.TODO()

payload := assignmentfilter.ValidateDeviceManagementAssignmentFiltersFilterRequest{
	// ...
}


read, err := client.ValidateDeviceManagementAssignmentFiltersFilter(ctx, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
