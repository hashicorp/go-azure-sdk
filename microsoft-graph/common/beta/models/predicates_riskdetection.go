package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionOperationPredicate struct {
	ActivityDateTime    *string
	AdditionalInfo      *string
	CorrelationId       *string
	DetectedDateTime    *string
	Id                  *string
	IpAddress           *string
	LastUpdatedDateTime *string
	ODataType           *string
	RequestId           *string
	RiskEventType       *string
	Source              *string
	UserDisplayName     *string
	UserId              *string
	UserPrincipalName   *string
}

func (p RiskDetectionOperationPredicate) Matches(input RiskDetection) bool {

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.AdditionalInfo != nil && (input.AdditionalInfo == nil || *p.AdditionalInfo != *input.AdditionalInfo) {
		return false
	}

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.DetectedDateTime != nil && (input.DetectedDateTime == nil || *p.DetectedDateTime != *input.DetectedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestId != nil && (input.RequestId == nil || *p.RequestId != *input.RequestId) {
		return false
	}

	if p.RiskEventType != nil && (input.RiskEventType == nil || *p.RiskEventType != *input.RiskEventType) {
		return false
	}

	if p.Source != nil && (input.Source == nil || *p.Source != *input.Source) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
