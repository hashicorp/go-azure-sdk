package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionOperationPredicate struct {
	AzureADAppId         *string
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Shortdescription     *string
	TeamsAppId           *string
	Version              *string
}

func (p TeamsAppDefinitionOperationPredicate) Matches(input TeamsAppDefinition) bool {

	if p.AzureADAppId != nil && (input.AzureADAppId == nil || *p.AzureADAppId != *input.AzureADAppId) {
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

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Shortdescription != nil && (input.Shortdescription == nil || *p.Shortdescription != *input.Shortdescription) {
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
