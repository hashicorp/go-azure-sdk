package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemOperationPredicate struct {
	Id                      *string
	InitiatedBy             *string
	IsDeleted               *bool
	IsProcessing            *bool
	ODataType               *string
	RiskLastUpdatedDateTime *string
	UserDisplayName         *string
	UserId                  *string
	UserPrincipalName       *string
}

func (p RiskyUserHistoryItemOperationPredicate) Matches(input RiskyUserHistoryItem) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InitiatedBy != nil && (input.InitiatedBy == nil || *p.InitiatedBy != *input.InitiatedBy) {
		return false
	}

	if p.IsDeleted != nil && (input.IsDeleted == nil || *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.IsProcessing != nil && (input.IsProcessing == nil || *p.IsProcessing != *input.IsProcessing) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RiskLastUpdatedDateTime != nil && (input.RiskLastUpdatedDateTime == nil || *p.RiskLastUpdatedDateTime != *input.RiskLastUpdatedDateTime) {
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
