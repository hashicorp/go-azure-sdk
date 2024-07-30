package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOrganizationSettings struct {
	EnableMEMAutoEnroll *bool                                       `json:"enableMEMAutoEnroll,omitempty"`
	EnableSingleSignOn  *bool                                       `json:"enableSingleSignOn,omitempty"`
	Id                  *string                                     `json:"id,omitempty"`
	ODataType           *string                                     `json:"@odata.type,omitempty"`
	OsVersion           *CloudPCOrganizationSettingsOsVersion       `json:"osVersion,omitempty"`
	UserAccountType     *CloudPCOrganizationSettingsUserAccountType `json:"userAccountType,omitempty"`
	WindowsSettings     *CloudPCWindowsSettings                     `json:"windowsSettings,omitempty"`
}
