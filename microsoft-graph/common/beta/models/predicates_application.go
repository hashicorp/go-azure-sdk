package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationOperationPredicate struct {
	AppId                      *string
	CreatedDateTime            *string
	DefaultRedirectUri         *string
	DeletedDateTime            *string
	Description                *string
	DisabledByMicrosoftStatus  *string
	DisplayName                *string
	GroupMembershipClaims      *string
	Id                         *string
	IsDeviceOnlyAuthSupported  *bool
	IsFallbackPublicClient     *bool
	Logo                       *string
	Notes                      *string
	ODataType                  *string
	PublisherDomain            *string
	SamlMetadataUrl            *string
	ServiceManagementReference *string
	SignInAudience             *string
	TokenEncryptionKeyId       *string
	UniqueName                 *string
}

func (p ApplicationOperationPredicate) Matches(input Application) bool {

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DefaultRedirectUri != nil && (input.DefaultRedirectUri == nil || *p.DefaultRedirectUri != *input.DefaultRedirectUri) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisabledByMicrosoftStatus != nil && (input.DisabledByMicrosoftStatus == nil || *p.DisabledByMicrosoftStatus != *input.DisabledByMicrosoftStatus) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.GroupMembershipClaims != nil && (input.GroupMembershipClaims == nil || *p.GroupMembershipClaims != *input.GroupMembershipClaims) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDeviceOnlyAuthSupported != nil && (input.IsDeviceOnlyAuthSupported == nil || *p.IsDeviceOnlyAuthSupported != *input.IsDeviceOnlyAuthSupported) {
		return false
	}

	if p.IsFallbackPublicClient != nil && (input.IsFallbackPublicClient == nil || *p.IsFallbackPublicClient != *input.IsFallbackPublicClient) {
		return false
	}

	if p.Logo != nil && (input.Logo == nil || *p.Logo != *input.Logo) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PublisherDomain != nil && (input.PublisherDomain == nil || *p.PublisherDomain != *input.PublisherDomain) {
		return false
	}

	if p.SamlMetadataUrl != nil && (input.SamlMetadataUrl == nil || *p.SamlMetadataUrl != *input.SamlMetadataUrl) {
		return false
	}

	if p.ServiceManagementReference != nil && (input.ServiceManagementReference == nil || *p.ServiceManagementReference != *input.ServiceManagementReference) {
		return false
	}

	if p.SignInAudience != nil && (input.SignInAudience == nil || *p.SignInAudience != *input.SignInAudience) {
		return false
	}

	if p.TokenEncryptionKeyId != nil && (input.TokenEncryptionKeyId == nil || *p.TokenEncryptionKeyId != *input.TokenEncryptionKeyId) {
		return false
	}

	if p.UniqueName != nil && (input.UniqueName == nil || *p.UniqueName != *input.UniqueName) {
		return false
	}

	return true
}
