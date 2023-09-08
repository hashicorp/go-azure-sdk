package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOperationPredicate struct {
	AadDeviceId              *string
	DisplayName              *string
	GracePeriodEndDateTime   *string
	Id                       *string
	ImageDisplayName         *string
	LastModifiedDateTime     *string
	ManagedDeviceId          *string
	ManagedDeviceName        *string
	ODataType                *string
	OnPremisesConnectionName *string
	ProvisioningPolicyId     *string
	ProvisioningPolicyName   *string
	ServicePlanId            *string
	ServicePlanName          *string
	UserPrincipalName        *string
}

func (p CloudPCOperationPredicate) Matches(input CloudPC) bool {

	if p.AadDeviceId != nil && (input.AadDeviceId == nil || *p.AadDeviceId != *input.AadDeviceId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.GracePeriodEndDateTime != nil && (input.GracePeriodEndDateTime == nil || *p.GracePeriodEndDateTime != *input.GracePeriodEndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ImageDisplayName != nil && (input.ImageDisplayName == nil || *p.ImageDisplayName != *input.ImageDisplayName) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ManagedDeviceName != nil && (input.ManagedDeviceName == nil || *p.ManagedDeviceName != *input.ManagedDeviceName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnPremisesConnectionName != nil && (input.OnPremisesConnectionName == nil || *p.OnPremisesConnectionName != *input.OnPremisesConnectionName) {
		return false
	}

	if p.ProvisioningPolicyId != nil && (input.ProvisioningPolicyId == nil || *p.ProvisioningPolicyId != *input.ProvisioningPolicyId) {
		return false
	}

	if p.ProvisioningPolicyName != nil && (input.ProvisioningPolicyName == nil || *p.ProvisioningPolicyName != *input.ProvisioningPolicyName) {
		return false
	}

	if p.ServicePlanId != nil && (input.ServicePlanId == nil || *p.ServicePlanId != *input.ServicePlanId) {
		return false
	}

	if p.ServicePlanName != nil && (input.ServicePlanName == nil || *p.ServicePlanName != *input.ServicePlanName) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
