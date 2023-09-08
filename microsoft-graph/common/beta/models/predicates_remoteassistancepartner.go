package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistancePartnerOperationPredicate struct {
	DisplayName                     *string
	Id                              *string
	LastConnectionDateTime          *string
	ODataType                       *string
	OnboardingRequestExpiryDateTime *string
	OnboardingUrl                   *string
}

func (p RemoteAssistancePartnerOperationPredicate) Matches(input RemoteAssistancePartner) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastConnectionDateTime != nil && (input.LastConnectionDateTime == nil || *p.LastConnectionDateTime != *input.LastConnectionDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OnboardingRequestExpiryDateTime != nil && (input.OnboardingRequestExpiryDateTime == nil || *p.OnboardingRequestExpiryDateTime != *input.OnboardingRequestExpiryDateTime) {
		return false
	}

	if p.OnboardingUrl != nil && (input.OnboardingUrl == nil || *p.OnboardingUrl != *input.OnboardingUrl) {
		return false
	}

	return true
}
