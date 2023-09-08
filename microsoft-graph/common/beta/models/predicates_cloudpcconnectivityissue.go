package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityIssueOperationPredicate struct {
	DeviceId          *string
	ErrorCode         *string
	ErrorDateTime     *string
	ErrorDescription  *string
	Id                *string
	ODataType         *string
	RecommendedAction *string
	UserId            *string
}

func (p CloudPCConnectivityIssueOperationPredicate) Matches(input CloudPCConnectivityIssue) bool {

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.ErrorDateTime != nil && (input.ErrorDateTime == nil || *p.ErrorDateTime != *input.ErrorDateTime) {
		return false
	}

	if p.ErrorDescription != nil && (input.ErrorDescription == nil || *p.ErrorDescription != *input.ErrorDescription) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecommendedAction != nil && (input.RecommendedAction == nil || *p.RecommendedAction != *input.RecommendedAction) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
