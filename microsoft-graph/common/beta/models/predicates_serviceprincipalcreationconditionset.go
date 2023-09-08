package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipalCreationConditionSetOperationPredicate struct {
	ApplicationsFromVerifiedPublisherOnly *bool
	CertifiedApplicationsOnly             *bool
	Id                                    *string
	ODataType                             *string
}

func (p ServicePrincipalCreationConditionSetOperationPredicate) Matches(input ServicePrincipalCreationConditionSet) bool {

	if p.ApplicationsFromVerifiedPublisherOnly != nil && (input.ApplicationsFromVerifiedPublisherOnly == nil || *p.ApplicationsFromVerifiedPublisherOnly != *input.ApplicationsFromVerifiedPublisherOnly) {
		return false
	}

	if p.CertifiedApplicationsOnly != nil && (input.CertifiedApplicationsOnly == nil || *p.CertifiedApplicationsOnly != *input.CertifiedApplicationsOnly) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
