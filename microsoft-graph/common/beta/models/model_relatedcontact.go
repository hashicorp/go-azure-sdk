package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedContact struct {
	AccessConsent *bool                       `json:"accessConsent,omitempty"`
	DisplayName   *string                     `json:"displayName,omitempty"`
	EmailAddress  *string                     `json:"emailAddress,omitempty"`
	Id            *string                     `json:"id,omitempty"`
	MobilePhone   *string                     `json:"mobilePhone,omitempty"`
	ODataType     *string                     `json:"@odata.type,omitempty"`
	Relationship  *RelatedContactRelationship `json:"relationship,omitempty"`
}
