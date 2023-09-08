package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSettings struct {
	DurationUntilExternalUserDeletedAfterBlocked *string                                                   `json:"durationUntilExternalUserDeletedAfterBlocked,omitempty"`
	ExternalUserLifecycleAction                  *EntitlementManagementSettingsExternalUserLifecycleAction `json:"externalUserLifecycleAction,omitempty"`
	Id                                           *string                                                   `json:"id,omitempty"`
	ODataType                                    *string                                                   `json:"@odata.type,omitempty"`
}
