package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDelegatedScopeAppSetting struct {
	AppDetail *AppListItem                                         `json:"appDetail,omitempty"`
	AppScopes *AndroidDeviceOwnerDelegatedScopeAppSettingAppScopes `json:"appScopes,omitempty"`
	ODataType *string                                              `json:"@odata.type,omitempty"`
}
