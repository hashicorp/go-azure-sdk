package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnsupportedGroupPolicyExtension struct {
	ExtensionType *string                                      `json:"extensionType,omitempty"`
	Id            *string                                      `json:"id,omitempty"`
	NamespaceUrl  *string                                      `json:"namespaceUrl,omitempty"`
	NodeName      *string                                      `json:"nodeName,omitempty"`
	ODataType     *string                                      `json:"@odata.type,omitempty"`
	SettingScope  *UnsupportedGroupPolicyExtensionSettingScope `json:"settingScope,omitempty"`
}
