package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionOperationPredicate struct {
	AdDomainName           *string
	AdDomainPassword       *string
	AdDomainUsername       *string
	AlternateResourceUrl   *string
	DisplayName            *string
	Id                     *string
	InUse                  *bool
	ODataType              *string
	OrganizationalUnit     *string
	ResourceGroupId        *string
	SubnetId               *string
	SubscriptionId         *string
	SubscriptionName       *string
	VirtualNetworkId       *string
	VirtualNetworkLocation *string
}

func (p CloudPCOnPremisesConnectionOperationPredicate) Matches(input CloudPCOnPremisesConnection) bool {

	if p.AdDomainName != nil && (input.AdDomainName == nil || *p.AdDomainName != *input.AdDomainName) {
		return false
	}

	if p.AdDomainPassword != nil && (input.AdDomainPassword == nil || *p.AdDomainPassword != *input.AdDomainPassword) {
		return false
	}

	if p.AdDomainUsername != nil && (input.AdDomainUsername == nil || *p.AdDomainUsername != *input.AdDomainUsername) {
		return false
	}

	if p.AlternateResourceUrl != nil && (input.AlternateResourceUrl == nil || *p.AlternateResourceUrl != *input.AlternateResourceUrl) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InUse != nil && (input.InUse == nil || *p.InUse != *input.InUse) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrganizationalUnit != nil && (input.OrganizationalUnit == nil || *p.OrganizationalUnit != *input.OrganizationalUnit) {
		return false
	}

	if p.ResourceGroupId != nil && (input.ResourceGroupId == nil || *p.ResourceGroupId != *input.ResourceGroupId) {
		return false
	}

	if p.SubnetId != nil && (input.SubnetId == nil || *p.SubnetId != *input.SubnetId) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil || *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	if p.SubscriptionName != nil && (input.SubscriptionName == nil || *p.SubscriptionName != *input.SubscriptionName) {
		return false
	}

	if p.VirtualNetworkId != nil && (input.VirtualNetworkId == nil || *p.VirtualNetworkId != *input.VirtualNetworkId) {
		return false
	}

	if p.VirtualNetworkLocation != nil && (input.VirtualNetworkLocation == nil || *p.VirtualNetworkLocation != *input.VirtualNetworkLocation) {
		return false
	}

	return true
}
