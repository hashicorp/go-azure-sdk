package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentPoolOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p AgentPoolOperationPredicate) Matches(input AgentPool) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type BareMetalMachineOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p BareMetalMachineOperationPredicate) Matches(input BareMetalMachine) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type BareMetalMachineKeySetOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p BareMetalMachineKeySetOperationPredicate) Matches(input BareMetalMachineKeySet) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type BmcKeySetOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p BmcKeySetOperationPredicate) Matches(input BmcKeySet) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type CloudServicesNetworkOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p CloudServicesNetworkOperationPredicate) Matches(input CloudServicesNetwork) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ClusterOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ClusterOperationPredicate) Matches(input Cluster) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ClusterManagerOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ClusterManagerOperationPredicate) Matches(input ClusterManager) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ClusterMetricsConfigurationOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ClusterMetricsConfigurationOperationPredicate) Matches(input ClusterMetricsConfiguration) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type ConsoleOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p ConsoleOperationPredicate) Matches(input Console) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type KubernetesClusterOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p KubernetesClusterOperationPredicate) Matches(input KubernetesCluster) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type L2NetworkOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p L2NetworkOperationPredicate) Matches(input L2Network) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type L3NetworkOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p L3NetworkOperationPredicate) Matches(input L3Network) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type RackOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p RackOperationPredicate) Matches(input Rack) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type RackSkuOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p RackSkuOperationPredicate) Matches(input RackSku) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type StorageApplianceOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p StorageApplianceOperationPredicate) Matches(input StorageAppliance) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type TrunkedNetworkOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p TrunkedNetworkOperationPredicate) Matches(input TrunkedNetwork) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type VirtualMachineOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p VirtualMachineOperationPredicate) Matches(input VirtualMachine) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}

type VolumeOperationPredicate struct {
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p VolumeOperationPredicate) Matches(input Volume) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && *p.Location != input.Location {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
