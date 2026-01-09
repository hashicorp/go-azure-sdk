package androidmanagedstoreaccountenterprisesetting

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppRequest struct {
	ApproveAllPermissions *bool     `json:"approveAllPermissions,omitempty"`
	PackageIds            *[]string `json:"packageIds,omitempty"`
}
