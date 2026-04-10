package networkinterfaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonLoadBalancerOperationPredicate struct {
	Etag     *string
	Id       *string
	Location *string
	Name     *string
	Type     *string
}

func (p CommonLoadBalancerOperationPredicate) Matches(input CommonLoadBalancer) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
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

type CommonNetworkInterfaceOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p CommonNetworkInterfaceOperationPredicate) Matches(input CommonNetworkInterface) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

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

type CommonNetworkInterfaceIPConfigurationOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p CommonNetworkInterfaceIPConfigurationOperationPredicate) Matches(input CommonNetworkInterfaceIPConfiguration) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

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

type CommonNetworkInterfaceTapConfigurationOperationPredicate struct {
	Etag *string
	Id   *string
	Name *string
	Type *string
}

func (p CommonNetworkInterfaceTapConfigurationOperationPredicate) Matches(input CommonNetworkInterfaceTapConfiguration) bool {

	if p.Etag != nil && (input.Etag == nil || *p.Etag != *input.Etag) {
		return false
	}

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

type EffectiveNetworkSecurityGroupOperationPredicate struct {
}

func (p EffectiveNetworkSecurityGroupOperationPredicate) Matches(input EffectiveNetworkSecurityGroup) bool {

	return true
}

type EffectiveRouteOperationPredicate struct {
	DisableBgpRoutePropagation *bool
	Name                       *string
}

func (p EffectiveRouteOperationPredicate) Matches(input EffectiveRoute) bool {

	if p.DisableBgpRoutePropagation != nil && (input.DisableBgpRoutePropagation == nil || *p.DisableBgpRoutePropagation != *input.DisableBgpRoutePropagation) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	return true
}
