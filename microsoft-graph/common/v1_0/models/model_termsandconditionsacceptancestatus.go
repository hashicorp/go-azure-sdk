package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionsAcceptanceStatus struct {
	AcceptedDateTime   *string             `json:"acceptedDateTime,omitempty"`
	AcceptedVersion    *int64              `json:"acceptedVersion,omitempty"`
	Id                 *string             `json:"id,omitempty"`
	ODataType          *string             `json:"@odata.type,omitempty"`
	TermsAndConditions *TermsAndConditions `json:"termsAndConditions,omitempty"`
	UserDisplayName    *string             `json:"userDisplayName,omitempty"`
	UserPrincipalName  *string             `json:"userPrincipalName,omitempty"`
}
