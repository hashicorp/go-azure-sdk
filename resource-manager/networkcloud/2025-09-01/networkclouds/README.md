
## `github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/networkclouds` Documentation

The `networkclouds` SDK allows for interaction with Azure Resource Manager `networkcloud` (API Version `2025-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
import "github.com/hashicorp/go-azure-sdk/resource-manager/networkcloud/2025-09-01/networkclouds"
```


### Client Initialization

```go
client := networkclouds.NewNetworkcloudsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NetworkcloudsClient.AgentPoolsCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

payload := networkclouds.AgentPool{
	// ...
}


if err := client.AgentPoolsCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultAgentPoolsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.AgentPoolsDelete`

```go
ctx := context.TODO()
id := networkclouds.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

if err := client.AgentPoolsDeleteThenPoll(ctx, id, networkclouds.DefaultAgentPoolsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.AgentPoolsGet`

```go
ctx := context.TODO()
id := networkclouds.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

read, err := client.AgentPoolsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.AgentPoolsListByKubernetesCluster`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

// alternatively `client.AgentPoolsListByKubernetesCluster(ctx, id, networkclouds.DefaultAgentPoolsListByKubernetesClusterOperationOptions())` can be used to do batched pagination
items, err := client.AgentPoolsListByKubernetesClusterComplete(ctx, id, networkclouds.DefaultAgentPoolsListByKubernetesClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.AgentPoolsUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewAgentPoolID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "agentPoolName")

payload := networkclouds.AgentPoolPatchParameters{
	// ...
}


if err := client.AgentPoolsUpdateThenPoll(ctx, id, payload, networkclouds.DefaultAgentPoolsUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachineKeySetsCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

payload := networkclouds.BareMetalMachineKeySet{
	// ...
}


if err := client.BareMetalMachineKeySetsCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBareMetalMachineKeySetsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachineKeySetsDelete`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

if err := client.BareMetalMachineKeySetsDeleteThenPoll(ctx, id, networkclouds.DefaultBareMetalMachineKeySetsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachineKeySetsGet`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

read, err := client.BareMetalMachineKeySetsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachineKeySetsListByCluster`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.BareMetalMachineKeySetsListByCluster(ctx, id, networkclouds.DefaultBareMetalMachineKeySetsListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.BareMetalMachineKeySetsListByClusterComplete(ctx, id, networkclouds.DefaultBareMetalMachineKeySetsListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachineKeySetsUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bareMetalMachineKeySetName")

payload := networkclouds.BareMetalMachineKeySetPatchParameters{
	// ...
}


if err := client.BareMetalMachineKeySetsUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBareMetalMachineKeySetsUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesCordon`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineCordonParameters{
	// ...
}


if err := client.BareMetalMachinesCordonThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachine{
	// ...
}


if err := client.BareMetalMachinesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBareMetalMachinesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.BareMetalMachinesDeleteThenPoll(ctx, id, networkclouds.DefaultBareMetalMachinesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesGet`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

read, err := client.BareMetalMachinesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.BareMetalMachinesListByResourceGroup(ctx, id, networkclouds.DefaultBareMetalMachinesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.BareMetalMachinesListByResourceGroupComplete(ctx, id, networkclouds.DefaultBareMetalMachinesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.BareMetalMachinesListBySubscription(ctx, id, networkclouds.DefaultBareMetalMachinesListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.BareMetalMachinesListBySubscriptionComplete(ctx, id, networkclouds.DefaultBareMetalMachinesListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesPowerOff`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachinePowerOffParameters{
	// ...
}


if err := client.BareMetalMachinesPowerOffThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesReimage`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.BareMetalMachinesReimageThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesReplace`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineReplaceParameters{
	// ...
}


if err := client.BareMetalMachinesReplaceThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesRestart`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.BareMetalMachinesRestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesRunCommand`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineRunCommandParameters{
	// ...
}


if err := client.BareMetalMachinesRunCommandThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesRunDataExtracts`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineRunDataExtractsParameters{
	// ...
}


if err := client.BareMetalMachinesRunDataExtractsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesRunDataExtractsRestricted`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineRunDataExtractsParameters{
	// ...
}


if err := client.BareMetalMachinesRunDataExtractsRestrictedThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesRunReadCommands`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachineRunReadCommandsParameters{
	// ...
}


if err := client.BareMetalMachinesRunReadCommandsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesStart`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.BareMetalMachinesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesUncordon`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

if err := client.BareMetalMachinesUncordonThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BareMetalMachinesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBareMetalMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "bareMetalMachineName")

payload := networkclouds.BareMetalMachinePatchParameters{
	// ...
}


if err := client.BareMetalMachinesUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBareMetalMachinesUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BmcKeySetsCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

payload := networkclouds.BmcKeySet{
	// ...
}


if err := client.BmcKeySetsCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBmcKeySetsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BmcKeySetsDelete`

```go
ctx := context.TODO()
id := networkclouds.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

if err := client.BmcKeySetsDeleteThenPoll(ctx, id, networkclouds.DefaultBmcKeySetsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.BmcKeySetsGet`

```go
ctx := context.TODO()
id := networkclouds.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

read, err := client.BmcKeySetsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.BmcKeySetsListByCluster`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.BmcKeySetsListByCluster(ctx, id, networkclouds.DefaultBmcKeySetsListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.BmcKeySetsListByClusterComplete(ctx, id, networkclouds.DefaultBmcKeySetsListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.BmcKeySetsUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewBmcKeySetID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "bmcKeySetName")

payload := networkclouds.BmcKeySetPatchParameters{
	// ...
}


if err := client.BmcKeySetsUpdateThenPoll(ctx, id, payload, networkclouds.DefaultBmcKeySetsUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

payload := networkclouds.CloudServicesNetwork{
	// ...
}


if err := client.CloudServicesNetworksCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultCloudServicesNetworksCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksDelete`

```go
ctx := context.TODO()
id := networkclouds.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

if err := client.CloudServicesNetworksDeleteThenPoll(ctx, id, networkclouds.DefaultCloudServicesNetworksDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksGet`

```go
ctx := context.TODO()
id := networkclouds.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

read, err := client.CloudServicesNetworksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.CloudServicesNetworksListByResourceGroup(ctx, id, networkclouds.DefaultCloudServicesNetworksListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.CloudServicesNetworksListByResourceGroupComplete(ctx, id, networkclouds.DefaultCloudServicesNetworksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.CloudServicesNetworksListBySubscription(ctx, id, networkclouds.DefaultCloudServicesNetworksListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.CloudServicesNetworksListBySubscriptionComplete(ctx, id, networkclouds.DefaultCloudServicesNetworksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.CloudServicesNetworksUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewCloudServicesNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "cloudServicesNetworkName")

payload := networkclouds.CloudServicesNetworkPatchParameters{
	// ...
}


if err := client.CloudServicesNetworksUpdateThenPoll(ctx, id, payload, networkclouds.DefaultCloudServicesNetworksUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

payload := networkclouds.ClusterManager{
	// ...
}


if err := client.ClusterManagersCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultClusterManagersCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersDelete`

```go
ctx := context.TODO()
id := networkclouds.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

if err := client.ClusterManagersDeleteThenPoll(ctx, id, networkclouds.DefaultClusterManagersDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersGet`

```go
ctx := context.TODO()
id := networkclouds.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

read, err := client.ClusterManagersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ClusterManagersListByResourceGroup(ctx, id, networkclouds.DefaultClusterManagersListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ClusterManagersListByResourceGroupComplete(ctx, id, networkclouds.DefaultClusterManagersListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ClusterManagersListBySubscription(ctx, id, networkclouds.DefaultClusterManagersListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ClusterManagersListBySubscriptionComplete(ctx, id, networkclouds.DefaultClusterManagersListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.ClusterManagersUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewClusterManagerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterManagerName")

payload := networkclouds.ClusterManagerPatchParameters{
	// ...
}


read, err := client.ClusterManagersUpdate(ctx, id, payload, networkclouds.DefaultClusterManagersUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.ClustersContinueUpdateVersion`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.ClusterContinueUpdateVersionParameters{
	// ...
}


if err := client.ClustersContinueUpdateVersionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.Cluster{
	// ...
}


if err := client.ClustersCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultClustersCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersDelete`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

if err := client.ClustersDeleteThenPoll(ctx, id, networkclouds.DefaultClustersDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersDeploy`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.ClusterDeployParameters{
	// ...
}


if err := client.ClustersDeployThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersGet`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

read, err := client.ClustersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.ClustersListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ClustersListByResourceGroup(ctx, id, networkclouds.DefaultClustersListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.ClustersListByResourceGroupComplete(ctx, id, networkclouds.DefaultClustersListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.ClustersListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.ClustersListBySubscription(ctx, id, networkclouds.DefaultClustersListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.ClustersListBySubscriptionComplete(ctx, id, networkclouds.DefaultClustersListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.ClustersScanRuntime`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.ClusterScanRuntimeParameters{
	// ...
}


if err := client.ClustersScanRuntimeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.ClusterPatchParameters{
	// ...
}


if err := client.ClustersUpdateThenPoll(ctx, id, payload, networkclouds.DefaultClustersUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ClustersUpdateVersion`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

payload := networkclouds.ClusterUpdateVersionParameters{
	// ...
}


if err := client.ClustersUpdateVersionThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ConsolesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

payload := networkclouds.Console{
	// ...
}


if err := client.ConsolesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultConsolesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ConsolesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

if err := client.ConsolesDeleteThenPoll(ctx, id, networkclouds.DefaultConsolesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.ConsolesGet`

```go
ctx := context.TODO()
id := networkclouds.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

read, err := client.ConsolesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.ConsolesListByVirtualMachine`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

// alternatively `client.ConsolesListByVirtualMachine(ctx, id, networkclouds.DefaultConsolesListByVirtualMachineOperationOptions())` can be used to do batched pagination
items, err := client.ConsolesListByVirtualMachineComplete(ctx, id, networkclouds.DefaultConsolesListByVirtualMachineOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.ConsolesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewConsoleID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName", "consoleName")

payload := networkclouds.ConsolePatchParameters{
	// ...
}


if err := client.ConsolesUpdateThenPoll(ctx, id, payload, networkclouds.DefaultConsolesUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClusterFeaturesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

payload := networkclouds.KubernetesClusterFeature{
	// ...
}


if err := client.KubernetesClusterFeaturesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultKubernetesClusterFeaturesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClusterFeaturesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

if err := client.KubernetesClusterFeaturesDeleteThenPoll(ctx, id, networkclouds.DefaultKubernetesClusterFeaturesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClusterFeaturesGet`

```go
ctx := context.TODO()
id := networkclouds.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

read, err := client.KubernetesClusterFeaturesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClusterFeaturesListByKubernetesCluster`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

// alternatively `client.KubernetesClusterFeaturesListByKubernetesCluster(ctx, id, networkclouds.DefaultKubernetesClusterFeaturesListByKubernetesClusterOperationOptions())` can be used to do batched pagination
items, err := client.KubernetesClusterFeaturesListByKubernetesClusterComplete(ctx, id, networkclouds.DefaultKubernetesClusterFeaturesListByKubernetesClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClusterFeaturesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewFeatureID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName", "featureName")

payload := networkclouds.KubernetesClusterFeaturePatchParameters{
	// ...
}


if err := client.KubernetesClusterFeaturesUpdateThenPoll(ctx, id, payload, networkclouds.DefaultKubernetesClusterFeaturesUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

payload := networkclouds.KubernetesCluster{
	// ...
}


if err := client.KubernetesClustersCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultKubernetesClustersCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersDelete`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

if err := client.KubernetesClustersDeleteThenPoll(ctx, id, networkclouds.DefaultKubernetesClustersDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersGet`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

read, err := client.KubernetesClustersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.KubernetesClustersListByResourceGroup(ctx, id, networkclouds.DefaultKubernetesClustersListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.KubernetesClustersListByResourceGroupComplete(ctx, id, networkclouds.DefaultKubernetesClustersListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.KubernetesClustersListBySubscription(ctx, id, networkclouds.DefaultKubernetesClustersListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.KubernetesClustersListBySubscriptionComplete(ctx, id, networkclouds.DefaultKubernetesClustersListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersRestartNode`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

payload := networkclouds.KubernetesClusterRestartNodeParameters{
	// ...
}


if err := client.KubernetesClustersRestartNodeThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.KubernetesClustersUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewKubernetesClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "kubernetesClusterName")

payload := networkclouds.KubernetesClusterPatchParameters{
	// ...
}


if err := client.KubernetesClustersUpdateThenPoll(ctx, id, payload, networkclouds.DefaultKubernetesClustersUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewL2NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l2NetworkName")

payload := networkclouds.L2Network{
	// ...
}


if err := client.L2NetworksCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultL2NetworksCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksDelete`

```go
ctx := context.TODO()
id := networkclouds.NewL2NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l2NetworkName")

if err := client.L2NetworksDeleteThenPoll(ctx, id, networkclouds.DefaultL2NetworksDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksGet`

```go
ctx := context.TODO()
id := networkclouds.NewL2NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l2NetworkName")

read, err := client.L2NetworksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.L2NetworksListByResourceGroup(ctx, id, networkclouds.DefaultL2NetworksListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.L2NetworksListByResourceGroupComplete(ctx, id, networkclouds.DefaultL2NetworksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.L2NetworksListBySubscription(ctx, id, networkclouds.DefaultL2NetworksListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.L2NetworksListBySubscriptionComplete(ctx, id, networkclouds.DefaultL2NetworksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.L2NetworksUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewL2NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l2NetworkName")

payload := networkclouds.L2NetworkPatchParameters{
	// ...
}


read, err := client.L2NetworksUpdate(ctx, id, payload, networkclouds.DefaultL2NetworksUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

payload := networkclouds.L3Network{
	// ...
}


if err := client.L3NetworksCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultL3NetworksCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksDelete`

```go
ctx := context.TODO()
id := networkclouds.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

if err := client.L3NetworksDeleteThenPoll(ctx, id, networkclouds.DefaultL3NetworksDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksGet`

```go
ctx := context.TODO()
id := networkclouds.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

read, err := client.L3NetworksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.L3NetworksListByResourceGroup(ctx, id, networkclouds.DefaultL3NetworksListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.L3NetworksListByResourceGroupComplete(ctx, id, networkclouds.DefaultL3NetworksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.L3NetworksListBySubscription(ctx, id, networkclouds.DefaultL3NetworksListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.L3NetworksListBySubscriptionComplete(ctx, id, networkclouds.DefaultL3NetworksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.L3NetworksUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewL3NetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "l3NetworkName")

payload := networkclouds.L3NetworkPatchParameters{
	// ...
}


read, err := client.L3NetworksUpdate(ctx, id, payload, networkclouds.DefaultL3NetworksUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.MetricsConfigurationsCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

payload := networkclouds.ClusterMetricsConfiguration{
	// ...
}


if err := client.MetricsConfigurationsCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultMetricsConfigurationsCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.MetricsConfigurationsDelete`

```go
ctx := context.TODO()
id := networkclouds.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

if err := client.MetricsConfigurationsDeleteThenPoll(ctx, id, networkclouds.DefaultMetricsConfigurationsDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.MetricsConfigurationsGet`

```go
ctx := context.TODO()
id := networkclouds.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

read, err := client.MetricsConfigurationsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.MetricsConfigurationsListByCluster`

```go
ctx := context.TODO()
id := networkclouds.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName")

// alternatively `client.MetricsConfigurationsListByCluster(ctx, id, networkclouds.DefaultMetricsConfigurationsListByClusterOperationOptions())` can be used to do batched pagination
items, err := client.MetricsConfigurationsListByClusterComplete(ctx, id, networkclouds.DefaultMetricsConfigurationsListByClusterOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.MetricsConfigurationsUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewMetricsConfigurationID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterName", "metricsConfigurationName")

payload := networkclouds.ClusterMetricsConfigurationPatchParameters{
	// ...
}


if err := client.MetricsConfigurationsUpdateThenPoll(ctx, id, payload, networkclouds.DefaultMetricsConfigurationsUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.RackSkusGet`

```go
ctx := context.TODO()
id := networkclouds.NewRackSkuID("12345678-1234-9876-4563-123456789012", "rackSkuName")

read, err := client.RackSkusGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.RackSkusListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.RackSkusListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.RackSkusListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.RacksCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

payload := networkclouds.Rack{
	// ...
}


if err := client.RacksCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultRacksCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.RacksDelete`

```go
ctx := context.TODO()
id := networkclouds.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

if err := client.RacksDeleteThenPoll(ctx, id, networkclouds.DefaultRacksDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.RacksGet`

```go
ctx := context.TODO()
id := networkclouds.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

read, err := client.RacksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.RacksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.RacksListByResourceGroup(ctx, id, networkclouds.DefaultRacksListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.RacksListByResourceGroupComplete(ctx, id, networkclouds.DefaultRacksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.RacksListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.RacksListBySubscription(ctx, id, networkclouds.DefaultRacksListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.RacksListBySubscriptionComplete(ctx, id, networkclouds.DefaultRacksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.RacksUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewRackID("12345678-1234-9876-4563-123456789012", "example-resource-group", "rackName")

payload := networkclouds.RackPatchParameters{
	// ...
}


if err := client.RacksUpdateThenPoll(ctx, id, payload, networkclouds.DefaultRacksUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := networkclouds.StorageAppliance{
	// ...
}


if err := client.StorageAppliancesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultStorageAppliancesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

if err := client.StorageAppliancesDeleteThenPoll(ctx, id, networkclouds.DefaultStorageAppliancesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesDisableRemoteVendorManagement`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

if err := client.StorageAppliancesDisableRemoteVendorManagementThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesEnableRemoteVendorManagement`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := networkclouds.StorageApplianceEnableRemoteVendorManagementParameters{
	// ...
}


if err := client.StorageAppliancesEnableRemoteVendorManagementThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesGet`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

read, err := client.StorageAppliancesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.StorageAppliancesListByResourceGroup(ctx, id, networkclouds.DefaultStorageAppliancesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.StorageAppliancesListByResourceGroupComplete(ctx, id, networkclouds.DefaultStorageAppliancesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.StorageAppliancesListBySubscription(ctx, id, networkclouds.DefaultStorageAppliancesListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.StorageAppliancesListBySubscriptionComplete(ctx, id, networkclouds.DefaultStorageAppliancesListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesRunReadCommands`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := networkclouds.StorageApplianceRunReadCommandsParameters{
	// ...
}


if err := client.StorageAppliancesRunReadCommandsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.StorageAppliancesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewStorageApplianceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "storageApplianceName")

payload := networkclouds.StorageAppliancePatchParameters{
	// ...
}


if err := client.StorageAppliancesUpdateThenPoll(ctx, id, payload, networkclouds.DefaultStorageAppliancesUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

payload := networkclouds.TrunkedNetwork{
	// ...
}


if err := client.TrunkedNetworksCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultTrunkedNetworksCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksDelete`

```go
ctx := context.TODO()
id := networkclouds.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

if err := client.TrunkedNetworksDeleteThenPoll(ctx, id, networkclouds.DefaultTrunkedNetworksDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksGet`

```go
ctx := context.TODO()
id := networkclouds.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

read, err := client.TrunkedNetworksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.TrunkedNetworksListByResourceGroup(ctx, id, networkclouds.DefaultTrunkedNetworksListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.TrunkedNetworksListByResourceGroupComplete(ctx, id, networkclouds.DefaultTrunkedNetworksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.TrunkedNetworksListBySubscription(ctx, id, networkclouds.DefaultTrunkedNetworksListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.TrunkedNetworksListBySubscriptionComplete(ctx, id, networkclouds.DefaultTrunkedNetworksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.TrunkedNetworksUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewTrunkedNetworkID("12345678-1234-9876-4563-123456789012", "example-resource-group", "trunkedNetworkName")

payload := networkclouds.TrunkedNetworkPatchParameters{
	// ...
}


read, err := client.TrunkedNetworksUpdate(ctx, id, payload, networkclouds.DefaultTrunkedNetworksUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesAssignRelay`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := networkclouds.VirtualMachineAssignRelayParameters{
	// ...
}


if err := client.VirtualMachinesAssignRelayThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := networkclouds.VirtualMachine{
	// ...
}


if err := client.VirtualMachinesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultVirtualMachinesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.VirtualMachinesDeleteThenPoll(ctx, id, networkclouds.DefaultVirtualMachinesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesGet`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

read, err := client.VirtualMachinesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.VirtualMachinesListByResourceGroup(ctx, id, networkclouds.DefaultVirtualMachinesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.VirtualMachinesListByResourceGroupComplete(ctx, id, networkclouds.DefaultVirtualMachinesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.VirtualMachinesListBySubscription(ctx, id, networkclouds.DefaultVirtualMachinesListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.VirtualMachinesListBySubscriptionComplete(ctx, id, networkclouds.DefaultVirtualMachinesListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesPowerOff`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := networkclouds.VirtualMachinePowerOffParameters{
	// ...
}


if err := client.VirtualMachinesPowerOffThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesReimage`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.VirtualMachinesReimageThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesRestart`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.VirtualMachinesRestartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesStart`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

if err := client.VirtualMachinesStartThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VirtualMachinesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewVirtualMachineID("12345678-1234-9876-4563-123456789012", "example-resource-group", "virtualMachineName")

payload := networkclouds.VirtualMachinePatchParameters{
	// ...
}


if err := client.VirtualMachinesUpdateThenPoll(ctx, id, payload, networkclouds.DefaultVirtualMachinesUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VolumesCreateOrUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

payload := networkclouds.Volume{
	// ...
}


if err := client.VolumesCreateOrUpdateThenPoll(ctx, id, payload, networkclouds.DefaultVolumesCreateOrUpdateOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VolumesDelete`

```go
ctx := context.TODO()
id := networkclouds.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

if err := client.VolumesDeleteThenPoll(ctx, id, networkclouds.DefaultVolumesDeleteOperationOptions()); err != nil {
	// handle the error
}
```


### Example Usage: `NetworkcloudsClient.VolumesGet`

```go
ctx := context.TODO()
id := networkclouds.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

read, err := client.VolumesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `NetworkcloudsClient.VolumesListByResourceGroup`

```go
ctx := context.TODO()
id := commonids.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.VolumesListByResourceGroup(ctx, id, networkclouds.DefaultVolumesListByResourceGroupOperationOptions())` can be used to do batched pagination
items, err := client.VolumesListByResourceGroupComplete(ctx, id, networkclouds.DefaultVolumesListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.VolumesListBySubscription`

```go
ctx := context.TODO()
id := commonids.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.VolumesListBySubscription(ctx, id, networkclouds.DefaultVolumesListBySubscriptionOperationOptions())` can be used to do batched pagination
items, err := client.VolumesListBySubscriptionComplete(ctx, id, networkclouds.DefaultVolumesListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `NetworkcloudsClient.VolumesUpdate`

```go
ctx := context.TODO()
id := networkclouds.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "volumeName")

payload := networkclouds.VolumePatchParameters{
	// ...
}


read, err := client.VolumesUpdate(ctx, id, payload, networkclouds.DefaultVolumesUpdateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
