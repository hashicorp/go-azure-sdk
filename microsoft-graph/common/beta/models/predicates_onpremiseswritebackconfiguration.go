package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesWritebackConfigurationOperationPredicate struct {
	ODataType             *string
	UnifiedGroupContainer *string
	UserContainer         *string
}

func (p OnPremisesWritebackConfigurationOperationPredicate) Matches(input OnPremisesWritebackConfiguration) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UnifiedGroupContainer != nil && (input.UnifiedGroupContainer == nil || *p.UnifiedGroupContainer != *input.UnifiedGroupContainer) {
		return false
	}

	if p.UserContainer != nil && (input.UserContainer == nil || *p.UserContainer != *input.UserContainer) {
		return false
	}

	return true
}
