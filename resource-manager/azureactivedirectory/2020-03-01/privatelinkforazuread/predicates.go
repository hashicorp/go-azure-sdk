package privatelinkforazuread

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkPolicyOperationPredicate struct {
	AllTenants     *bool
	Id             *string
	Name           *string
	OwnerTenantId  *string
	ResourceGroup  *string
	ResourceName   *string
	SubscriptionId *string
	Type           *string
}

func (p PrivateLinkPolicyOperationPredicate) Matches(input PrivateLinkPolicy) bool {

	if p.AllTenants != nil && (input.AllTenants == nil && *p.AllTenants != *input.AllTenants) {
		return false
	}

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.OwnerTenantId != nil && (input.OwnerTenantId == nil && *p.OwnerTenantId != *input.OwnerTenantId) {
		return false
	}

	if p.ResourceGroup != nil && (input.ResourceGroup == nil && *p.ResourceGroup != *input.ResourceGroup) {
		return false
	}

	if p.ResourceName != nil && (input.ResourceName == nil && *p.ResourceName != *input.ResourceName) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil && *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
