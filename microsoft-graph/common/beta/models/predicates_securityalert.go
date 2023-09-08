package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertOperationPredicate struct {
	ActorDisplayName      *string
	AlertPolicyId         *string
	AlertWebUrl           *string
	AssignedTo            *string
	Category              *string
	CreatedDateTime       *string
	Description           *string
	DetectorId            *string
	FirstActivityDateTime *string
	Id                    *string
	IncidentId            *string
	IncidentWebUrl        *string
	LastActivityDateTime  *string
	LastUpdateDateTime    *string
	ODataType             *string
	ProviderAlertId       *string
	RecommendedActions    *string
	ResolvedDateTime      *string
	TenantId              *string
	ThreatDisplayName     *string
	ThreatFamilyName      *string
	Title                 *string
}

func (p SecurityAlertOperationPredicate) Matches(input SecurityAlert) bool {

	if p.ActorDisplayName != nil && (input.ActorDisplayName == nil || *p.ActorDisplayName != *input.ActorDisplayName) {
		return false
	}

	if p.AlertPolicyId != nil && (input.AlertPolicyId == nil || *p.AlertPolicyId != *input.AlertPolicyId) {
		return false
	}

	if p.AlertWebUrl != nil && (input.AlertWebUrl == nil || *p.AlertWebUrl != *input.AlertWebUrl) {
		return false
	}

	if p.AssignedTo != nil && (input.AssignedTo == nil || *p.AssignedTo != *input.AssignedTo) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DetectorId != nil && (input.DetectorId == nil || *p.DetectorId != *input.DetectorId) {
		return false
	}

	if p.FirstActivityDateTime != nil && (input.FirstActivityDateTime == nil || *p.FirstActivityDateTime != *input.FirstActivityDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IncidentId != nil && (input.IncidentId == nil || *p.IncidentId != *input.IncidentId) {
		return false
	}

	if p.IncidentWebUrl != nil && (input.IncidentWebUrl == nil || *p.IncidentWebUrl != *input.IncidentWebUrl) {
		return false
	}

	if p.LastActivityDateTime != nil && (input.LastActivityDateTime == nil || *p.LastActivityDateTime != *input.LastActivityDateTime) {
		return false
	}

	if p.LastUpdateDateTime != nil && (input.LastUpdateDateTime == nil || *p.LastUpdateDateTime != *input.LastUpdateDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProviderAlertId != nil && (input.ProviderAlertId == nil || *p.ProviderAlertId != *input.ProviderAlertId) {
		return false
	}

	if p.RecommendedActions != nil && (input.RecommendedActions == nil || *p.RecommendedActions != *input.RecommendedActions) {
		return false
	}

	if p.ResolvedDateTime != nil && (input.ResolvedDateTime == nil || *p.ResolvedDateTime != *input.ResolvedDateTime) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.ThreatDisplayName != nil && (input.ThreatDisplayName == nil || *p.ThreatDisplayName != *input.ThreatDisplayName) {
		return false
	}

	if p.ThreatFamilyName != nil && (input.ThreatFamilyName == nil || *p.ThreatFamilyName != *input.ThreatFamilyName) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
