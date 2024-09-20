
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelserverlogcollectionresponse` Documentation

The `microsofttunnelserverlogcollectionresponse` SDK allows for interaction with Microsoft Graph `devicemanagement` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/devicemanagement/beta/microsofttunnelserverlogcollectionresponse"
```


### Client Initialization

```go
client := microsofttunnelserverlogcollectionresponse.NewMicrosoftTunnelServerLogCollectionResponseClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.CreateMicrosoftTunnelServerLogCollectionResponse`

```go
ctx := context.TODO()

payload := microsofttunnelserverlogcollectionresponse.MicrosoftTunnelServerLogCollectionResponse{
	// ...
}


read, err := client.CreateMicrosoftTunnelServerLogCollectionResponse(ctx, payload, microsofttunnelserverlogcollectionresponse.DefaultCreateMicrosoftTunnelServerLogCollectionResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.CreateMicrosoftTunnelServerLogCollectionResponseDownloadUrl`

```go
ctx := context.TODO()
id := microsofttunnelserverlogcollectionresponse.NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

read, err := client.CreateMicrosoftTunnelServerLogCollectionResponseDownloadUrl(ctx, id, microsofttunnelserverlogcollectionresponse.DefaultCreateMicrosoftTunnelServerLogCollectionResponseDownloadUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.CreateMicrosoftTunnelServerLogCollectionResponseGenerateDownloadUrl`

```go
ctx := context.TODO()
id := microsofttunnelserverlogcollectionresponse.NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

read, err := client.CreateMicrosoftTunnelServerLogCollectionResponseGenerateDownloadUrl(ctx, id, microsofttunnelserverlogcollectionresponse.DefaultCreateMicrosoftTunnelServerLogCollectionResponseGenerateDownloadUrlOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.DeleteMicrosoftTunnelServerLogCollectionResponse`

```go
ctx := context.TODO()
id := microsofttunnelserverlogcollectionresponse.NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

read, err := client.DeleteMicrosoftTunnelServerLogCollectionResponse(ctx, id, microsofttunnelserverlogcollectionresponse.DefaultDeleteMicrosoftTunnelServerLogCollectionResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.GetMicrosoftTunnelServerLogCollectionResponse`

```go
ctx := context.TODO()
id := microsofttunnelserverlogcollectionresponse.NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

read, err := client.GetMicrosoftTunnelServerLogCollectionResponse(ctx, id, microsofttunnelserverlogcollectionresponse.DefaultGetMicrosoftTunnelServerLogCollectionResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.GetMicrosoftTunnelServerLogCollectionResponsesCount`

```go
ctx := context.TODO()


read, err := client.GetMicrosoftTunnelServerLogCollectionResponsesCount(ctx, microsofttunnelserverlogcollectionresponse.DefaultGetMicrosoftTunnelServerLogCollectionResponsesCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.ListMicrosoftTunnelServerLogCollectionResponses`

```go
ctx := context.TODO()


// alternatively `client.ListMicrosoftTunnelServerLogCollectionResponses(ctx, microsofttunnelserverlogcollectionresponse.DefaultListMicrosoftTunnelServerLogCollectionResponsesOperationOptions())` can be used to do batched pagination
items, err := client.ListMicrosoftTunnelServerLogCollectionResponsesComplete(ctx, microsofttunnelserverlogcollectionresponse.DefaultListMicrosoftTunnelServerLogCollectionResponsesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MicrosoftTunnelServerLogCollectionResponseClient.UpdateMicrosoftTunnelServerLogCollectionResponse`

```go
ctx := context.TODO()
id := microsofttunnelserverlogcollectionresponse.NewDeviceManagementMicrosoftTunnelServerLogCollectionResponseID("microsoftTunnelServerLogCollectionResponseId")

payload := microsofttunnelserverlogcollectionresponse.MicrosoftTunnelServerLogCollectionResponse{
	// ...
}


read, err := client.UpdateMicrosoftTunnelServerLogCollectionResponse(ctx, id, payload, microsofttunnelserverlogcollectionresponse.DefaultUpdateMicrosoftTunnelServerLogCollectionResponseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
