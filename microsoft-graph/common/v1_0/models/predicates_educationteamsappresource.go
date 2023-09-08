package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationTeamsAppResourceOperationPredicate struct {
	AppIconWebUrl           *string
	AppId                   *string
	CreatedDateTime         *string
	DisplayName             *string
	LastModifiedDateTime    *string
	ODataType               *string
	TeamsEmbeddedContentUrl *string
	WebUrl                  *string
}

func (p EducationTeamsAppResourceOperationPredicate) Matches(input EducationTeamsAppResource) bool {

	if p.AppIconWebUrl != nil && (input.AppIconWebUrl == nil || *p.AppIconWebUrl != *input.AppIconWebUrl) {
		return false
	}

	if p.AppId != nil && (input.AppId == nil || *p.AppId != *input.AppId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamsEmbeddedContentUrl != nil && (input.TeamsEmbeddedContentUrl == nil || *p.TeamsEmbeddedContentUrl != *input.TeamsEmbeddedContentUrl) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
