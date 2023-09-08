package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestOperationPredicate struct {
	ClosedDateTime         *string
	ContentQuery           *string
	CreatedDateTime        *string
	Description            *string
	DisplayName            *string
	ExternalId             *string
	Id                     *string
	IncludeAllVersions     *bool
	IncludeAuthoredContent *bool
	InternalDueDateTime    *string
	LastModifiedDateTime   *string
	ODataType              *string
	PauseAfterEstimate     *bool
}

func (p SubjectRightsRequestOperationPredicate) Matches(input SubjectRightsRequest) bool {

	if p.ClosedDateTime != nil && (input.ClosedDateTime == nil || *p.ClosedDateTime != *input.ClosedDateTime) {
		return false
	}

	if p.ContentQuery != nil && (input.ContentQuery == nil || *p.ContentQuery != *input.ContentQuery) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IncludeAllVersions != nil && (input.IncludeAllVersions == nil || *p.IncludeAllVersions != *input.IncludeAllVersions) {
		return false
	}

	if p.IncludeAuthoredContent != nil && (input.IncludeAuthoredContent == nil || *p.IncludeAuthoredContent != *input.IncludeAuthoredContent) {
		return false
	}

	if p.InternalDueDateTime != nil && (input.InternalDueDateTime == nil || *p.InternalDueDateTime != *input.InternalDueDateTime) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PauseAfterEstimate != nil && (input.PauseAfterEstimate == nil || *p.PauseAfterEstimate != *input.PauseAfterEstimate) {
		return false
	}

	return true
}
