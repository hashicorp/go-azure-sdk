package hypervsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteHealthSummaryOperationPredicate struct {
	AffectedObjectsCount *int64
	AffectedResourceType *string
	ApplianceName        *string
	ErrorCode            *string
	ErrorId              *int64
	ErrorMessage         *string
	HitCount             *int64
	RemediationGuidance  *string
	Severity             *string
	SummaryMessage       *string
}

func (p SiteHealthSummaryOperationPredicate) Matches(input SiteHealthSummary) bool {

	if p.AffectedObjectsCount != nil && (input.AffectedObjectsCount == nil || *p.AffectedObjectsCount != *input.AffectedObjectsCount) {
		return false
	}

	if p.AffectedResourceType != nil && (input.AffectedResourceType == nil || *p.AffectedResourceType != *input.AffectedResourceType) {
		return false
	}

	if p.ApplianceName != nil && (input.ApplianceName == nil || *p.ApplianceName != *input.ApplianceName) {
		return false
	}

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.ErrorId != nil && (input.ErrorId == nil || *p.ErrorId != *input.ErrorId) {
		return false
	}

	if p.ErrorMessage != nil && (input.ErrorMessage == nil || *p.ErrorMessage != *input.ErrorMessage) {
		return false
	}

	if p.HitCount != nil && (input.HitCount == nil || *p.HitCount != *input.HitCount) {
		return false
	}

	if p.RemediationGuidance != nil && (input.RemediationGuidance == nil || *p.RemediationGuidance != *input.RemediationGuidance) {
		return false
	}

	if p.Severity != nil && (input.Severity == nil || *p.Severity != *input.Severity) {
		return false
	}

	if p.SummaryMessage != nil && (input.SummaryMessage == nil || *p.SummaryMessage != *input.SummaryMessage) {
		return false
	}

	return true
}
