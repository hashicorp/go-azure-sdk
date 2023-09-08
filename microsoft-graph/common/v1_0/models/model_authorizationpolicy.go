package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationPolicy struct {
	AllowEmailVerifiedUsersToJoinOrganization *bool                                `json:"allowEmailVerifiedUsersToJoinOrganization,omitempty"`
	AllowInvitesFrom                          *AuthorizationPolicyAllowInvitesFrom `json:"allowInvitesFrom,omitempty"`
	AllowUserConsentForRiskyApps              *bool                                `json:"allowUserConsentForRiskyApps,omitempty"`
	AllowedToSignUpEmailBasedSubscriptions    *bool                                `json:"allowedToSignUpEmailBasedSubscriptions,omitempty"`
	AllowedToUseSSPR                          *bool                                `json:"allowedToUseSSPR,omitempty"`
	BlockMsolPowerShell                       *bool                                `json:"blockMsolPowerShell,omitempty"`
	DefaultUserRolePermissions                *DefaultUserRolePermissions          `json:"defaultUserRolePermissions,omitempty"`
	DeletedDateTime                           *string                              `json:"deletedDateTime,omitempty"`
	Description                               *string                              `json:"description,omitempty"`
	DisplayName                               *string                              `json:"displayName,omitempty"`
	GuestUserRoleId                           *string                              `json:"guestUserRoleId,omitempty"`
	Id                                        *string                              `json:"id,omitempty"`
	ODataType                                 *string                              `json:"@odata.type,omitempty"`
}
