package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetails struct {
	DefaultMfaMethod                              *UserRegistrationDetailsDefaultMfaMethod                              `json:"defaultMfaMethod,omitempty"`
	Id                                            *string                                                               `json:"id,omitempty"`
	IsAdmin                                       *bool                                                                 `json:"isAdmin,omitempty"`
	IsMfaCapable                                  *bool                                                                 `json:"isMfaCapable,omitempty"`
	IsMfaRegistered                               *bool                                                                 `json:"isMfaRegistered,omitempty"`
	IsPasswordlessCapable                         *bool                                                                 `json:"isPasswordlessCapable,omitempty"`
	IsSsprCapable                                 *bool                                                                 `json:"isSsprCapable,omitempty"`
	IsSsprEnabled                                 *bool                                                                 `json:"isSsprEnabled,omitempty"`
	IsSsprRegistered                              *bool                                                                 `json:"isSsprRegistered,omitempty"`
	IsSystemPreferredAuthenticationMethodEnabled  *bool                                                                 `json:"isSystemPreferredAuthenticationMethodEnabled,omitempty"`
	LastUpdatedDateTime                           *string                                                               `json:"lastUpdatedDateTime,omitempty"`
	MethodsRegistered                             *[]string                                                             `json:"methodsRegistered,omitempty"`
	ODataType                                     *string                                                               `json:"@odata.type,omitempty"`
	SystemPreferredAuthenticationMethods          *[]string                                                             `json:"systemPreferredAuthenticationMethods,omitempty"`
	UserDisplayName                               *string                                                               `json:"userDisplayName,omitempty"`
	UserPreferredMethodForSecondaryAuthentication *UserRegistrationDetailsUserPreferredMethodForSecondaryAuthentication `json:"userPreferredMethodForSecondaryAuthentication,omitempty"`
	UserPrincipalName                             *string                                                               `json:"userPrincipalName,omitempty"`
	UserType                                      *UserRegistrationDetailsUserType                                      `json:"userType,omitempty"`
}
