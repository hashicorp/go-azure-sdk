package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationOperationPredicate struct {
	ApiVersion           *int64
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	IsActive             *bool
	LastModifiedDateTime *string
	ODataType            *string
	Url                  *string
}

func (p WorkforceIntegrationOperationPredicate) Matches(input WorkforceIntegration) bool {

	if p.ApiVersion != nil && (input.ApiVersion == nil || *p.ApiVersion != *input.ApiVersion) {
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

	if p.IsActive != nil && (input.IsActive == nil || *p.IsActive != *input.IsActive) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Url != nil && (input.Url == nil || *p.Url != *input.Url) {
		return false
	}

	return true
}
