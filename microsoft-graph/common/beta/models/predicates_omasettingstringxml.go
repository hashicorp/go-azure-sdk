package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OmaSettingStringXmlOperationPredicate struct {
	Description            *string
	DisplayName            *string
	FileName               *string
	IsEncrypted            *bool
	ODataType              *string
	OmaUri                 *string
	SecretReferenceValueId *string
	Value                  *string
}

func (p OmaSettingStringXmlOperationPredicate) Matches(input OmaSettingStringXml) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FileName != nil && (input.FileName == nil || *p.FileName != *input.FileName) {
		return false
	}

	if p.IsEncrypted != nil && (input.IsEncrypted == nil || *p.IsEncrypted != *input.IsEncrypted) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OmaUri != nil && (input.OmaUri == nil || *p.OmaUri != *input.OmaUri) {
		return false
	}

	if p.SecretReferenceValueId != nil && (input.SecretReferenceValueId == nil || *p.SecretReferenceValueId != *input.SecretReferenceValueId) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
