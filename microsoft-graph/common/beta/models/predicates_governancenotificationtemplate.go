package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceNotificationTemplateOperationPredicate struct {
	Culture   *string
	Id        *string
	ODataType *string
	Source    *string
	Type      *string
	Version   *string
}

func (p GovernanceNotificationTemplateOperationPredicate) Matches(input GovernanceNotificationTemplate) bool {

	if p.Culture != nil && (input.Culture == nil || *p.Culture != *input.Culture) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Source != nil && (input.Source == nil || *p.Source != *input.Source) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
