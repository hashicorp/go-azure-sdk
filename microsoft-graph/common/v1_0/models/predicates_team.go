package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamOperationPredicate struct {
	Classification  *string
	CreatedDateTime *string
	Description     *string
	DisplayName     *string
	Id              *string
	InternalId      *string
	IsArchived      *bool
	ODataType       *string
	TenantId        *string
	WebUrl          *string
}

func (p TeamOperationPredicate) Matches(input Team) bool {

	if p.Classification != nil && (input.Classification == nil || *p.Classification != *input.Classification) {
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

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InternalId != nil && (input.InternalId == nil || *p.InternalId != *input.InternalId) {
		return false
	}

	if p.IsArchived != nil && (input.IsArchived == nil || *p.IsArchived != *input.IsArchived) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
