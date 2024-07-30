package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCAccountManagerPolicy struct {
	AccountDeletionPolicy                 *SharedPCAccountManagerPolicyAccountDeletionPolicy `json:"accountDeletionPolicy,omitempty"`
	CacheAccountsAboveDiskFreePercentage  *int64                                             `json:"cacheAccountsAboveDiskFreePercentage,omitempty"`
	InactiveThresholdDays                 *int64                                             `json:"inactiveThresholdDays,omitempty"`
	ODataType                             *string                                            `json:"@odata.type,omitempty"`
	RemoveAccountsBelowDiskFreePercentage *int64                                             `json:"removeAccountsBelowDiskFreePercentage,omitempty"`
}
