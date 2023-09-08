package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrength struct {
	AuthenticationStrengthId     *string                                             `json:"authenticationStrengthId,omitempty"`
	AuthenticationStrengthResult *AuthenticationStrengthAuthenticationStrengthResult `json:"authenticationStrengthResult,omitempty"`
	DisplayName                  *string                                             `json:"displayName,omitempty"`
	ODataType                    *string                                             `json:"@odata.type,omitempty"`
}
