package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessRoot struct {
	AuthenticationContextClassReferences *[]AuthenticationContextClassReference `json:"authenticationContextClassReferences,omitempty"`
	AuthenticationStrength               *AuthenticationStrengthRoot            `json:"authenticationStrength,omitempty"`
	AuthenticationStrengths              *AuthenticationStrengthRoot            `json:"authenticationStrengths,omitempty"`
	Id                                   *string                                `json:"id,omitempty"`
	NamedLocations                       *[]NamedLocation                       `json:"namedLocations,omitempty"`
	ODataType                            *string                                `json:"@odata.type,omitempty"`
	Policies                             *[]ConditionalAccessPolicy             `json:"policies,omitempty"`
	Templates                            *[]ConditionalAccessTemplate           `json:"templates,omitempty"`
}
