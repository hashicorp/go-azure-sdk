package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationOperationPredicate struct {
	AuthenticationBlockPeriodInMinutes      *int64
	AuthenticationPeriodInSeconds           *int64
	AuthenticationRetryDelayPeriodInSeconds *int64
	CacheCredentials                        *bool
	CreatedDateTime                         *string
	Description                             *string
	DisableUserPromptForServerValidation    *bool
	DisplayName                             *string
	EapolStartPeriodInSeconds               *int64
	Enforce8021X                            *bool
	ForceFIPSCompliance                     *bool
	Id                                      *string
	LastModifiedDateTime                    *string
	MaximumAuthenticationFailures           *int64
	MaximumEAPOLStartMessages               *int64
	ODataType                               *string
	OuterIdentityPrivacyTemporaryValue      *string
	PerformServerValidation                 *bool
	RequireCryptographicBinding             *bool
	SupportsScopeTags                       *bool
	Version                                 *int64
}

func (p WindowsWiredNetworkConfigurationOperationPredicate) Matches(input WindowsWiredNetworkConfiguration) bool {

	if p.AuthenticationBlockPeriodInMinutes != nil && (input.AuthenticationBlockPeriodInMinutes == nil || *p.AuthenticationBlockPeriodInMinutes != *input.AuthenticationBlockPeriodInMinutes) {
		return false
	}

	if p.AuthenticationPeriodInSeconds != nil && (input.AuthenticationPeriodInSeconds == nil || *p.AuthenticationPeriodInSeconds != *input.AuthenticationPeriodInSeconds) {
		return false
	}

	if p.AuthenticationRetryDelayPeriodInSeconds != nil && (input.AuthenticationRetryDelayPeriodInSeconds == nil || *p.AuthenticationRetryDelayPeriodInSeconds != *input.AuthenticationRetryDelayPeriodInSeconds) {
		return false
	}

	if p.CacheCredentials != nil && (input.CacheCredentials == nil || *p.CacheCredentials != *input.CacheCredentials) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisableUserPromptForServerValidation != nil && (input.DisableUserPromptForServerValidation == nil || *p.DisableUserPromptForServerValidation != *input.DisableUserPromptForServerValidation) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EapolStartPeriodInSeconds != nil && (input.EapolStartPeriodInSeconds == nil || *p.EapolStartPeriodInSeconds != *input.EapolStartPeriodInSeconds) {
		return false
	}

	if p.Enforce8021X != nil && (input.Enforce8021X == nil || *p.Enforce8021X != *input.Enforce8021X) {
		return false
	}

	if p.ForceFIPSCompliance != nil && (input.ForceFIPSCompliance == nil || *p.ForceFIPSCompliance != *input.ForceFIPSCompliance) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MaximumAuthenticationFailures != nil && (input.MaximumAuthenticationFailures == nil || *p.MaximumAuthenticationFailures != *input.MaximumAuthenticationFailures) {
		return false
	}

	if p.MaximumEAPOLStartMessages != nil && (input.MaximumEAPOLStartMessages == nil || *p.MaximumEAPOLStartMessages != *input.MaximumEAPOLStartMessages) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OuterIdentityPrivacyTemporaryValue != nil && (input.OuterIdentityPrivacyTemporaryValue == nil || *p.OuterIdentityPrivacyTemporaryValue != *input.OuterIdentityPrivacyTemporaryValue) {
		return false
	}

	if p.PerformServerValidation != nil && (input.PerformServerValidation == nil || *p.PerformServerValidation != *input.PerformServerValidation) {
		return false
	}

	if p.RequireCryptographicBinding != nil && (input.RequireCryptographicBinding == nil || *p.RequireCryptographicBinding != *input.RequireCryptographicBinding) {
		return false
	}

	if p.SupportsScopeTags != nil && (input.SupportsScopeTags == nil || *p.SupportsScopeTags != *input.SupportsScopeTags) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
