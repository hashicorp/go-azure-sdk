
## `github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listqnamakerendpointkeys` Documentation

The `listqnamakerendpointkeys` SDK allows for interaction with the Azure Resource Manager Service `botservice` (API Version `2022-09-15`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/botservice/2022-09-15/listqnamakerendpointkeys"
```


### Client Initialization

```go
client := listqnamakerendpointkeys.NewListQnAMakerEndpointKeysClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ListQnAMakerEndpointKeysClient.QnAMakerEndpointKeysGet`

```go
ctx := context.TODO()
id := listqnamakerendpointkeys.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := listqnamakerendpointkeys.QnAMakerEndpointKeysRequestBody{
	// ...
}


read, err := client.QnAMakerEndpointKeysGet(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
