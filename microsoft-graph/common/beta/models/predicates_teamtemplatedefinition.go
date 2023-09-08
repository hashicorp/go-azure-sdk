package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateDefinitionOperationPredicate struct {
	Description          *string
	DisplayName          *string
	IconUrl              *string
	Id                   *string
	LanguageTag          *string
	LastModifiedDateTime *string
	ODataType            *string
	ParentTemplateId     *string
	PublisherName        *string
	ShortDescription     *string
}

func (p TeamTemplateDefinitionOperationPredicate) Matches(input TeamTemplateDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.IconUrl != nil && (input.IconUrl == nil || *p.IconUrl != *input.IconUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LanguageTag != nil && (input.LanguageTag == nil || *p.LanguageTag != *input.LanguageTag) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentTemplateId != nil && (input.ParentTemplateId == nil || *p.ParentTemplateId != *input.ParentTemplateId) {
		return false
	}

	if p.PublisherName != nil && (input.PublisherName == nil || *p.PublisherName != *input.PublisherName) {
		return false
	}

	if p.ShortDescription != nil && (input.ShortDescription == nil || *p.ShortDescription != *input.ShortDescription) {
		return false
	}

	return true
}
