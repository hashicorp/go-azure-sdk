
## `github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/serveradvisors` Documentation

The `serveradvisors` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-11-01/serveradvisors"
```


### Client Initialization

```go
client := serveradvisors.NewServerAdvisorsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerAdvisorsClient.Get`

```go
ctx := context.TODO()
id := serveradvisors.NewAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "advisorValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerAdvisorsClient.ListByServer`

```go
ctx := context.TODO()
id := serveradvisors.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

read, err := client.ListByServer(ctx, id, serveradvisors.DefaultListByServerOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerAdvisorsClient.Update`

```go
ctx := context.TODO()
id := serveradvisors.NewAdvisorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "advisorValue")

payload := serveradvisors.Advisor{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
