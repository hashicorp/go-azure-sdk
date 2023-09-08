package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFeatureUpdateProfileOperationPredicate struct {
	CreatedDateTime              *string
	DeployableContentDisplayName *string
	Description                  *string
	DisplayName                  *string
	EndOfSupportDate             *string
	FeatureUpdateVersion         *string
	Id                           *string
	LastModifiedDateTime         *string
	ODataType                    *string
}

func (p WindowsFeatureUpdateProfileOperationPredicate) Matches(input WindowsFeatureUpdateProfile) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeployableContentDisplayName != nil && (input.DeployableContentDisplayName == nil || *p.DeployableContentDisplayName != *input.DeployableContentDisplayName) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndOfSupportDate != nil && (input.EndOfSupportDate == nil || *p.EndOfSupportDate != *input.EndOfSupportDate) {
		return false
	}

	if p.FeatureUpdateVersion != nil && (input.FeatureUpdateVersion == nil || *p.FeatureUpdateVersion != *input.FeatureUpdateVersion) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
