package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentOperationPredicate struct {
	AssignedTo         *string
	CreatedDateTime    *string
	Description        *string
	DisplayName        *string
	Id                 *string
	IncidentWebUrl     *string
	LastModifiedBy     *string
	LastUpdateDateTime *string
	ODataType          *string
	RecommendedActions *string
	RedirectIncidentId *string
	TenantId           *string
}

func (p SecurityIncidentOperationPredicate) Matches(input SecurityIncident) bool {

	if p.AssignedTo != nil && (input.AssignedTo == nil || *p.AssignedTo != *input.AssignedTo) {
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

	if p.IncidentWebUrl != nil && (input.IncidentWebUrl == nil || *p.IncidentWebUrl != *input.IncidentWebUrl) {
		return false
	}

	if p.LastModifiedBy != nil && (input.LastModifiedBy == nil || *p.LastModifiedBy != *input.LastModifiedBy) {
		return false
	}

	if p.LastUpdateDateTime != nil && (input.LastUpdateDateTime == nil || *p.LastUpdateDateTime != *input.LastUpdateDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecommendedActions != nil && (input.RecommendedActions == nil || *p.RecommendedActions != *input.RecommendedActions) {
		return false
	}

	if p.RedirectIncidentId != nil && (input.RedirectIncidentId == nil || *p.RedirectIncidentId != *input.RedirectIncidentId) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
