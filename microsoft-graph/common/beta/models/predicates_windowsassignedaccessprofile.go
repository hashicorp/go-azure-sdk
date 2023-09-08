package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAssignedAccessProfileOperationPredicate struct {
	Id                 *string
	ODataType          *string
	ProfileName        *string
	ShowTaskBar        *bool
	StartMenuLayoutXml *string
}

func (p WindowsAssignedAccessProfileOperationPredicate) Matches(input WindowsAssignedAccessProfile) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProfileName != nil && (input.ProfileName == nil || *p.ProfileName != *input.ProfileName) {
		return false
	}

	if p.ShowTaskBar != nil && (input.ShowTaskBar == nil || *p.ShowTaskBar != *input.ShowTaskBar) {
		return false
	}

	if p.StartMenuLayoutXml != nil && (input.StartMenuLayoutXml == nil || *p.StartMenuLayoutXml != *input.StartMenuLayoutXml) {
		return false
	}

	return true
}
