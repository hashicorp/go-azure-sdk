package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCitationTemplateOperationPredicate struct {
	CitationJurisdiction *string
	CitationUrl          *string
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	ODataType            *string
}

func (p SecurityCitationTemplateOperationPredicate) Matches(input SecurityCitationTemplate) bool {

	if p.CitationJurisdiction != nil && (input.CitationJurisdiction == nil || *p.CitationJurisdiction != *input.CitationJurisdiction) {
		return false
	}

	if p.CitationUrl != nil && (input.CitationUrl == nil || *p.CitationUrl != *input.CitationUrl) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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
