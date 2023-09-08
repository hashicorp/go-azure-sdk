package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionOperationPredicate struct {
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	ShortDescription     *string
	TeamsAppId           *string
	Version              *string
}

func (p TeamsAppDefinitionOperationPredicate) Matches(input TeamsAppDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.ShortDescription != nil && (input.ShortDescription == nil || *p.ShortDescription != *input.ShortDescription) {
		return false
	}

	if p.TeamsAppId != nil && (input.TeamsAppId == nil || *p.TeamsAppId != *input.TeamsAppId) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
