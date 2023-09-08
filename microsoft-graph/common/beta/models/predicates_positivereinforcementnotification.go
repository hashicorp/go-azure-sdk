package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositiveReinforcementNotificationOperationPredicate struct {
	DefaultLanguage *string
	ODataType       *string
}

func (p PositiveReinforcementNotificationOperationPredicate) Matches(input PositiveReinforcementNotification) bool {

	if p.DefaultLanguage != nil && (input.DefaultLanguage == nil || *p.DefaultLanguage != *input.DefaultLanguage) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
