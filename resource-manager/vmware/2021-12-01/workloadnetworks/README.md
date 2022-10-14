
## `github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/workloadnetworks` Documentation

The `workloadnetworks` SDK allows for interaction with the Azure Resource Manager Service `vmware` (API Version `2021-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/vmware/2021-12-01/workloadnetworks"
```


### Client Initialization

```go
client := workloadnetworks.NewWorkloadNetworksClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkloadNetworksClient.CreateDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

payload := workloadnetworks.WorkloadNetworkDhcp{
	// ...
}


if err := client.CreateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

payload := workloadnetworks.WorkloadNetworkDnsService{
	// ...
}


if err := client.CreateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

payload := workloadnetworks.WorkloadNetworkDnsZone{
	// ...
}


if err := client.CreateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreatePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

payload := workloadnetworks.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.CreatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreatePublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

payload := workloadnetworks.WorkloadNetworkPublicIP{
	// ...
}


if err := client.CreatePublicIPThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

payload := workloadnetworks.WorkloadNetworkSegment{
	// ...
}


if err := client.CreateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.CreateVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVmGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

payload := workloadnetworks.WorkloadNetworkVMGroup{
	// ...
}


if err := client.CreateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

if err := client.DeleteDhcpThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

if err := client.DeleteDnsServiceThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

if err := client.DeleteDnsZoneThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeletePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

if err := client.DeletePortMirroringThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeletePublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

if err := client.DeletePublicIPThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteSegment`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

if err := client.DeleteSegmentThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.DeleteVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVmGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

if err := client.DeleteVMGroupThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.GetDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

read, err := client.GetDhcp(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

read, err := client.GetDnsService(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

read, err := client.GetDnsZone(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetGateway`

```go
ctx := context.TODO()
id := workloadnetworks.NewGatewayID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "gatewayIdValue")

read, err := client.GetGateway(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetPortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

read, err := client.GetPortMirroring(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetPublicIP`

```go
ctx := context.TODO()
id := workloadnetworks.NewPublicIPID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "publicIPIdValue")

read, err := client.GetPublicIP(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetSegment`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

read, err := client.GetSegment(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVmGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

read, err := client.GetVMGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.GetVirtualMachine`

```go
ctx := context.TODO()
id := workloadnetworks.NewDefaultVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "virtualMachineIdValue")

read, err := client.GetVirtualMachine(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkloadNetworksClient.ListDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListDhcp(ctx, id)` can be used to do batched pagination
items, err := client.ListDhcpComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListDnsServices`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListDnsServices(ctx, id)` can be used to do batched pagination
items, err := client.ListDnsServicesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListDnsZones`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListDnsZones(ctx, id)` can be used to do batched pagination
items, err := client.ListDnsZonesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListGateways`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListGateways(ctx, id)` can be used to do batched pagination
items, err := client.ListGatewaysComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListPortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListPortMirroring(ctx, id)` can be used to do batched pagination
items, err := client.ListPortMirroringComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListPublicIPs`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListPublicIPs(ctx, id)` can be used to do batched pagination
items, err := client.ListPublicIPsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListSegments(ctx, id)` can be used to do batched pagination
items, err := client.ListSegmentsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListVMGroups`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListVMGroups(ctx, id)` can be used to do batched pagination
items, err := client.ListVMGroupsComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.ListVirtualMachines`

```go
ctx := context.TODO()
id := workloadnetworks.NewPrivateCloudID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue")

// alternatively `client.ListVirtualMachines(ctx, id)` can be used to do batched pagination
items, err := client.ListVirtualMachinesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDhcp`

```go
ctx := context.TODO()
id := workloadnetworks.NewDhcpConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dhcpIdValue")

payload := workloadnetworks.WorkloadNetworkDhcp{
	// ...
}


if err := client.UpdateDhcpThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDnsService`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsServiceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsServiceIdValue")

payload := workloadnetworks.WorkloadNetworkDnsService{
	// ...
}


if err := client.UpdateDnsServiceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateDnsZone`

```go
ctx := context.TODO()
id := workloadnetworks.NewDnsZoneID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "dnsZoneIdValue")

payload := workloadnetworks.WorkloadNetworkDnsZone{
	// ...
}


if err := client.UpdateDnsZoneThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdatePortMirroring`

```go
ctx := context.TODO()
id := workloadnetworks.NewPortMirroringProfileID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "portMirroringIdValue")

payload := workloadnetworks.WorkloadNetworkPortMirroring{
	// ...
}


if err := client.UpdatePortMirroringThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateSegments`

```go
ctx := context.TODO()
id := workloadnetworks.NewSegmentID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "segmentIdValue")

payload := workloadnetworks.WorkloadNetworkSegment{
	// ...
}


if err := client.UpdateSegmentsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `WorkloadNetworksClient.UpdateVMGroup`

```go
ctx := context.TODO()
id := workloadnetworks.NewVmGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "privateCloudValue", "vmGroupIdValue")

payload := workloadnetworks.WorkloadNetworkVMGroup{
	// ...
}


if err := client.UpdateVMGroupThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
