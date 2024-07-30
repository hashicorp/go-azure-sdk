package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCTenantEncryptionSetting struct {
	LastSyncDateTime         *string                                                 `json:"lastSyncDateTime,omitempty"`
	ODataType                *string                                                 `json:"@odata.type,omitempty"`
	TenantDiskEncryptionType *CloudPCTenantEncryptionSettingTenantDiskEncryptionType `json:"tenantDiskEncryptionType,omitempty"`
}
