package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleSummary struct {
	ElevatedCount *int64                       `json:"elevatedCount,omitempty"`
	Id            *string                      `json:"id,omitempty"`
	ManagedCount  *int64                       `json:"managedCount,omitempty"`
	MfaEnabled    *bool                        `json:"mfaEnabled,omitempty"`
	ODataType     *string                      `json:"@odata.type,omitempty"`
	Status        *PrivilegedRoleSummaryStatus `json:"status,omitempty"`
	UsersCount    *int64                       `json:"usersCount,omitempty"`
}
