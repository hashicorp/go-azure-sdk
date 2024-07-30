package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupLifecyclePolicy struct {
	AlternateNotificationEmails *string `json:"alternateNotificationEmails,omitempty"`
	GroupLifetimeInDays         *int64  `json:"groupLifetimeInDays,omitempty"`
	Id                          *string `json:"id,omitempty"`
	ManagedGroupTypes           *string `json:"managedGroupTypes,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
}
