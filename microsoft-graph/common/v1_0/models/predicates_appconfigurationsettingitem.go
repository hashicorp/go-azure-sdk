package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConfigurationSettingItemOperationPredicate struct {
	AppConfigKey      *string
	AppConfigKeyValue *string
	ODataType         *string
}

func (p AppConfigurationSettingItemOperationPredicate) Matches(input AppConfigurationSettingItem) bool {

	if p.AppConfigKey != nil && (input.AppConfigKey == nil || *p.AppConfigKey != *input.AppConfigKey) {
		return false
	}

	if p.AppConfigKeyValue != nil && (input.AppConfigKeyValue == nil || *p.AppConfigKeyValue != *input.AppConfigKeyValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
