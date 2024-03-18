package vmwares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDhcpOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkDhcpOperationPredicate) Matches(input WorkloadNetworkDhcp) bool {

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

type WorkloadNetworkDnsServiceOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkDnsServiceOperationPredicate) Matches(input WorkloadNetworkDnsService) bool {

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

type WorkloadNetworkDnsZoneOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkDnsZoneOperationPredicate) Matches(input WorkloadNetworkDnsZone) bool {

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

type WorkloadNetworkPortMirroringOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkPortMirroringOperationPredicate) Matches(input WorkloadNetworkPortMirroring) bool {

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

type WorkloadNetworkPublicIPOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkPublicIPOperationPredicate) Matches(input WorkloadNetworkPublicIP) bool {

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

type WorkloadNetworkVMGroupOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p WorkloadNetworkVMGroupOperationPredicate) Matches(input WorkloadNetworkVMGroup) bool {

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
