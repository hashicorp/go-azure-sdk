package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityProfileOperationPredicate struct {
	AzureSubscriptionId    *string
	AzureTenantId          *string
	CreatedDateTime        *string
	DeploymentPackageUrl   *string
	DestinationServiceName *string
	Id                     *string
	IsSigned               *bool
	LastModifiedDateTime   *string
	Manifest               *string
	Name                   *string
	ODataType              *string
	Platform               *string
	PolicyName             *string
	Publisher              *string
	RiskScore              *string
	Type                   *string
}

func (p CloudAppSecurityProfileOperationPredicate) Matches(input CloudAppSecurityProfile) bool {

	if p.AzureSubscriptionId != nil && (input.AzureSubscriptionId == nil || *p.AzureSubscriptionId != *input.AzureSubscriptionId) {
		return false
	}

	if p.AzureTenantId != nil && (input.AzureTenantId == nil || *p.AzureTenantId != *input.AzureTenantId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeploymentPackageUrl != nil && (input.DeploymentPackageUrl == nil || *p.DeploymentPackageUrl != *input.DeploymentPackageUrl) {
		return false
	}

	if p.DestinationServiceName != nil && (input.DestinationServiceName == nil || *p.DestinationServiceName != *input.DestinationServiceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSigned != nil && (input.IsSigned == nil || *p.IsSigned != *input.IsSigned) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Manifest != nil && (input.Manifest == nil || *p.Manifest != *input.Manifest) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Platform != nil && (input.Platform == nil || *p.Platform != *input.Platform) {
		return false
	}

	if p.PolicyName != nil && (input.PolicyName == nil || *p.PolicyName != *input.PolicyName) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.RiskScore != nil && (input.RiskScore == nil || *p.RiskScore != *input.RiskScore) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
