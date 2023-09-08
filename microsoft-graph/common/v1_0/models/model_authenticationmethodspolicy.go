package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicy struct {
	AuthenticationMethodConfigurations *[]AuthenticationMethodConfiguration             `json:"authenticationMethodConfigurations,omitempty"`
	Description                        *string                                          `json:"description,omitempty"`
	DisplayName                        *string                                          `json:"displayName,omitempty"`
	Id                                 *string                                          `json:"id,omitempty"`
	LastModifiedDateTime               *string                                          `json:"lastModifiedDateTime,omitempty"`
	ODataType                          *string                                          `json:"@odata.type,omitempty"`
	PolicyMigrationState               *AuthenticationMethodsPolicyPolicyMigrationState `json:"policyMigrationState,omitempty"`
	PolicyVersion                      *string                                          `json:"policyVersion,omitempty"`
	ReconfirmationInDays               *int64                                           `json:"reconfirmationInDays,omitempty"`
	RegistrationEnforcement            *RegistrationEnforcement                         `json:"registrationEnforcement,omitempty"`
}
