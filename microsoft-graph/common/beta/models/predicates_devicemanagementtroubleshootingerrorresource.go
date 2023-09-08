package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTroubleshootingErrorResourceOperationPredicate struct {
	Link      *string
	ODataType *string
	Text      *string
}

func (p DeviceManagementTroubleshootingErrorResourceOperationPredicate) Matches(input DeviceManagementTroubleshootingErrorResource) bool {

	if p.Link != nil && (input.Link == nil || *p.Link != *input.Link) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	return true
}
