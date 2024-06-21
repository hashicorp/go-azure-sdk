package user

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProperties struct {
	AdditionalUsageQuota   *string                 `json:"additionalUsageQuota,omitempty"`
	DisplayName            *string                 `json:"displayName,omitempty"`
	Email                  string                  `json:"email"`
	InvitationSent         *string                 `json:"invitationSent,omitempty"`
	InvitationState        *InvitationState        `json:"invitationState,omitempty"`
	ProvisioningState      *ProvisioningState      `json:"provisioningState,omitempty"`
	RegistrationState      *RegistrationState      `json:"registrationState,omitempty"`
	ResourceOperationError *ResourceOperationError `json:"resourceOperationError,omitempty"`
	TotalUsage             *string                 `json:"totalUsage,omitempty"`
}
