package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyUploadedDefinitionFileOperationPredicate struct {
	Content              *string
	DefaultLanguageCode  *string
	Description          *string
	DisplayName          *string
	FileName             *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Revision             *string
	TargetNamespace      *string
	TargetPrefix         *string
	UploadDateTime       *string
}

func (p GroupPolicyUploadedDefinitionFileOperationPredicate) Matches(input GroupPolicyUploadedDefinitionFile) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.DefaultLanguageCode != nil && (input.DefaultLanguageCode == nil || *p.DefaultLanguageCode != *input.DefaultLanguageCode) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
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

	if p.Revision != nil && (input.Revision == nil || *p.Revision != *input.Revision) {
		return false
	}

	if p.TargetNamespace != nil && (input.TargetNamespace == nil || *p.TargetNamespace != *input.TargetNamespace) {
		return false
	}

	if p.TargetPrefix != nil && (input.TargetPrefix == nil || *p.TargetPrefix != *input.TargetPrefix) {
		return false
	}

	if p.UploadDateTime != nil && (input.UploadDateTime == nil || *p.UploadDateTime != *input.UploadDateTime) {
		return false
	}

	return true
}
