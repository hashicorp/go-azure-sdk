package androidmanagedstoreaccountenterprisesetting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAndroidManagedStoreAccountEnterpriseSettingApproveAppRequest struct {
	ApproveAllPermissions *bool     `json:"approveAllPermissions,omitempty"`
	PackageIds            *[]string `json:"packageIds,omitempty"`
}
