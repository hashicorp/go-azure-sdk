package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConfigurationSettingItem struct {
	AppConfigKey      *string                                      `json:"appConfigKey,omitempty"`
	AppConfigKeyType  *AppConfigurationSettingItemAppConfigKeyType `json:"appConfigKeyType,omitempty"`
	AppConfigKeyValue *string                                      `json:"appConfigKeyValue,omitempty"`
	ODataType         *string                                      `json:"@odata.type,omitempty"`
}
