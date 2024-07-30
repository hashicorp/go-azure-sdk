package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetails struct {
	AuthMethod        *UserCredentialUsageDetailsAuthMethod `json:"authMethod,omitempty"`
	EventDateTime     *string                               `json:"eventDateTime,omitempty"`
	FailureReason     *string                               `json:"failureReason,omitempty"`
	Feature           *UserCredentialUsageDetailsFeature    `json:"feature,omitempty"`
	Id                *string                               `json:"id,omitempty"`
	IsSuccess         *bool                                 `json:"isSuccess,omitempty"`
	ODataType         *string                               `json:"@odata.type,omitempty"`
	UserDisplayName   *string                               `json:"userDisplayName,omitempty"`
	UserPrincipalName *string                               `json:"userPrincipalName,omitempty"`
}
