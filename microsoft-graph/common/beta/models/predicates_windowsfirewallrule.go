package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleOperationPredicate struct {
	Description             *string
	DisplayName             *string
	FilePath                *string
	LocalUserAuthorizations *string
	ODataType               *string
	PackageFamilyName       *string
	Protocol                *int64
	ServiceName             *string
}

func (p WindowsFirewallRuleOperationPredicate) Matches(input WindowsFirewallRule) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FilePath != nil && (input.FilePath == nil || *p.FilePath != *input.FilePath) {
		return false
	}

	if p.LocalUserAuthorizations != nil && (input.LocalUserAuthorizations == nil || *p.LocalUserAuthorizations != *input.LocalUserAuthorizations) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PackageFamilyName != nil && (input.PackageFamilyName == nil || *p.PackageFamilyName != *input.PackageFamilyName) {
		return false
	}

	if p.Protocol != nil && (input.Protocol == nil || *p.Protocol != *input.Protocol) {
		return false
	}

	if p.ServiceName != nil && (input.ServiceName == nil || *p.ServiceName != *input.ServiceName) {
		return false
	}

	return true
}
